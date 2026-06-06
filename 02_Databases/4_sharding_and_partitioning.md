# Sharding and Partitioning

* **Sharding:** Method of distributing data across **multiple machines**.
* **Partitioning:** Splitting a subset of data **within the same instance**.

## How a db is scaled?
![](/diagrams/dbserver.png)



You put your database in production, serving real traffic (e.g., 100 WPS).
![](/diagrams/db1.png)

You are getting more users, that your DB is unable to manage:
* You **scale up** your DB... give it more **CPU, RAM and Disk**.
![](/diagrams/db2.png)

But, after a certain stage you know you would not be able to scale "up" your DB because **vertical scaling has a limit**.

So, you will have to resort to **Horizontal Scaling**.

Say, one DB server was handling **1000 WPS** and we cannot scale up beyond that, but we are getting **1500 WPS**. We scale horizontally and **split** the data.
![](/diagrams/horizontalscaling.png)

By adding one more database server, we reduced the load to **750 WPS** on each node and thus handled **higher throughput**.

Each database server is thus a **shard**, and we say that the data is **partitioned**.

Overall, a database is **sharded** while the data is **partitioned** (split across).
> **Note (Oversimplification):** Most people use the terms interchangeably.

---

### Partition Allocation

You partitioned the **100 GB** of total data into **5 mutually exclusive partitions** (e.g., `30 GB`, `10 GB`, `30 GB`, `20 GB`, `10 GB`).

Each of these partitions can either live on **one database server**, or a couple of them can **share one server**.

*And this depends on the **# of shards** you have.*

---

## How to partition the data?

There are two categories of partitioning:
1. **Horizontal Partitioning**
2. **Vertical Partitioning**

When we "split" the 100 GB data, we could have used either of the ways, but deciding which one to pick depends on **load, usecase, and access pattern**.

---

## Pros & Cons of Sharding

| Advantages of Sharding | Disadvantages of Sharding |
| :--- | :--- |
| - Handle large Reads and writes | - Operationally complex |
| - Increase overall storage capacity | - Cross-shard queries expensive |
| - Higher availability | |
