console.log(`
    Welcome to the KafkaJS Demo!
    =============================
    This demo uses the following:
    - NodeJS
    - KafkaJS
    - Docker
    - Docker Compose
    - Kafka
    - Zookeeper
    
    To get started, run the following commands:
    $ docker-compose up -d
    $ npm install
    $ npm run create-topic
    $ npm run producer <message>
    $ npm run consumer
    
    To stop the demo, run the following command:
    $ docker-compose down
    `)