package infra

import (
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/streadway/amqp"
)

var (
	rabbitConn *amqp.Connection
)

const (
	RETRIES    int           = 3
	RETRY_TIME time.Duration = 5
)

// TODO this is single channle for testing need to send it to client
var channel *amqp.Channel

// RabbitConnectionProperties properties
type RabbitConnectionProperties struct {
	User       string
	Password   string
	RabbitHost string
	Port       int
}

func (rcp *RabbitConnectionProperties) amqpURL() string {
	return fmt.Sprintf("amqp://%s:%s@%s:%d/", rcp.User, rcp.Password, rcp.RabbitHost, rcp.Port)
}

// ConnectRabbit connection to rabbit mq as per details provided rcp
// In container the liveness and readiness check should work so no need to
func ConnectRabbit(rcp *RabbitConnectionProperties) error {
	log.Print("Connecting to rabbit")
	var try int = 1
	var err error
	for try <= RETRIES {
		log.Printf("Connecting rabbit %d time", try)
		log.Print("Connecting to rabbit not connected yet")
		rabbitConn, err = amqp.Dial(rcp.amqpURL())
		if err == nil {
			log.Print("Is connected breaking")
			break
		}
		try++
		log.Print("Sleeping ")
		time.Sleep(RETRY_TIME * time.Second)
	}

	if err != nil {
		return err
	}

	log.Print("Connected to rabbit, opening channel")

	channel, err = rabbitConn.Channel()
	if err != nil {
		return err
	}

	log.Print("Channel opened")

	return nil
}

// DeclareQueue creates a queue with queueName
func DeclareQueue(queueName string) error {
	_, err := channel.QueueDeclare(queueName, false, false, false, false, nil)
	return err
}

// PostMessage sends message to rabbit
func PostMessage(exchange string, queue string, message string) error {
	if rabbitConn == nil || rabbitConn.IsClosed() {
		return errors.New("no connection to tabbit")
	}
	if err := channel.Publish(exchange, queue, false, false, amqp.Publishing{ContentType: "text/plain", Body: []byte(message)}); err != nil {
		return err
	}
	return nil
}

// CloseRabbitResources closes the rabbitConn if not already closed
func CloseRabbitResources() error {
	log.Print("Closing resources")
	err := channel.Close()
	if err != nil {
		return err
	}

	log.Print("Channel closed")

	if !rabbitConn.IsClosed() {
		return rabbitConn.Close()
	}

	log.Print("Connection closed")

	return nil
}
