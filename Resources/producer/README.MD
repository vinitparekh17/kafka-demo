## Producer

The producer is responsible for publishing messages to Kafka topics. This section explains how to create a producer and send messages to Kafka using various programming languages and libraries.

- Producers automatically know to which broker and partition to send the message to.
- Producers can choose to receive acknowledgments of data writes.
- Producers can choose to send a key with the message. If a key is sent, then all messages with that key will arrive to the same partition.