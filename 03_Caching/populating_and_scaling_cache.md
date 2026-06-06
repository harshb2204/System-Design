# Populating the cache

Cache between the API server and the database.



There are two ways to populate the cache:

### 1. Lazy population (most popular)

- Read first go to the cache
- If data exist, return data
- Else
  - go to database / do heavy operation
  - persist in the cache
  - return data

> **Note:** Whenever we set something in cache, we set an expiry

#### Example: Caching Blogs (multiple joins)
- Fetching a blog from DB is expensive
- Hence, when someone accesses it we fetch from DB and cache it on Redis
- Subsequent requests are served from cache

### 2. Eager population

1. Writes go to both database and cache in the same request call.

#### Example: Live cricket score
- Thousands of people are watching cricket score.
- You will be serving it from cache.
- So, why not update cache & DB at once and save the cache miss.
![](/diagrams/cricinfo.png)

So when the score is updated if we wait for cache eviction/ the expiration of the value in cache there will be an inconsitency. Therefore when we are writing to db we are proactively write to cache as well. You are not waiting for key to expire. 

2. Proactively push data to cache because you anticipate the need.

#### Example: When a celebrity tweets / posts something
- When an account with 100,000 followers posts something, proactively push it to the cache.
- We will anyway need it and we save a cache miss.

# Scaling Cache

Cache is just like a database, hence scaling technique for a cache like Redis is similar to a regular database.

### 1. Vertical scaling
Make your cache bigger to handle more data / load.

### 2. Horizontal Scaling - Replica (scaling reads)
Same data replicated across multiple nodes so that reads could scale.

### 3. Horizontal Scaling - Sharding (scaling writes)
Data partitioned across multiple shards, so that writes could scale.
- Each shard can have a replica.
- Shards are mutually exclusive.