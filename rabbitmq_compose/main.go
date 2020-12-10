package main

import (
	"context"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gorilla/mux"
	"github.com/streadway/amqp"
	"github.com/sudosapan/rabbitmq_compose/rabbitmq/infra"
)

const myqueue string = "MYQUEUE"

var myserver http.Server
var done = make(chan bool)

func rootpath(w http.ResponseWriter, r *http.Request) {
	log.Print("Received request at /")
	log.Print("Seever request at /")
}

func publish(w http.ResponseWriter, r *http.Request) {
	log.Printf("Body is  %v", r.Body)
	if r.Body != nil {
		defer r.Body.Close()
	}

	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Coudl not reads body %v", err)
		return
	}
	message := string(data)
	log.Printf("MEssage received %s", message)
	if err = infra.PostMessage("", myqueue, message); err != nil {
		log.Printf("Coudl not send mesasge %v", err)
	}

}

var channel *amqp.Channel

func _init() {
	go signalHandler()
	// COnnect and declare queues and channels.
	err := infra.ConnectRabbit(&infra.RabbitConnectionProperties{Password: "guest", User: "guest", RabbitHost: "localhost", Port: 5672})
	if err != nil {
		log.Fatalf("Could not connect to rabbit due to %v", err)
	}
	err = infra.DeclareQueue(myqueue)
	if err != nil {
		log.Fatalf("Could not declare queue %v", err)
	}
}

//docker run -d --hostname my-rabbit --name some-rabbit -p 15672:15672 -p 5672:5672 rabbitmq:3-management-apline
func main1() {
	log.Print("Starting rabbit mq client")

	myrouter := mux.NewRouter()
	myrouter.HandleFunc("/publish", publish).Methods(http.MethodPost)
	myrouter.HandleFunc("/", rootpath).Methods(http.MethodGet)

	myserver = http.Server{Addr: ":8080", Handler: myrouter}
	if err := myserver.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatalf("Failed to start http due to %v", err)
	} else {
		log.Printf("HTTP stopped by shutdown %v", err)
	}

	<-done
	log.Print("Done main ")

}

func signalHandler() {
	log.Print("Starting signal listener")
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGTERM, syscall.SIGINT)
	sig := <-sigs
	log.Printf("Recieved %v signal", sig)
	log.Print("Closing rabbit resources ")
	if err := infra.CloseRabbitResources(); err != nil {
		log.Printf("Could not close rabbit due to %v ", err)
	}

	log.Print("Closing http resources ")
	if err := myserver.Shutdown(context.Background()); err != nil {
		log.Printf("Could not close http resources due to %v ", err)
	}

	log.Print("Closed resources ")
	done <- true

}
