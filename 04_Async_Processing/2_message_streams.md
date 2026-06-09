# Message Streams

* **Message streams are similar to message queues... with a few differences**
* **To understand, let's take an example:**

### Example: Medium.com
Say we are building Medium.com, and upon every blog published, we:
* **Need to index it in search engine (ElasticSearch)**
* **Need to do `count++` for user's total blogs**

#### Approach 1: We use one message broker and add logic in consumer
![](/diagrams/rabbitmq.png)

* **Consumers are doing two things: `count++` and index**
* **Issue: What if write to one succeeded but the other one failed?**
  * `count++` ✓ but not there in search
  * OR there in search but count does not match

#### Approach 2: Two brokers and two set of consumers
![](/diagrams/rabbitmq2.png)

* **API server writes to two Brokers & each has its own set of consumers**
* **This still does not solve the problem!**
  * When API server writes to two RabbitMQ, one of them fails.
  * We end up in the same spot.

---

* **Hence, we want "write to one" and "read by many" semantic**
* **This is where Message Streams come into the picture (e.g., Kafka, Kinesis)**

### Message Streams

* **Message streams are similar to message brokers with one change:**
  * **Multiple types of consumers read the same message**

#### Approach 3: Using streams and multiple types of consumers
![](/diagrams/kafka.png)

* **API server pushes *one* message in Kafka; search and counter services both read the message and do their work**

## Message Queues vs Message Streams

### Message queues
![](/diagrams/messagequeues.png)
Lets say we have 4 messages to send and we have three consumers doing the exact same thing. When you send a message it goes and sits in the front of the queue. 
Second comes after that and so on. 
When the consumer is running one of the consumer makes a call to the queue and pulls one message out. 
Consumer three read the second message and so on. It does not matter which consumer we are sending the message to. All are doing the same thing. 


### Message Streams
![](/diagrams/messagestreams.png)
Lets say we have the same 4 events. The 4 messages are published and now we have bunch of messages in the message stream. Now we have different types of consumers which are called consumer groups. We have search consumer group which can have multiple servers and the counter consumer group. We have 2 totally different use cases that need to process the exact same event. When a blog is published we put one event. We want the same event to be consumed by the search service and the counter service. In the message queues the consumer was pulling the message out. Here the consumers iterate over the messages consuming one at a time. Messages lie in kafka forever. Whichever consumer wants to spin up it can create a new consumer group and they can start consuming messages from the first message. They all iterate through the same message in the same order.
Consumer would still pull the message over here but visually the messages are still there in the queue. They are not deleted. The same set of messages in message streams are iterated by different types of consumers 



# Kafka
![](/diagrams/kafka1.png)

* **Kafka is a message stream that holds the messages**
* **Internally, Kafka has topics**
* **Every topic has 'n' partitions**
* **Message is sent to a topic and depending on the configured hash key it is put into a partition**
* **Within a partition, messages are ordered**
  * *No ordering guarantee across partitions*

  ## Limitations of kafka
  #consumers = #partitions
  ![](/diagrams/kafka2.png)
  One consumer can iterate through one partition only. this is for a consumer in a consumer group.
  Same message can be iterated by different consumer groups 

  When you are done reading some set of messages you can commit and kafka will store that for you.