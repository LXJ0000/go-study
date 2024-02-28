package main

import (
	"context"
	"log"
	"strings"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

func bodyFrom(args []string) (str string) {
	if len(args) < 2 {
		return "hello"
	} else {
		str = strings.Join(args[1:], "")
	}
	return
}

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"hello1", // name
		true,     // durable 声明为持久队列，实现消息持久化
		false,    // delete when unused
		false,    // exclusive
		false,    // no-wait
		nil,      // arguments
	)
	failOnError(err, "Failed to declare a queue")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	//body := bodyFrom(os.Args)
	msgs := []string{"msg1.", "msg2..", "msg3....................", "msg4..", "msg5.."}
	for _, msg := range msgs {
		err = ch.PublishWithContext(ctx,
			"",     // exchange
			q.Name, // routing key
			false,  // mandatory 强制
			false,  // immediate 立即
			amqp.Publishing{
				DeliveryMode: amqp.Persistent, // 持久（交付模式：瞬态 | 持久）
				ContentType:  "text/plain",
				Body:         []byte(msg),
			})
		failOnError(err, "Failed to publish a message")
		log.Printf(" [x] Sent %s\n", msg)
	}

}
