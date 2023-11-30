package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	handler "github.com/vinitparekh17/kafka-demo/Handler"
)

func main() {
	var msg []string = os.Args[1:]
	byteMsg := []byte(strings.Join(msg, " "))

	fmt.Println("Trying to connect the producer...")
	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": "localhost:9092"})
	handler.CheckError(err, "Failed to create a producer")
	fmt.Println("Producer has been connected..." + p.String())

	topic := "Myorders"
	deliverychan := make(chan kafka.Event, 10000)
	partitionNum := rand.Int31n(2) // random partition between 0 and 1
	println("Partition number: ", partitionNum)
	err = p.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: partitionNum},
		Value:          byteMsg},
		deliverychan)

	handler.CheckError(err, "Something went wrong...")

	e := <-deliverychan
	m := e.(*kafka.Message)

	if m.TopicPartition.Error != nil {
		fmt.Printf("Delivery failed: %v\n", m.TopicPartition.Error)
	} else {
		fmt.Printf("Delivered message to topic %s [%d] at offset %v\n",
			*m.TopicPartition.Topic, m.TopicPartition.Partition, m.TopicPartition.Offset)
	}
	close(deliverychan)
}
