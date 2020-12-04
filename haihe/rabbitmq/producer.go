package rabbitmq

import (
	"fmt"
	"log"
	"time"

	"github.com/streadway/amqp"
)

type Producer struct {
	topic          string
	logger         *log.Logger
	connection     *amqp.Connection
	channel        *amqp.Channel
	msgs           chan *Message
	doneConn       chan bool
	donePublish    chan bool
	notifyClose    chan *amqp.Error
	notifyConfirm  chan amqp.Confirmation
	isConnected    bool
	confirmHandler func(id string) error
}

type Message struct {
	RoutingKey string
	Id         string
	Msg        string
}

func NewProducer(host, port, username, password, topic string, logger *log.Logger) *Producer {
	addr := fmt.Sprintf("amqp://%s:%s@%s:%s/",
		username,
		password,
		host,
		port)

	producer := Producer{
		logger:      logger,
		topic:       topic,
		doneConn:    make(chan bool),
		donePublish: make(chan bool),
		msgs:        make(chan *Message, 1024),
	}
	go producer.handleReconnect(addr)
	go producer.handlePublish()
	return &producer
}

func (producer *Producer) RegisterConfirmHandler(confirmHandler func(id string) error) {
	producer.confirmHandler = confirmHandler
}

func (producer *Producer) handleReconnect(addr string) {
	for {
		producer.isConnected = false
		log.Println("[info] Attempting to connect")
		for !producer.connect(addr) {
			log.Println("[warning] Failed to connect. Retrying...")
			time.Sleep(reconnectDelay)
		}
		select {
		case <-producer.doneConn:
			return
		case <-producer.notifyClose:
		}
	}
}

func (producer *Producer) connect(addr string) bool {
	conn, err := amqp.Dial(addr)
	if err != nil {
		return false
	}
	ch, err := conn.Channel()
	if err != nil {
		return false
	}

	err = ch.Confirm(false)
	if err != nil {
		producer.logger.Printf("[error] Set confirm mode error : %s", err.Error())
		return false
	}

	err = ch.ExchangeDeclare(
		producer.topic,
		"topic",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return false
	}
	producer.changeConnection(conn, ch)
	producer.isConnected = true
	log.Println("Connected!")
	return true
}

func (producer *Producer) changeConnection(connection *amqp.Connection, channel *amqp.Channel) {
	producer.connection = connection
	producer.channel = channel
	producer.notifyClose = make(chan *amqp.Error)
	producer.notifyConfirm = make(chan amqp.Confirmation, 1)
	producer.channel.NotifyClose(producer.notifyClose)
	producer.channel.NotifyPublish(producer.notifyConfirm)
}

func (producer *Producer) Publish(routingKey string, msgId string, msg string) {
	message := &Message{RoutingKey: routingKey, Id: msgId, Msg: msg}
	producer.msgs <- message
}

func (producer *Producer) handlePublish() {
	for {
		select {
		case <-producer.donePublish:
			return
		case msg, ok := <-producer.msgs:
			if !ok {
				producer.logger.Println("[error] Cannot get message from chan")
				continue
			}
			err := producer.publish(msg.RoutingKey, msg.Msg)
			if err != nil {
				producer.logger.Printf("[error] Push message failed: %s %s", msg.Id, err.Error())
				continue
			}
			if confirmed := <-producer.notifyConfirm; confirmed.Ack {
				if producer.confirmHandler != nil {
					producer.confirmHandler(msg.Id)
					producer.logger.Printf("[info] Push message confirmed: %s", msg.Id)
				}
			}
		}

	}
}

func (producer *Producer) publish(routingKey string, msg string) error {
	if !producer.isConnected {
		return errNotConnected
	}
	producer.logger.Println("publish mq = " + routingKey + ": " + msg)
	return producer.channel.Publish(
		producer.topic, // Exchange
		routingKey,     // Routing key
		false,          // Mandatory
		false,          // Immediate
		amqp.Publishing{
			DeliveryMode: 2,
			ContentType:  "text/plain",
			Body:         []byte(msg),
			Timestamp:    time.Now(),
		},
	)
}

// 关闭连接/信道
func (producer *Producer) Close() error {
	if !producer.isConnected {
		return errAlreadyClosed
	}
	err := producer.channel.Close()
	if err != nil {
		return err
	}
	err = producer.connection.Close()
	if err != nil {
		return err
	}
	close(producer.doneConn)
	close(producer.donePublish)
	producer.isConnected = false
	return nil
}
