package main

import (
	"context"
	"fmt"

	kafka "github.com/segmentio/kafka-go"
)

func main() {
	// make a new reader that consumes from topic-A, partikktion 0, at offset 42
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  []string{"localhost:9092"},
		GroupID:  "consumer-group-id",
		Topic:    "topic-A",
		MinBytes: 10e3, // 10KB
		MaxBytes: 10e6, // 10MB
	})
	fmt.Printf("---begin consumer---\n")
	for {
		m, err := r.ReadMessage(context.Background())
		if err != nil {
			fmt.Printf("---error---:%s \n", err)
			break
		}

		fmt.Printf("message at offset %d: %s = %s\n", m.Offset, string(m.Key), string(m.Value))
	}
	fmt.Printf("---end consumer---\n")
	r.Close()
}
