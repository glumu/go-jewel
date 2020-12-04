package rabbitmq

import (
	"bytes"
	"errors"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/streadway/amqp"
)

type Producer struct {
	name            string
	logger          *log.Logger
	connection      *amqp.Connection
	channel         *amqp.Channel
	done            chan bool
	notifyConnClose chan *amqp.Error
	notifyChanClose chan *amqp.Error
	notifyConfirm   chan amqp.Confirmation
	isReady         bool

	addr        string
	topic       string
	recvHandler func(msg string)
	routingKeys map[string][]string
}

const (
	reconnectDelay = 5 * time.Second // 连接断开后多久重连
	reInitDelay    = 2 * time.Second // When setting up the channel after a channel exception
	resendDelay    = 5 * time.Second // 消息发送失败后，多久重发
	resendTime     = 3               // 消息重发次数
)

var (
	errNotConnected  = errors.New("not connected to the producer")
	errAlreadyClosed = errors.New("already closed: not connected to the producer")
	errShutdown      = errors.New("session is shutting down")
)

func NewProducer(username string, password string, host string, port string, vhost string,
	node string, topic string, routingKeys map[string][]string) *Producer {
	addr := fmt.Sprintf("amqp://%s:%s@%s:%s/", username, password, host, port)
	if vhost != "" {
		addr = addr + vhost
	}

	session := Producer{
		logger:      log.New(os.Stdout, "", log.LstdFlags),
		name:        node,
		addr:        addr,
		topic:       topic,
		done:        make(chan bool),
		routingKeys: routingKeys,
	}
	go session.handleReconnect(addr)
	return &session
}

func (s *Producer) Start(recvHandler func(msg string)) {
	s.recvHandler = recvHandler
}

// 如果连接失败会不断重连
// 如果连接断开会重新连接
func (s *Producer) handleReconnect(addr string) {
	for {
		s.isReady = false
		log.Println("Attempting to connect")

		conn, err := s.connect(addr)
		if err != nil {
			log.Println("Failed to connect. Retrying...")

			select {
			case <-s.done:
				s.logger.Println("Session connect exception!")
				return
			case <-time.After(reconnectDelay):
				s.logger.Println("Do session reconnect!")
			}
			continue
		}

		if done := s.handleReInit(conn); done {
			break
		}
	}
}

// 连接rabbitmq，以生产者的name定义一个队列
// connect will create a new AMQP connection
func (s *Producer) connect(addr string) (*amqp.Connection, error) {
	conn, err := amqp.Dial(addr)

	if err != nil {
		return nil, err
	}

	s.changeConnection(conn)
	log.Println("Connected!")
	return conn, nil
}

// handleReconnect will wait for a channel error
// and then continuously attempt to re-initialize both channels
func (s *Producer) handleReInit(conn *amqp.Connection) bool {
	for {
		s.isReady = false

		err := s.init(conn)

		if err != nil {
			log.Println("Failed to initialize channel. Retrying...")

			select {
			case <-s.done:
				s.logger.Println("Session init exception!")
				return true
			case <-time.After(reInitDelay):
				s.logger.Println("Session reinit!")
			}
			continue
		}

		select {
		case <-s.done:
			return true
		case <-s.notifyConnClose:
			log.Println("Connection closed. Reconnecting...")
			return false
		case <-s.notifyChanClose:
			log.Println("Channel closed. Re-running init...")
		}
	}
}

// init will initialize channel & declare queue
func (s *Producer) init(conn *amqp.Connection) error {
	ch, err := conn.Channel()

	if err != nil {
		return err
	}

	err = ch.Confirm(false)

	if err != nil {
		return err
	}

	_, err = ch.QueueDeclare(
		s.name,
		true,  // Durable
		false, // Delete when unused
		false, // Exclusive
		false, // No-wait
		nil,   // Arguments
	)

	if err != nil {
		return err
	}

	err = ch.ExchangeDeclare(
		s.topic,
		"topic",
		true,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		return err
	}

	s.changeChannel(ch)
	s.bindQueue()
	s.receive()
	s.isReady = true
	log.Println("Setup init!")

	return nil
}

// 监听Rabbit channel的状态
// changeConnection takes a new connection to the queue,
// and updates the close listener to reflect this.
func (s *Producer) changeConnection(connection *amqp.Connection) {
	s.connection = connection
	s.notifyConnClose = make(chan *amqp.Error)
	s.connection.NotifyClose(s.notifyConnClose)

	//producer.connection = connection
	//producer.channel = channel
	//// channels没有必要主动关闭。如果没有协程使用它，它会被垃圾收集器收拾
	//producer.notifyClose = make(chan *amqp.Error)
	//producer.notifyConfirm = make(chan amqp.Confirmation)
	//producer.channel.NotifyClose(producer.notifyClose)
	//producer.channel.NotifyPublish(producer.notifyConfirm)
}

