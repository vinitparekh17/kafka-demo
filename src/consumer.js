const { Kafka, Partitioners } = require('kafkajs');
run();
async function run() {
    try {
        const kafka = new Kafka({
            "clientId": "apple",
            "brokers": ["localhost:9092"]
        })

        const consumer = kafka.consumer({ "groupId": "demo" })
        console.log("connecting...");
        await consumer.connect()
        console.log("connection successed!");

        await consumer.subscribe({ "topic": "MyOrders", "fromBeginning": true })
        await consumer.run({
            "eachMessage": async ({ message, partition, topic }) => {
                console.table({ topic, partition, "value": message.value.toString() })
            }
        })
    } catch (error) {
        console.error(`Something went wrong ${error}`);
    }
}