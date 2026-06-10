# Realtime Pubsub

Both message brokers and message streams require consumers to **"pull"** the messages out.

### Advantage:
* **Consumers pull at their own pace:** Consumers can retrieve messages when they are ready.
* **Consumers do not get overwhelmed:** A slow consumer will not be flooded with more messages than it can handle.

### Disadvantage:
* **Consumption lag:** When there is high ingestion, there can be a delay (lag) before consumers pull and process the messages.

---

### What if we want low latency & zero lag?-> **Realtime PubSub**

> **Realtime Pubsub** makes things reactive instead of continuous polling.

Instead of consumers pulling the message, **the message is pushed to them**.

* **Example:** Redis PubSub

![](/diagrams/pubsub.png)

---

* **This way we get really fast delivery time**
* **But it can overwhelm the consumers**
  * *What if consumers receive messages faster than they could process?*

### Practical Usecase:
1. **Message Broadcast**
2. **Configuration Push**
   * *All servers receive updates without polling for data.*

