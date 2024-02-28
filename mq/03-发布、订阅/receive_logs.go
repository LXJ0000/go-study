package main

import (
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
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

	err = ch.ExchangeDeclare(
		"logs",
		"fanout",
		true,
		false,
		false,
		false,
		nil,
	)
	failOnErr(err, "声明Exchange失败")

	q, err := ch.QueueDeclare(
		"", // 随机key 随机定义一个队列 因为生产者没有定义队列 而是交换机 因此不需要特定队列
		false,
		false,
		true,
		false,
		nil,
	)
	failOnErr(err, "声明队列失败")

	err = ch.QueueBind(
		q.Name,
		"",
		"logs",
		false,
		nil,
	)
	failOnErr(err, "队列绑定失败")

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
