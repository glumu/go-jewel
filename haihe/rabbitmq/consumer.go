package rabbitmq

import (
	"errors"
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/streadway/amqp"
)

type Consumer struct {
	addr        string
	name        string
	topic       string
	routingKeys []string
	logger      *log.Logger
	connection  *amqp.Connection
	channel     *amqp.Channel
	done        chan bool
	notifyClose chan *amqp.Error
	isConnected bool
	recvHandler func(msg *[]byte) error
}

const (
	reconnectDelay = 5 * time.Second // 连接断开后多久重连
)

var (
	errNotConnected  = errors.New("not connected to the producer")
	errAlreadyClosed = errors.New("already closed: not connected to the producer")
)

func NewConsumer(host, port, username, password, name, topic string, logger *log.Logger) *Consumer {
	addr := fmt.Sprintf("amqp://%s:%s@%s:%s/",
		username,
		password,
		host,
		port)

	consumer := Consumer{
		logger: logger,
		addr:   addr,
		name:   name,
		topic:  topic,
		done:   make(chan bool),
	}
	return &consumer
}

func (consumer *Consumer) RegisterConsumerHandler(recvHandler func(msg *[]byte) error) {
	consumer.recvHandler = recvHandler
}

func (consumer *Consumer) RegisterRoutingKey(routingKey string) {
	consumer.routingKeys = append(consumer.routingKeys, routingKey)
}

func (consumer *Consumer) Start() {
	go consumer.handleReconnect(consumer.addr)
}

func (consumer *Consumer) handleReconnect(addr string) {
	for {
		consumer.isConnected = false
		consumer.logger.Println("[info] Attempting to connect")
		for !consumer.connect(addr) {
			consumer.logger.Println("[warning] Failed to connect. Retrying...")
			time.Sleep(reconnectDelay)
		}
		select {
		case <-consumer.done:
			return
		case <-consumer.notifyClose:
		}
	}
}

func (consumer *Consumer) connect(addr string) bool {
	conn, err := amqp.Dial(addr)
	if err != nil {
		consumer.logger.Printf("[error] Connect to mq error : %s", err.Error())
		return false
	}
	ch, err := conn.Channel()
	if err != nil {
		consumer.logger.Printf("[error] Open channel error : %s", err.Error())
		return false
	}

	err = ch.ExchangeDeclare(
		consumer.topic,
		"topic",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		consumer.logger.Printf("[error] Declare exchange error : %s", err.Error())
		return false
	}

	_, err = ch.QueueDeclare(
		consumer.name,
		true,  // Durable
		false, // Delete when unused
		false, // Exclusive
		false, // No-wait
		nil,   // Arguments
	)
	if err != nil {
		consumer.logger.Printf("[error] Declare queue error : %s", err.Error())
		return false
	}

	consumer.changeConnection(conn, ch)
	err = consumer.bindQueue()
	if err != nil {
		consumer.logger.Printf("[error] Bind queue error : %s", err.Error())
		return false
	}

	consumer.receive()
	if err != nil {
		consumer.logger.Printf("[error] Consume message error : %s", err.Error())
		return false
	}

	consumer.isConnected = true
	consumer.logger.Printf("Connected!")
	return true
}

func (consumer *Consumer) bindQueue() (err error) {
	for _, key := range consumer.routingKeys {
		err = consumer.channel.QueueBind(consumer.name, key, consumer.topic, false, nil)
		if err != nil {
			return err
		}
	}
	return nil
}

func (consumer *Consumer) receive() (err error) {
	consumer.channel.Qos(1, 0, true)
	consumerTag := fmt.Sprintf("amq.ctag-%s", getRandomString(22))
	msgs, err := consumer.channel.Consume(consumer.name, consumerTag, false, false, false, false, nil)
	if err != nil {
		return err
	}
	go func() {
		if consumer.recvHandler != nil {
			for d := range msgs {
				consumer.recvHandler(&d.Body)
				d.Ack(false)
			}
		}
	}()
	return nil
}

func (consumer *Consumer) changeConnection(connection *amqp.Connection, channel *amqp.Channel) {
	consumer.connection = connection
	consumer.channel = channel
	consumer.notifyClose = make(chan *amqp.Error)
	consumer.channel.NotifyClose(consumer.notifyClose)
}

func (consumer *Consumer) Close() error {
	if !consumer.isConnected {
		return errAlreadyClosed
	}
	err := consumer.channel.Close()
	if err != nil {
		return err
	}
	err = consumer.connection.Close()
	if err != nil {
		return err
	}
	close(consumer.done)
	consumer.isConnected = false
	return nil
}

func getRandomString(l int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyz"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < l; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}
