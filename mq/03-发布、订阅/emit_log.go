package main

import (
	"context"
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Println(msg)
	}
}

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "RMQ连接失败")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "声明Channel失败")
	defer conn.Close()
	// fanout 广播
	err = ch.ExchangeDeclare("logs", "fanout", true, false, false, false, nil)
	failOnError(err, "声明Exchange失败")

	bodys := []string{"msg1.", "msg2..", "msg3...", "msg4...."}
	for _, body := range bodys {
		log.Printf(" [x] Sent %s", body)
		err = ch.PublishWithContext(
			context.Background(),
			"logs",
			"", // routing key
			false,
			false,
			amqp.Publishing{
				ContentType: "text/plain",
				Body:        []byte(body),
			},
		)
		failOnError(err, fmt.Sprintf("消息：%s，发送失败", body))
	}

}
