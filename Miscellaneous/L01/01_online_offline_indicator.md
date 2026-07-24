## Online Offline Indicator

![](/diagrams/day0sketch00.png)
We have not decided yet on the db. Access pattern is key value, but there is no such thing as key value db. If you say dynamodb is kv db it supports range based operations on multi shards. You can still do a primary key lookup in mysql

#### Updating the database

Push based model : Users push their status periodically

Our API servers cannot "pull" from client because we cannot proactively talk to client * unless there is a persistent connection. 

Every user periodically sends heartbeat to the service.

POST / heartbeat → the authenticated user will be marked as "alive".
So when is a user offline? ->   When we did not receive "hearbeat" for long(subjective) enough 

Offline : When we do not receive heartbeat long enough
 Handling through business logic say 30 seconds

In database store "time you received the last heartbeat"

![](/diagrams/hearbeat00.png)
Now our db contains last hearbeat so our query needs to change. 

#### Scale Estimation

**User Scale:**
- 100 users → 100 entries
- 1,000 users → 1,000 entries
- 1M users → 1M entries
- 1B users → 1B entries

**Database Schema:**
Each entry has 2 columns:
- `user_id` (int, 4B)
- `last_hb` (int, 4B)

**Storage Calculation:**
- Size of each entry = 8B (1 billion bytes = 1 GB)
- Total storage for 1B entries = 8 GB

Always critically challenge your system

* does not mean you alter your system

But how to find ?? → when you have 'n' paths, evaluate all of them

[Framework of the Opposites]

What if we dont store all the entries?? Dense vs Sparse
Requirement: We only care if user is online or offline

absence == offline ?
Idea: if user not present in the DB, we return offline.
So, let's expire the entries after 30 seconds -> delete
if we delete entries → we save a bunch of space by not storing data of inactive users
Total entries = active users
if 1B total users and 100k are active then total entries = 100K
total size = 800KB

#### How to auto delete?

Approach 1: Write a CRON job that deletes expired entries
Approach 2: Can we not offload this to our datastore ?

DB with KV + expiration → Redis (OSS)
                        → DynamoDB

Upon receiving an heartbeat
- update entry in Redis / DynamoDB with ttl = 30 sec

Every heartbeat moves the expiration time forward !

![](/diagrams/redisvsdynamodb.png)

Note: In real world, WebSockets are used in such systems, but this is Day 1, hence we keep things simple.

Assume: you are using socket IO library


* this is not a websocket specification [AFC 6455]

Socket IO implement its own heartbeat mechanism

1. own protocol over websocket
2. provides automatic connection monitoring and reconnection
Scaling websockets is hard.
* if you are using any other lib, you can always implement your own HB mechanism

#### How is our DB doing ? [Poll based approach]

Heartbeat is sent every 10 seconds
So, one user in 1 minute sends 6 heartbeats
If there are 1m active users, our system will get 6m req/m
Each heartbeat request results in 1 DB call


Our db needs to handle 6M updates per minute. micro reads/updates -> hence n/w bottlneck.
Every request we are getting is actually making a call to db and updating the value. 
Naive implementation:- We get the req, the api server establishes the connection with the db. fires the query, gets the response and breaks the connection.  If for every request it needs to setup a connection (3 way handshake, one req one response, 2 way teardown) it is very hectic.

#### Connection Pool
We can instead reuse connections instead of creating one evertime.
The idea is your server creates a pool of connections with the db and when the req comes to the server, server picks one of them whichever is free and fires the query.



##### How is connection pool implemented ?

Blocking queue → you wait unless there is something to remove

Queue data structure
implemented as array → Each element is TCP connection 0

![](/diagrams/connectionpooling.png)

To decide max number of connections db can handle the only way is to do load testing. See when db starts to throttle. (1.2 * (limit/no. of api servers))
