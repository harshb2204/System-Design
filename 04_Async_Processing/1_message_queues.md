## Asynchronous Processing


* **user sent the request and you immediately handle it is synchronous**

* **loading insta feed is synchronous**
* **login on website is synchronous**
* **payments are synchronous**
  * *(Most interactions on web are synchronous)*
* **some things that should not be synchronous**

1. Spinning up virtual machine
![](/diagrams/vmspinup.png)
* **spinning up vm takes minutes, and user will not wait on the same page, waiting for the response**
* **Instead, he/she would love to move around and keep checking status once in a while. This is asynchronous**


Client makes a request, it comes to the api server. Api server registers something in the db or pushes a task to the message queue or a message broker
and immediately responds back to the user. The workers nodes pull the task out and update the db when it is done. 
![](/diagrams/asyncprocessing.png)

## Message Queues

* **Brokers help two services/applications communicate through messages.**
* **We use message brokers when we want to do something asynchronously.**
  1. **Long-running task**
  2. **Trigger dependent tasks across machines**
* **Example: Video processing**

![](/diagrams/asyncvidresolution.png)

Once the video is uploaded we need to convert it to 720p, 360p, , 480p
When the video upload service uploads the video to s3 it pushes a message to the message queue or a message broker. (sqs, rabbitmq). When it sends the message to the message broker it is consumed by the video processing service.
In the message u send details like video id, who uploaded the video etc, that the video processing service would use to process and reupload them to s3 in different formats.
When the message is received by the video processing worker, from that message it would know the video id, from the video id it would know the path the video lies in s3. It downloads the video from s3 on its machine, it would convert it to different resolutions and upload them back to s3. 

## Features of message brokers

1. **Brokers help us connect different sub-systems**


2. **Brokers act as a buffer for the messages**
   * **i.e. consumers can consume at their own pace**
   * **eg: Notification system** *(no synchronous load on connected systems)*

3. **Brokers can retain messages for 'n' days** *(depends on broker you use)*

4. **Brokers can re-queue the message if not deleted**
  Someone placed an order , it pushed the message to the broker. The message was consumed by the email sender but before it could make an api call it crashed. In this case every broker has a different way of handling it. Broker gives u apis to call to delete a particular message. When you consume a message you send what u want to and then u delete the message. When someone says i read the message from the broker does not mean that the message is deleted from the broker. You have to explicitly consume and then delete the message. 
  If now before deleting your job killed. Then after some time period the message would reappear at the head of the queue to be consumed by some other consumer. Its never that because u did not delete the message the message is permanantly gone, you did not delete the message so that it would show up after some time that is visibility timeout (sqs property). It is possible tha you consumed the message made an email api, but before you could delete the message the process crashed. So some users might get email twice. Always be vary of the fact that some messages may be delivered more than once. 
   * **eg: consumer read the message but before it could delete it, it crashed**

## Typical flow while using message queue

#### Example Auto-Subtitle (auto-captioning)
![](/diagrams/autocaptioning.png)

1. **user uploads video to s3 through video service**
2. **video service puts a message (after upload completes) to broker**
3. **and returns response to the user**
   * *user sees upload complete*
4. **message is asynchronously read by 'captioner' service**
5. **captioner downloads the video**
6. **captioner generates captions and updates in the DB**
7. **user now sees caption button enabled**
