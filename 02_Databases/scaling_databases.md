# Scaling Databases

Databases are the most important component of any system out there. It makes or breaks any system...

Hence, it is critical to understand how to scale them...

## Vertical Scaling

- Add more CPU, RAM, Disk to the database
- Requires downtime during reboot
- Gives you ability to handle "scale", more load
- Vertical scaling has a physical hardware limitation
![](/diagrams/verticalscaling.png)

## Horizontal Scaling: Read Replicas

- When read:write = 90:10
- You move reads to other database so that "master" is free to do writes
- API servers should know which DB to connect to get things done

## Replication

Changes on one database (Master) need to be sent to Replica to maintain consistency.

### Two modes of replication
1. Synchronous replication

When write comes to the api server, it makes the write to the db. 
The writes from the master go to replica. The api server doesnt reply to the client unless the write has happended on master and replica.
![](/diagrams/synchronousreplication.png)
- Strong consistency
- slower writes
- zero replication lag

2. Asynchronous replication
Write coming to api server is written to the master and then the response is taken by api server, then sent back to the user.
- eventual consistency
- some replication lag
- faster writes
![](/diagrams/asynchronousreplication.png)

## Horizontal Scaling: Sharding

Because one node cannot handle the data/load, we split it into multiple exclusive subsets.

Writes on a particular row/document will go to one particular shard.

This way, we scale our overall database load.

- **Note:** Shards are independent (no replication between them).
- API server needs to know whom to connect to, to get things done.
- **Note:** Some databases have a proxy that takes care of routing.
- Each shard can have its own replica (if needed).




