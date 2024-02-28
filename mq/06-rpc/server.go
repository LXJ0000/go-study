package main

import (
	"context"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
	"strconv"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalln(err, " : ", msg)
	}
}

func fib(n int) int {

	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	a, b := 1, 1
	for i := 2; i <= n; i++ {
		a, b = b, a+b
	}
	return a
}

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "conn error")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "channel error")
	defer ch.Close()

	//定义一个队列 用于接受参数 n
	q, err := ch.QueueDeclare("rpc_queue", false, false, false, false, nil)
	failOnError(err, "queue error")

	err = ch.Qos(1, 0, false)

	msgs, err := ch.Consume(q.Name, "", false, false, false, false, nil)
	failOnError(err, "consume error")

	var forever chan struct{}

	go func() {
		for msg := range msgs { // 消费参数 n
			n, err := strconv.Atoi(string(msg.Body))
			failOnError(err, "strconv.Atoi fail")

			log.Printf(" [.] fib(%d)", n)
			resp := fib(n)

			//消费完成 返回结果到msg指定的队列 并携带CorrelationId让对方知道是哪个被消费
			err = ch.PublishWithContext(
				context.Background(),
				"",
				msg.ReplyTo,
				false,
				false,
				amqp.Publishing{
					CorrelationId: msg.CorrelationId,
					Body:          []byte(strconv.Itoa(resp)),
					ContentType:   "text/plain",
				},
			)
			failOnError(err, "Failed to publish a message")

			msg.Ack(false) // 只确认当前处理的消息 使用 msg.Ack(true) 可以批量确认消息
		}
	}()

	log.Printf(" [*] Awaiting RPC requests")
	<-forever
}
