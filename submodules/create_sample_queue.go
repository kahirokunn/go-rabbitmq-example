package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/streadway/amqp"
)

var (
	// RabbitMQのURLはパラメータで指定
	rabbitmqURL = flag.String("rabbitmqUrl", "localhost:5672", "Your RabbtMQ URL")
	queueName   = "SampleQueue"
)

func main() {
	flag.Parse()

	if *rabbitmqURL == "" {
		log.Fatalln("[ERROR] require rabbitmqUrl")
	}

	log.Println("consumer start")

	url := fmt.Sprintf("amqp://%s", *rabbitmqURL)

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)

	conn, err := amqp.Dial(url)
	if err != nil {
		log.Printf("[ERROR] %s", err.Error())
		return
	}
	defer conn.Close()

	// チャンネル開いて・・・
	ch, err := conn.Channel()
	if err != nil {
		log.Printf("[ERROR] %s", err.Error())
		return
	}
	defer ch.Close()

	// Exchangeを作って・・・
	if err := ch.ExchangeDeclare("sample", "direct", false, true, false, false, nil); err != nil {
		log.Printf("[ERROR] %s", err.Error())
		return
	}

	q, err := ch.QueueDeclare(queueName, true, false, false, false, nil)
	if err != nil {
		log.Printf("[ERROR] %s", err)
		return
	}

	log.Printf("Success to create new queue. [NAME] %s", q.Name)
}
