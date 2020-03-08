package rabbitmq

import (
	"fmt"
	"log"
	"time"

	"github.com/streadway/amqp"
)

// Config : rabbitmq
type Config struct {
	Host     string
	Port     string
	User     string
	Password string
}

// RabbitMQ struct
type RabbitMQ struct {
	conf       *Config
	queueNames []RabbitMQName

	channel    *amqp.Channel
	connection *amqp.Connection

	deliveryChannel chan amqp.Delivery
	errorChannel    chan *amqp.Error

	close chan bool
}

// New returns new RabbitMQ ptr
func New(config *Config, qNames []RabbitMQName) *RabbitMQ {
	rbmq := new(RabbitMQ)
	rbmq.conf = config
	rbmq.queueNames = qNames

	rbmq.Connect()
	// go rbmq.Reconnect()

	return rbmq
}

// Connect connects to rabbitmq server
func (rbmq *RabbitMQ) Connect() {
	connURL := fmt.Sprintf("amqp://%s:%s@%s:%s/", rbmq.conf.User, rbmq.conf.Password, rbmq.conf.Host, rbmq.conf.Port)

	for {
		log.Printf("Connecting to rabbitmq on %s\n", connURL)
		conn, err := amqp.Dial(connURL)
		if err == nil {
			rbmq.deliveryChannel = make(chan amqp.Delivery)
			rbmq.errorChannel = make(chan *amqp.Error)
			rbmq.close = make(chan bool)

			rbmq.connection = conn
			rbmq.connection.NotifyClose(rbmq.errorChannel)

			log.Println("Connection established!")

			if err = rbmq.openChannel(); err != nil {
				logError("Opening channel failed", err)
				rbmq.Close()
				sleepDueToError(err)
				continue
			}

			if err = rbmq.declareQueue(); err != nil {
				logError("Queue declaration failed", err)
				rbmq.Close()
				sleepDueToError(err)
				continue
			}

			return
		}

		sleepDueToError(err)
	}
}

func sleepDueToError(err error) {
	logError("Connection to rabbitmq failed. Retrying in 1 sec... ", err)
	time.Sleep(1000 * time.Millisecond)
}

func (rbmq *RabbitMQ) openChannel() error {
	channel, err := rbmq.connection.Channel()
	if err != nil {
		return err
	}
	rbmq.channel = channel
	return nil
}

func (rbmq *RabbitMQ) declareQueue() error {
	for _, queueName := range rbmq.queueNames {
		_, err := rbmq.channel.QueueDeclare(
			string(queueName), // name
			true,              // durable
			false,             // delete when unused
			false,             // exclusive
			false,             // no-wait
			nil,               // arguments
		)
		if err != nil {
			return err
		}
	}
	return nil
}

// ReconnectSender : reconnects to rabbitmq server if connection is lost
func (rbmq *RabbitMQ) ReconnectSender() {
	for {
		select {
		case <-rbmq.close:
			return
		case err := <-rbmq.errorChannel:
			logError("Reconnecting after connection error", err)
			rbmq.Connect()
		}
	}
}

// ReconnectReceiver : reconnects to rabbitmq server if connection is lost
func (rbmq *RabbitMQ) ReconnectReceiver() {
	for {
		select {
		case <-rbmq.close:
			return
		case err := <-rbmq.errorChannel:
			logError("Reconnecting after connection error", err)
			rbmq.Connect()
			rbmq.ConsumingMessage()
		}
	}
}

// registerConsumingMessage : register channel consume messages for each queue
func (rbmq *RabbitMQ) registerConsumingMessage() error {
	for _, queueName := range rbmq.queueNames {
		msgs, err := rbmq.channel.Consume(
			string(queueName), // queue
			"",                // messageConsumer
			true,              // auto-ack
			false,             // exclusive
			false,             // no-local
			false,             // no-wait
			nil,               // args
		)
		if err != nil {
			logError("Consuming messages from queue failed", err)
			return err
		}
		go func() {
			for msg := range msgs {
				rbmq.deliveryChannel <- msg
			}
		}()
	}

	return nil
}

// ConsumingMessage : channel consume messages each queue
func (rbmq *RabbitMQ) ConsumingMessage() {
	err := rbmq.registerConsumingMessage()
	if err != nil {
		log.Printf("Consuming message failed : %v\n", err)
		return
	}

	for {
		select {
		case <-rbmq.close:
			return
		case delivery := <-rbmq.deliveryChannel:
			log.Printf("Received message : %s", string(delivery.Body[:]))
		}
	}
}

// SendMessage : channel publish messages each queue
func (rbmq *RabbitMQ) SendMessage(message string) {
	for _, qName := range rbmq.queueNames {
		err := rbmq.channel.Publish(
			"",            // exchange
			string(qName), // routing key
			false,         // mandatory
			false,         // immediate
			amqp.Publishing{
				ContentType: "text/plain",
				Body:        []byte(message),
			})
		logError("Sending message to queue failed", err)
	}
}

// Close : closes connection and channel
func (rbmq *RabbitMQ) Close() {
	log.Println("Closing connection")

	rbmq.close <- true

	if rbmq.channel != nil {
		rbmq.channel.Close()
		rbmq.channel = nil
	}

	if rbmq.connection != nil {
		rbmq.connection.Close()
		rbmq.connection = nil
	}
}
