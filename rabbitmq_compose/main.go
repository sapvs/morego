package main

import (
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

//docker run -d --hostname my-rabbit --name some-rabbit -p 15672:15672 -p 5672:5672 rabbitmq:3-management-apline
func main() {
	fmt.Println("Starting rabbit mq client")

	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672")
	if err != nil {
		log.Fatal("Could not connect to amqp")
	}

	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatal("Could not get channel ")
	}
	defer ch.Close()

	que, err := ch.QueueDeclare("sapan", false, false, false, false, nil)
	if err != nil {
		fmt.Printf("Could not create queu %v", err)
		panic(err)
	}

	fmt.Println(que)

	// publishg
	err = ch.Publish(
		"",
		"sapan",
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte("Helllo Worl"),
		})

	if err != nil {
		fmt.Printf("Could not publish message %v", err)
		panic(err)
	}

	fmt.Println("Done with connection")
}
