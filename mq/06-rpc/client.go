package main

import (
	"context"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

func failOnErr(err error, msg string) {
	if err != nil {
		log.Fatalln(err, " : ", msg)
	}
}

func randString(l int) string {
	bytes := make([]byte, l)
	for i := 0; i < l; i++ {
		bytes[i] = byte(randInt(65, 90))
	}
	return string(bytes)
}

func randInt(min int, max int) int {
	return min + rand.Intn(max-min)
}

func main() {
	n := bodyFrom(os.Args)

	log.Printf(" [x] Requesting fib(%d)", n)
	res, err := rpc(n)
	failOnErr(err, "Failed to handle RPC request")
	log.Println(" [.] Got ", res)
}

func rpc(n int) (res int, err error) {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnErr(err, "conn error")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnErr(err, "channel error")
	defer ch.Close()

	//定义随机队列 用于接受CorrelationId
	q, err := ch.QueueDeclare("", false, false, true, false, nil)
	failOnErr(err, "queue error")

	msgs, err := ch.Consume(q.Name, "", true, false, false, false, nil)
	failOnErr(err, "consume error")

	corrId := randString(32)

	err = ch.PublishWithContext(
		context.Background(),
		"",
		"rpc_queue",
		false, false,
		amqp.Publishing{
			ReplyTo:       q.Name,
			CorrelationId: corrId,
			Body:          []byte(strconv.Itoa(n)),
			ContentType:   "text/plain",
		},
	)
	failOnErr(err, "publish error")

	for msg := range msgs {
		if corrId == msg.CorrelationId {
			res, err = strconv.Atoi(string(msg.Body))
			failOnErr(err, "strconv.Atoi error")
			return res, nil
		}
	}
	return -1, err
}

func bodyFrom(args []string) int {
	var s string
	if (len(args) < 2) || os.Args[1] == "" {
		s = "30"
	} else {
		s = strings.Join(args[1:], " ")
	}
	n, err := strconv.Atoi(s)
	failOnErr(err, "Failed to convert arg to integer")
	return n
}
