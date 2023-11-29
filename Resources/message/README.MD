## Message

In Kafka, messages are the unit of data exchanged between producers and consumers.

- It is a key-value pair with an optional timestamp and optional metadata.
- The key and value are both byte arrays. 
- The key is used to decide which partition to send the message to and the value is the actual message data.
- Messages are persisted and durable within the Kafka cluster until the retention period configured for the topic has expired.

![img](https://developers.redhat.com/sites/default/files/styles/article_floated/public/RHOSAK%20LP1%20Fig6.webp?itok=iqmUWCwt)
