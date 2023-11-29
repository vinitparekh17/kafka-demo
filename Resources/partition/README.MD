## Partitions

> A partition is a unit of parallelism in Kafka. 

Partitions allows to parallelize a topic by splitting the data in a particular topic across multiple brokers — each partition can be placed on a separate machine to allow for multiple consumers to read from a topic in parallel.

- Each partition is an ordered, immutable sequence of records that is continually appended to—a structured commit log.
- Partitions can be configured to have replication factor > 1 to provide fault tolerance.
- One of the brokers acts as a leader for a partition and the rest of the brokers act as followers. The leader handles all read and write requests for the partition while the followers passively replicate the leader. If the leader fails, one of the followers will

![img](https://developers.redhat.com/sites/default/files/styles/article_floated/public/RHOSAK%20%20Fig3.png?itok=Z3Z3Z3Z3)
<center> <b>Figure 3. </b> Partitioning a topic into multiple partitions. </center>

___

![img](https://developers.redhat.com/sites/default/files/styles/article_floated/public/RHOSAK%20LP1%20Fig4.png?itok=zGET7b7K)
<center> <b>Figure 4. </b> Each partition has a leader and zero or more followers. </center>
