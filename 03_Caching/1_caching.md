# Caching

Caches are anything that helps you avoid an **expensive** network I/O, disk I/O, or computation.

1. API call to get profile information
2. Reading a specific line from a file
3. doing multiple table joins

* **Performance Improvement**


Store frequently accessed data in a temporary storage location
![](/diagrams/caching.png)

Caches are faster and expensive.
* Hence we do not cache all the data (just a subset of it which is most likely to be accessed)

**Note:** Caches are not restricted to RAM based storage.
Any storage that is 'nearer' and helps you avoid something expensive is a cache for you!

In their simplest form,
* Caches are just glorified hash tables

## Some examples

1. **Google News:** Most recent news articles are more likely to be accessed, hence served from cache.
2. **Auth Tokens:** Authentication tokens are cached in "cache" to avoid load on database.
   - [tokens are checked on every request]
3. **Live Stream:** Last 10 min of Live Stream is cached on CDN as it will be accessed the most.


