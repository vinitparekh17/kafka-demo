package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	handler "github.com/vinitparekh17/kafka-demo/Handler"
)

func main() {
	fmt.Println("Creating consumer...")
	c, err := kafka.NewConsumer(&kafka.ConfigMap{"bootstrap.servers": "localhost:9092", "group.id": "grp1"})
	handler.CheckError(err, "Failed to create consumer...")
	fmt.Println("Consumer has been created...")

	topic := "Myorders"
	err = c.SubscribeTopics([]string{topic}, nil)
	handler.CheckError(err, "Failed to subscribe the topic...")
	fmt.Printf("Consumer has been subscribed to %s\n", topic)

	// channel to notify ctrl+c etc. terminal activity to explicitly stop
	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)
	run := true

	for run {
		select {
		case sig := <-sigchan:
			fmt.Printf("\nCaught signal %v: terminating\n", sig)
			run = false
		default:
			ev := c.Poll(1000)
			switch e := ev.(type) {
			case *kafka.Message:
				fmt.Printf("\n\nTopic: %s\nConsumed msg: %s\nPartition number: %d", *e.TopicPartition.Topic, string(e.Value), e.TopicPartition.Partition)
			case kafka.Error:
				handler.CheckError(err, "Kafka error while consuming messages...")
				run = false
			}
		}
	}
	defer c.Close()
}
