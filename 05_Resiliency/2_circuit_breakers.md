# Circuit Breakers

Circuit Breakers prevent **CASCADING FAILURES**!!

Say, you are building a social network that serves feed for a user:

1. user's request comes to Feed Service
2. Feed service pulls some info from Recommendation, some from Trending
3. Recommendation and Trending both relies on Profile service
   * to get profile detail of user who made the post
4. Recommendation and Trending depends on Post service
   * to get details of the post

There are lots of other services that depend on profile service.

If profile service DB is overwhelmed !! it slows down
and transitively all services dependent on it are affected.

"Timeouts" [higher response times]
![](/diagrams/circuitbreakers1.png)

1. complete outage -> Poor user experience
2. unresponsiveness

Idea: what if we make call to a service, only if it is healthy?

This is Circuit Breaker

we break the circuit down
when we see the failure cascade!

Circuit Breakers prevent entire product from collapsing.
by preventing cascading failures.

## How is it implemented?

- A common database holds the settings for each breaker
- services, before making calls to others, checks the config
   * cache the config to avoid checking the DB
![](/diagrams/circuitbreaker2.png)

In case of the outage,
the circuit is tripped and DB updated

Services will periodically check and stop sending req. to affected services
