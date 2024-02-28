package main

import (
	"context"
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
	"os"
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

	err = ch.ExchangeDeclare(
		"logs_direct", //
		"direct",      // 定向发消息
		true,
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "声明Exchange失败")

	bodys := []string{"msg1.", "msg2.."}
	for _, body := range bodys {
		log.Printf(" [x] Sent %s", body)
		err = ch.PublishWithContext(
			context.Background(),
			"logs_direct",
			severityFrom(os.Args), // 获取key 也就是接受的类型 要接受什么数据 info error
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

func severityFrom(args []string) string {
	var s string
	if (len(args) < 2) || os.Args[1] == "" {
		s = "info"
	} else {
		s = os.Args[1]
	}
	return s
}
