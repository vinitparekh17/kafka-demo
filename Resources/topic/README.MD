## Topics

Topics are the core abstraction in Kafka. Topics is the name of the stream to which records are published. For example, a topic named `orders` can be used to publish orders from an e-commerce website. To be more specific, 
> A topic is a category or feed name to which records are published.

- Topics are split into partitions.
- Each partition is ordered and immutable. This means that records are appended to a partition in a strictly sequential order and once written, they cannot be modified.
- Each record in a partition is assigned a sequential id number called the offset that uniquely identifies each record within the partition.

![img](https://developers.redhat.com/sites/default/files/styles/article_floated/public/RHOSAK%20%20Fig2.png?itok=snniV09u)