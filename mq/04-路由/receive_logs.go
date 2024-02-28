package main

import (
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
	"os"
)

func failOnErr(err error, msg string) {
	if err != nil {
		log.Println(msg)
	}
}

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnErr(err, "RMQ连接失败")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnErr(err, "声明Channel失败")
	defer conn.Close()

	//err = ch.ExchangeDeclare(
	//	"logs",
	//	"fanout",
	//	true,
	//	false,
	//	false,
	//	false,
	//	nil,
	//)
	//failOnErr(err, "声明Exchange失败")

	q, err := ch.QueueDeclare(
		"",
		false,
		false,
		true,
		false,
		nil,
	)
	failOnErr(err, "声明队列失败")

	if len(os.Args) < 2 {
		log.Printf("Usage: %s [info] [warning] [error]", os.Args[0])
		os.Exit(0)
	}

	for _, s := range os.Args[1:] {
		log.Printf("Binding queue %s to exchange %s with routing key %s",
			q.Name, "logs_direct", s)
		err = ch.QueueBind(
			q.Name,        // queue name
			s,             // routing key
			"logs_direct", // exchange
			false,
			nil)
		failOnErr(err, "Failed to bind a queue")
	}

	msgs, err := ch.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	failOnErr(err, "消费者注册失败")

	var forever chan struct{}
	go func() {
		for msg := range msgs {
			log.Printf(" [x] %s", msg.Body)
		}
	}()
	log.Printf(" [*] Waiting for logs. To exit press CTRL+C")

	<-forever

}
