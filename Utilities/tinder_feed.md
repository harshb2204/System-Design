# Designing Tinder Feed

Design Feed for Tinder and feature to swipe left or right. For a great User Experience, a user should not be shown a profile that he/she already swiped.

## Feed Criteria and Preference

- **Proximity** -> near the better
- **Common Interest** -> shared interest

## Capturing Proximity

User's device (app) continuously emits `(lat, long)` to our backend. We store this information in a database.

**Database Requirements:**
1. That supports **geo-spatial queries** (*needed for generating the feed*)
2. Can be **horizontally scaled through sharding** (*needed to handle the huge incoming load as users emit location every...*)
![](/diagrams/locationemitter.png)

## Key Insight

Amount of data is not the concern but the query to store and get is.
Query load is the problem.

Hence we are sharding.

## Capturing Common Interest

**Profile Information**

1. Ask user to provide those details
2. Capture the details with Social login (Google/Facebook/Twitter)

**Authorizing for profile / interest / connection scrapping.**

![](/diagrams/enrichersprofile.png)
User logs in using authentication service. Auth DB will store username, password etc. When user signs in with facebook, google etc. you take that temporary token and push it to kafka. Then you invoke enrichers, which are a set of workers who make calls to facebook api, google api etc to fetch information with the token that the user has provided when they were authorizing. It then puts it in profile db. For each user you have facebook scraped groups, interests etc.

## Generating Feed

Now that we have the interest information and current location we can generate feed for users...

**Behaviour:** Generate feed when user is about to exhaust the current one

1. Frontend (app) will make an API call to trigger population of the feed
   - **Better** - Users swipes fast
2. Maintain a feed counter in backend and check every time user swipes
   - **Slow ux** - Users swipes fast

## Feed Database

Feed database can potentially explode (n×n) and will require decent storage.

For each user, the DB holds **Feed Items**

Each feed item = **Profile Information** that user has not seen before

**Approach 1**

- **Shard:** `< user_id, candidate_user_id, created_at >`
- **Partition by:** user_id
- **Ordered by:** ocreated_at
Becuase we are storing candidate_user_id here a challenge is that, when you send information on the frontend you also want profile data not just id.

But whenever we are returning, we have to fetch profile info in real time and send it to user (additional n/w call)

**Approach 2:**

- **Shard:** `< user_id, candidate_profile, created_at >`
- **Partition by:** user_id
- **Ordered by:** created_at

We require a significant storage + risk of serving stale data
But we avoid making call to profile service on runtime

### Important Note

In either case, you should never store feed of a user is "list"

- it will bloat up document size on mongodb
- expensive serialization and deserialization
- iteration is difficult

If we ar using dynamodb for the storage over here, user_id becomes the hashkey and created_at becomes the range key. If someones feed has 10 items there would be 10 docs in dynamodb. 

![](/diagrams/feedgenerator1.png)
Whenever frontend exhausts the feed, it makes an api call. Feed API pushes some message into the queue. Feed generator service accepts it from the queue, takes that message goes to the location db and finds the nearest profile, then it goes to the profile db and gets the info. It then generates the feed item and pushes it to the feed database

## Storing Swipes



People swipe left or right to indic ate interest.

We do not need a separate DB to hold this info and we leverage Feed Database and just add `"is_interested"` in each item.

**Note:** You do not have to create a new DB for everything

**When A Swipes B:**

- mark `is_interested` in feed item
- check `<B, A>` in feed DB
- if entry does not exist or `is_interested = false` : do nothing
- else: create a match (in match DB)

## Ensuring No-Repetition

- When A registers a "swipe" for B
- A should never see B again in the feed
- While generating the feed, we have to check for past swipes & add only if new

We need a definite no approximation yes is fine

**Classic case of Bloom Filter**

We use Redis & Periodic persist to store swipe information

Bloom filter is consulted before adding any item to feed
![](/diagrams/redisbloomfilter.png)
