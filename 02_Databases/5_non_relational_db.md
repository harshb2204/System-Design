# Non-Relational Databases

It is a very broad generalization of databases
that are non-relational (MySQL, PostgreSQL, etc)

But this does not mean all non-relational databases are similar

## What makes Non-relational databases interesting?

Most non-relational databases **shard** out-of-the box!!
-> Horizontal Scalability

## Document DBs (MongoDB, Elasticsearch)

- Mostly JSON based
- Supports complex queries -> almost like relational (sql) databases
- Partial updates to documents possible
  - can do `total_posts += 1` without re-writing the entire document
- closest to relational database
- **Use cases**: in-app notification service, catalog service

```json
{
  "user_id": "_____",
  "total_posts": 270
}
```

## Key Value Stores (Redis, DynamoDB, Aerospike)

- Extremely simple databases (`GET (K)`)
- Limited functionalities (`GET`, `PUT`, `DEL`) (`PUT (K, V)`)
- Meant for key-based access pattern (`DEL (K)`)
  - -> *covers most of the use cases*
- Does not support complex queries (aggregations)
- Can be heavily sharded and partitioned
- **Use case**: profile data, order data, auth data, messages, etc.

*Note: You can use relational databases and document DBs as KV stores*

## Graph Databases (Neo4j, Neptune, DGraph)

* What if our graph data structure had a database?

- It stores data that are represented as nodes, edges, and relations
  - **eg:** `A --FOLLOWS--> B`
  - **eg:** `Arpit --BOUGHT--> iPAD`
- Great for running complex graph algorithms
- Powerful to model Social Networks, Recommendations & Fraud Detection
