## Caching

- Reduces response times by saving any heavy computation or I/O
* Cache are not only RAM based (ex CDN)
Typical use: reduce disk I/O or network I/O or computation
Caches are just glorified Hash Tables


You should be okay with stale data.
When data gets changed in your source of truth, it should invalidate the data in the cache. 
![](/diagrams/cache.png)

What is the problem in the approach above?
- Cost of cache miss is very high.
- When the operation is expensive and there is a huge demand for that operation in case there is a cache miss.
- Your blog b1 is there in the cache and database. The request are coming , going to the cache and getting served. Imagine that entry blog b1 got deleted from the cache. Lot of req came at the very same time, all of them went to the cache and all of them saw entry does not exist. All the requests went to the db and fire the same query and then db fails. This is what we want to prevent. High throughput of read req coming in -> expensive fallback to db 

#### Solution

Debouncing(also called request hedging):multiple request to cache at the same time. One goes through db, while others wait.(CDNs use this)
```java

;

public class BlogService {

    private final ConcurrentHashMap<String, Semaphore> semMap = new ConcurrentHashMap<>();
    private final ConcurrentHashMap<String, String> resMap = new ConcurrentHashMap<>();

    private final Cache cache;
    private final Database db;

    public BlogService(Cache cache, Database db) {
        this.cache = cache;
        this.db = db;
    }

    public String getBlog(String key) throws InterruptedException {

        // 1. Check cache
        String value = cache.get(key);
        if (value != null) {
            return value;
        }

        // 2. Try becoming the leader for this key
        Semaphore mySemaphore = new Semaphore(0);

        Semaphore existing =
                semMap.putIfAbsent(key, mySemaphore);

        // ----------------------------------------------------
        // Someone else is already fetching from DB
        // ----------------------------------------------------
        if (existing != null) {

            existing.acquire();

            value = resMap.get(key);

            // Allow next waiting thread to proceed
            existing.release();

            return value;
        }

        // ----------------------------------------------------
        // We are the first thread (leader)
        // ----------------------------------------------------
        try {

            value = db.get(key);

            cache.put(key, value);

            resMap.put(key, value);

            return value;

        } finally {

            Semaphore s = semMap.remove(key);

            if (s != null) {
                int waiters = s.getQueueLength();

                // Wake all waiting threads
                s.release(waiters + 1);
            }
        }
    }

    // ---------------------------------------
    // Dummy interfaces
    // ---------------------------------------

    interface Cache {
        String get(String key);
        void put(String key, String value);
    }

    interface Database {
        String get(String key);
    }
}

```
Suppose 1000 requests come for blog:123:

All check cache → miss.
First thread executes:
```java
semMap.putIfAbsent("blog:123", mySemaphore);
```
and succeeds.
Remaining 999 threads see a semaphore already present:

```java
existing.acquire();
```
and block.
First thread fetches from DB:
```java
value = db.get("blog:123");
cache.put(key, value);
```
First thread releases waiting threads.
All waiting threads wake up and read:
```java
resMap.get(key);
```
instead of hitting DB.
Result:

DB calls = 1
Requests = 1000
All 1000 get the same response.