// changeChannel takes a new channel to the queue,
// and updates the channel listeners to reflect this.
func (s *Producer) changeChannel(channel *amqp.Channel) {
	s.channel = channel
	s.notifyChanClose = make(chan *amqp.Error)
	s.notifyConfirm = make(chan amqp.Confirmation, 1)
	s.channel.NotifyClose(s.notifyChanClose)
	s.channel.NotifyPublish(s.notifyConfirm)
}

// 三次重传的发消息
func (s *Producer) Publish(routingKey string, msg string) error {
	fmt.Println("publish mq = " + routingKey + ": " + string(msg))
	if !s.isReady {
		s.logger.Println("Session not connected when publish!")
		return errors.New("failed to push: not connected")
	}

	var currentTime = 0
	for {
		s.logger.Println("Push start!")
		err := s.UnsafePush(routingKey, []byte(msg))
		s.logger.Println("Push end!")

		if err != nil {
			s.logger.Println("Push failed. Retrying...")
			currentTime += 1
			if currentTime < resendTime {
				continue
			} else {
				return err
			}
		}

		ticker := time.NewTicker(resendDelay)
		select {
		case confirm := <-s.notifyConfirm:
			if confirm.Ack {
				s.logger.Println("Push confirmed!")
				return nil
			} else {
				s.logger.Println("Push confirm not Ack!")
			}
		case <-ticker.C:
			s.logger.Println("Push ticker!")
		}
		s.logger.Println("Push didn'Failed to connect. Retryingt confirm. Retrying...")
	}
}

// 发送出去，不管是否接受的到
func (s *Producer) UnsafePush(routingKey string, data []byte) error {
	if !s.isReady {
		return errNotConnected
	}
	return s.channel.Publish(
		s.topic,    // Exchange
		routingKey, // Routing key
		false,      // Mandatory
		false,      // Immediate
		amqp.Publishing{
			DeliveryMode: 1,
			ContentType:  "text/plain",
			Body:         data,
			Timestamp:    time.Now(),
		},
	)
}

// Stream will continuously put queue items on the channel.
// It is required to call delivery.Ack when it has been
// successfully processed, or delivery.Nack when it fails.
// Ignoring this will cause data to build up on the server.
func (s *Producer) Stream() (<-chan amqp.Delivery, error) {
	if !s.isReady {
		return nil, errNotConnected
	}
	return s.channel.Consume(
		s.name,
		"",    // Consumer
		false, // Auto-Ack
		false, // Exclusive
		false, // No-local
		false, // No-Wait
		nil,   // Args
	)
}

// 关闭连接/信道
func (s *Producer) Close() error {
	if !s.isReady {
		return errAlreadyClosed
	}
	err := s.channel.Close()
	if err != nil {
		return err
	}
	err = s.connection.Close()
	if err != nil {
		return err
	}
	close(s.done)
	s.isReady = false
	return nil
}

func (s *Producer) bindQueue() (err error) {
	for topic, keys := range s.routingKeys {
		for _, key := range keys {
			err = s.channel.QueueBind(s.name, key, topic, false, nil)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (s *Producer) receive() (err error) {
	s.channel.Qos(1, 0, true)
	consumerTag := fmt.Sprintf("amq.ctag-%s", getRandomString(22))
	s.logger.Println("Consumer tag: %s", consumerTag)
	msgs, err := s.channel.Consume(
		s.name,
		consumerTag,
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}
	go func() {
		defer func() {
			if rec := recover(); rec != nil {
				err := fmt.Errorf("%v", rec).Error()
				s.logger.Println("Consumer panic: %s", err)
				panic(err)
			}
		}()
		for d := range msgs {
			msg := bytesToString(&(d.Body))
			if s.recvHandler != nil {
				s.recvHandler(*msg)
			}
			d.Ack(false)
		}
	}()
	return nil
}

func bytesToString(b *[]byte) *string {
	s := bytes.NewBuffer(*b)
	r := s.String()
	return &r
}

func getRandomString(l int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyz"
	_bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < l; i++ {
		result = append(result, _bytes[r.Intn(len(_bytes))])
	}
	return string(result)
}
