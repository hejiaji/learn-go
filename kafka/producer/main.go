package main

import (
	"context"
	"time"

	kafka "github.com/segmentio/kafka-go"
)

func main() {
	topic := "go-test"
	partition := 0
	println("-----")

	conn, _ := kafka.DialLeader(context.Background(), "tcp", "localhost:9092", topic, partition)
	println("-----")
	conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
	conn.WriteMessages(
		kafka.Message{Value: []byte("one!")},
		kafka.Message{Value: []byte("two!")},
		kafka.Message{Value: []byte("three!")},
	)

	conn.Close()
}
