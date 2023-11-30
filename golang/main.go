package main

import (
	"fmt"
)

func main() {
	fmt.Println(`
    Welcome to the KafkaJS Demo!
    =============================
	
    This demo uses the following:
    - Golang
    - confluent-kafka-go
    - Docker
    - Docker Compose
    - Kafka
    - Zookeeper
    
    To get started, run the following commands:
    $ docker-compose up -d
	$ go run main.go
	$ go run ./Producer/producer.go
	$ go run ./Consumer/consumer.go
	$ go run ./Topic/topic.go
    
    To stop the demo, run the following command:
    $ docker-compose down
	`)
}
