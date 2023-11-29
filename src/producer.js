const { Kafka, Partitioners } = require('kafkajs');
const msg = process.argv[2];
run();
async function run() {
    try {
        const kafka = new Kafka({
            "clientId": "apple",
            "brokers": ["localhost:9092"],
        })

        const producer = kafka.producer({ createPartitioner: Partitioners.LegacyPartitioner })
        console.log("connecting...");
        await producer.connect()
        console.log("connection successed!");

        const partition = msg[0] > "f" ? 0 : 1

        const res = await producer.send({
            "topic": "MyOrders",
            "messages": [{
                "key": "demo",
                "value": msg,
                "partition": partition
            }]
        })
        console.log("Message has been published... " + JSON.stringify(res));
        await producer.disconnect();
    } catch (error) {
        console.error(`Something went wrong ${error}`);
    } finally {
        process.exit(0)
    }
}