## Consumer

The consumer is responsible for subscribing to Kafka topics and processing the messages. This section explains how to create a consumer and consume messages from Kafka using various programming languages and libraries.

- Consumers read messages from a topic in order and are identified by a consumer group name.
- Each consumer within the group reads from a single partition and that partition is exclusive to the consumer until it disconnects.
- Consumers can be configured to be either in a group or alone.