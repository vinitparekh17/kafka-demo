const { Kafka } = require('kafkajs');
run();
async function run() {
    try {
        const kafka = new Kafka({
            "clientId": "apple12",
            "brokers": ["localhost:9092"]
        })

        const admin = kafka.admin()
        console.log("connecting...");
        await admin.connect()
        console.log("connection successed!");

        const res = await admin.createTopics({
            topics: [{
                "topic": "MyOrders",
                "numPartitions": 2
            }]
        })
        console.log(`user topic is creation ${res ? "done" : "failed"}!`);
        await admin.disconnect();
    } catch (error) {
        console.error(error);
    } finally {
        process.exit(0)
    }
}