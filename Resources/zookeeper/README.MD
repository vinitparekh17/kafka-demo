
## Zookeeper

As we know that Kafka is a distributed system, it needs to coordinate with multiple nodes. Zookeeper is a centralized service for maintaining configuration information, naming, providing distributed synchronization, and providing group services. All of these kinds of services are used in some form or another by distributed applications.

- Zookeeper tracks status of the Kafka cluster nodes and it also keeps track of Kafka topics, partitions etc.
- Zookeeper sends notifications to Kafka in case of any changes such as broker dies, comes up, new topic created, broker goes down, deletes a topic, etc.
- Zookeeper is also used for leader election for partitions/topics. For example, if a topic has 3 partitions and 3 brokers, then one broker will be the leader for each partition and Zookeeper is used for leader election.
- When zookeeper goes down, Kafka cluster will continue to work but no new rebalances will occur, which means that if a node goes down, it won’t be reflected in the cluster until zookeeper comes back up.

![img](https://blog.cloudxlab.com/wp-content/uploads/2021/04/Screenshot-21.png)

I know the following example which I am going to say does not corelates in the context of Kafka but zookeeper's work is as same as kube-apiserver in kubernetes. Both are the centralized service to maintain the cluster state. Very important to maintain the distributed system.

It's very important to understand the zookeeper's role in Kafka. I have seen many people who are not aware of zookeeper's role in Kafka. They think that zookeeper is just a dependency of Kafka. But it's not true. Zookeeper is the heart of Kafka. Without zookeeper, Kafka cannot work.

One more thing to notice that zookeeper is not a part of Kafka cluster. It is a separate service. It is
not only used in Kafka but also used in many other distributed systems such as Hadoop, HBase, Solr, etc.