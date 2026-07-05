# Designing Rate Limiter

Systems break down under tremendous load.
and we need to ensure that doesn't happen.

Design a rate limiter, that:
- limits the number of requests in a given period
- allows developers to configure threshold at a granular level
- does not add a massive additional overhead

## Rate Limiter Core Behavior

- It acts as the **First line of defence**.
- Any incoming request is first consulted against the rate limiter.
- If we are under limits, the request can **go through**.
- Otherwise, **reject with error (429)**.

## Where does it fit
![](/diagrams/ratelimiter1.png)

![](/diagrams/ratelimiter2.png)
Here the request comes to your service the first thing the service does is check with rate limiter. If the rate limiter says yes then it starts processing. 

## Dissecting Rate limiter

Rate limiter needs to track #requests in a given period.
Hence we need a database to hold the count, but which one?

1. for every incoming request, we would update the db
2. for every incoming request, we would read from db (aggregate)
3. we need ability to "clock" the time

Say, we rate limit per user:

1. we store #request in current period (`user_id -> count`)
2. once the period is over we reset the counter

Resetting the counter ≈ expiring the key
-  **KV store + Expiration -> Redis** -> *also gives us fast in-mem writes*

## Checking and updating

We have to pick one of two approaches:

1. **We expose a service (HTTP endpoint)** that talks to the rate limiter database.
2. **We let the services directly talk** to the rate limiter DB.
 - *saves one network hop*
 - *but exposes critical internal details*

We can make Rate limiter having its own Load Balancer and a bunch of API servers.
* But this adds substantial network hops.

![](/diagrams/ratelimiter4.png)

Nothing wrong in this, just be aware of the trade off.

To reduce the overhead of checking the Rate limiter, we let proxy / service directly manipulate the rate limiter database (saving 2 hops).

![](/diagrams/ratelimiter.png)
Rate limiter added as a library holding all the business logic.

## Scaling Rate Limiter

Given the services are directly talking to the Rate Limiter Database, **scaling rate limiter = scaling the database**.

1. **Should we scale vertically?** Yes.
2. **Should we add read replicas?** No! Traffic is not read heavy.
3. **Should we shard the database?** Yes.

↳ Any function in backend service:
  - will extract user_id from token
  - check rate limiter DB
  - if good, then proceed
  - otherwise, block

*Note: the access is key based:*
- `user_id -> count` (4 bytes + 4 bytes)
- `ip -> count` (16 bytes + 4 bytes)
- `token -> count` (16 bytes + 4 bytes)

Hence we can easily shard the database to handle more load.
![](/diagrams/ratelimiter5.png)
Storage is not an issue, compute is as you are doing count ++ because of which we are sharding the db. 

## Admin Console

It is better to have a small Admin Console (Frontend / Backend) that is used by the developers to reset counters and debug when things go wrong.

The service will also provide observability on infra and rate limiting DB like #keys, #request blocked, cpu/mem load, etc.