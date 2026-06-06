# Caching at different levels

Most common cache we saw was Redis
but that is not the only type of cache out there
or not the only place that can be used as a cache

Literally every piece/component in your infrastructure
can cache something for you. But should you?
↓
depends on the guarantees -> stale data and invalidation
also, too much caching is bad!

Caching financial summaries might not be a good idea, as u can see inconsistencies in data. Recency here is very important. You can still cache user details.

You can store something in a cache with an expiry. Once the cache expires it gets autodeleted and for the next reqeust it gets filled up again.
If we want to proactively update the cache its called invalidation (either deleting the key or updating with new content)
let's take a look at different places where we can cache

## Client Side Caching

Storing frequently accessed data on client side

eg: Browser, Mobile Devices, etc.

- cache near constant data
  (eg: images, js files, user information, etc)
- it should be okay serving cached info (stale)
- invalidation by time (expiry)

Massive performance boost, as we need not make
any request to backend

## Content Delivery Networks (CDN)

*used for caching*
- Live Streaming
- Serving Images, Videos, Audio, Bundles, etc

CDNs are a set of servers distributed across the world.

Request from a user, goes to the nearest
CDN server and hence user gets very quick response

US folks getting images from US servers is faster
than fetching it from India.

* CDN does lazy cache population
![](/diagrams/cdn.png)

user's request comes to CDN -> closest server
CDN server checks if it has the data
if yes, return the data
else,
  CDN makes the same request to origin
  gets the response
  caches the response
  return the data

like any other cache, when you put data on CDN
you set an expiry to it (post which CDN deletes data)

## Remote cache (Redis)

*Expensive & store data in main memory*

Remote cache is centralized cache that
we most commonly use (Redis). Multiple
API servers use it to store frequently
accessed data.

* Every key stored should have an expiration (memory is expensive)
* Size of cache is relatively very small as compared to database

## Database caching

Instead of computing total posts by users everytime
we store `total_posts` as column and update it once
a while. (saves an expensive DB computation)

```sql
select count(*) from posts
where user_id = 123
```
*(expensive query)*

**users**
| id | name | .... | total_posts |
|---|---|---|---|
| 123 | Arpit | | 77 |

Everytime a post is published, we
also update the
`users` table and do
`total_posts = total_posts + 1`


The more layers of cache you add the more stale data you are adding.