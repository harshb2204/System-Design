# Designing a Notification Service

Design a notification service that sends notifications to users across channels.
The system needs to be horizontally scalable and should support a very high fan-out.

Instead of directly jumping to send "millions" of notifications, lets start simple.

## Notification Template

We need a UI and simple backend to create notification templates that will be configured by some internal team.

Number of notification templates will not be huge and will fit on a single machine. Hence, we start with relational DB.
![](/diagrams/notificationtemplate.png)

When any team/product wants to send notification:
1. create a template using the Notification Control Service
2. specify the variables
3. note the ID of the defined template

## Notification Channels

User can be notified via multiple channels:
1. email
2. android push notn
3. apple push notn
4. SMS

For each of these channels there are providers that expose APIs. We invoke them programmatically to send notification at the very instant.

Some examples: SES, mailgun, OneSignal, Twilio, msg91, etc.

So, the servers that will send/trigger the actual notification will have to configure their SDKs/libraries.

**Note**: The API calls are expensive network calls with high latencies.
one machine will not be able to make millions of concurrent network calls.

## Simple Notification Flow (One user)

1. PM creates Notn Template
2. PM triggers notn, the control service notifies the user instantly

**What if PM wants to trigger thousands of notifications at the same time?**
![](/diagrams/notificationservice.png)
Here the notification control service itself is triggering notification to your user using twilio etc. What if this system is down. Retries will wear down the service. 

1. Triggering one notification for every user is pain for PM.
2. Control service becomes the bottleneck (network call takes time) + provider outage?

**This is a classic place of making thing asynchronous.**


PM wants to send notification n1 to user u1. It would make a call to notification control service, it would get the template from the meta server.
It will create a final message that it will push into the queue. The message would contain everything that your notification emiting servers have to just send the notification. The workers would have the SDKs installed 
![](/diagrams/notificationservicequeue.png)

This architecture solves retries does not overwhelm Notification Control.

## Bulk Notification

Above architecture works well when we have moderate traffic and notifications are triggered one on one.
a typical usecase: Notify 'everyone'

Someone has to take care of the iteration of of users. Control service is not a good place to do this.
User experience: PM submits a job to notify everyone and we need to take care of everything else.

### Approach 1: Control Server Iterates (Workers)

- Iterating through millions of rows takes time.
- Plus, it eats up the control service, whose main responsibility should just be to accept "notification requests".

### Approach 2: Control Server Delegates iteration

![](/diagrams/notificationservicesqs.png)
When there was a bulk notification that we were sending, the bulk message goes with all the filters into the queue picked up by the iterator. This iterator would read the message and say on the user's table I want to iterate with this criteria. It would tak with users table and for each user it would get the notification template and create a notification message and then put into the queue. 

**Queue Separation**:
- **SQS 1**: Notification Emitter
- **SQS 2**: User iterator, and anything that requires some processing.

## Important Notifications

Some notifications are more important than others.
*e.g. Appointment Reminder >>> Marketing push*

In our current architecture, one massive marketing campaign will keep all the executors busy, and all other important notifications will starve in the queue.

**To solve this**: instead of having one Notification Emitter Queue, have multiple—**priority wise** (e.g., P1, P2, P3)—with each having its own set of dedicated workers.
![](/diagrams/notificationservicepriority.png)
When you send a notification you define the priority. Depending upon the priority the your iterator would put it in the corresponding queue.

## Duplicate Notifications

-  **Horizontal scalability**: Add more queues, add more workers
-  **Load isolation and avoiding starvation**: Priority queues

**x Duplicate notifications**

It is really irritating to receive marketing notifications, and it is the worst if we receive multiple of the same campaign.

To ensure we do not accidentally send multiple notifications from the same campaign, we keep track in a database (KV store is fine).

100M x (4B + 4B) = 800MB
5 campaigns ~ 4GB

Keep track of user & marketing campaign:
`<u1, mc1>`

*Delete or archive the old data to keep system performant.*

Iterator iterates through all the users and putting the message in the queue, piling up the queue again, consumed by the workers eventually realising the message was already sent. 
The iterator can use the notification tracker db to check if it has sent the message.
![](/diagrams/notificationservicetracker.png)

### Notification Tracker DB Requirements

- Lightweight and shard-able
- In-memory with periodic persistence
- *Bloom Filter support is a plus for being space efficient* (e.g. `mc1: BF<...>`)

**Redis** satisfies all 3 requirements.
