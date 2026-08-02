## Delegation
Adding basic analytics to our blog.

`What does not need to be done in realtime(synchronous) should not be done in realtime.`
Do bare minimum in synchronous request context and take everything else asynchronous.


Core idea: Delegate and respond

![](/diagrams/messagebroker.png)
---

### Tasks that one delegated

outlash your request timeout

- long running tasks ( spin up EC2 / video encoding )
- heavy computation query
- batch & write
- anything that could be eventual

### Brokers
![](/diagrams/brokers.png)

---

### 2 common implementations
1. Message Queues
SQS, RabbitMQ
![](/diagrams/messagequeues1.png)
You enqueue the message in the broker and you have homogenous consumers. Each consumer in the broker is equally capable of handling the request. 
Consumers pull the message from the broker, they execute and delete the message.
---

2. Message Streams
Kafka, Kinesis
![](/diagrams/messagestreams1.png)

---

![](/diagrams/blogsincrementasync.png)
With asynchronous processing we have to be okay with lag.
Similarly we publish and consume ON_DELETE() event (total_blogs--)
But the worker here is also indexing the blogs on es. What if one work fails.
In that case there will be inconsistency. 
Ideally you want same message to be consumed by 2 different use cases. This is where message streams come into place.

In case of kafka, you have a broker in which messages are there. When you put a message they remain in the broker. You have consumer groups. The groups can iterate over the same messages at their own pace. The messages are deleted by deletion policy. 
![](/diagrams/kafkasearchandcounter.png)

`
---

## Kafka Essentials
Kafka is a message stream that holds the messages.
Internally kafka has a topic. One kafka cluster can have many topics.
Each kafka topic has multiple partitions. Message is sent to a topic and depending on the configured hash key it is put into partition.
Kafka can guarantee ordering within the partition.
How does kafka know message is processed -> Kafka gives u a commit call, every consumer when it read a message, it needs to call commit, help kafka register that until this point the messages have been read and processed.
If consumer dies and starts again, it can start from the last checkpoint. You have to decide when to commit -> 1. Before processing the message, but while processing your machine crashed. According to kafka you have consumed the message but in reality it has failed. You have to be okay with this. 2. When commiting after processing you have to be okay with message being sent again. (processing same message again). We have to deal with the idempotency here.

### Limitations of Kafka
Number of consumers in Kafka is limited by the number of partitions in the topic. 
![](/diagrams/kafkapartitions.png)
Your parellelism of kafka is limited by number of partitions in the topic. *You cannot reduce the number of partitions once you have increased it.*

### Example
- Imagine you are sending appointment reminders. It was to be sent at 4, you enqueued the message, lot of messages piled up. The appointment reminder which should have been sent at 3:55 is now sent at 4:15 which is a problem. Here business sensitivity comes in picture. The lag might be unacceptable in certain types of use cases. 