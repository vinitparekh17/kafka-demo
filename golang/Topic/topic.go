package main

import (
	"context"
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	handler "github.com/vinitparekh17/kafka-demo/Handler"
)

func main() {
	// create a kafka topic
	topic := "Myorders"

	// create a kafka admin client
	adminClient, err := kafka.NewAdminClient(&kafka.ConfigMap{"bootstrap.servers": "localhost:9092"})
	handler.CheckError(err, "Failed to create admin client...")
	fmt.Println("Admin client has been created...")
	defer adminClient.Close()
	defer fmt.Println("Admin client has been closed...")
	// create a topic
	topicSpec := kafka.TopicSpecification{
		Topic:             topic,
		NumPartitions:     2,
		ReplicationFactor: 1,
	}
	topics := []kafka.TopicSpecification{topicSpec}
	fmt.Println("Creating topic...")
	results, err := adminClient.CreateTopics(context.Background(), topics)
	handler.CheckError(err, "Failed to create topic...")
	fmt.Printf("Topic %s has been created with results %v\n", topic, results)
}
