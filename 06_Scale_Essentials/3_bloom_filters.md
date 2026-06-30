# Bloom Filters

Bloom Filters are **approximate data structures** that says with 100% certainty that an element does not belong to a set.
-> *also probabilistic data structure*

For example:
- Instagram wants to recommend reels.
- But it does not want to recommend something you saw already.

**Naive way**: keep track of everything that a user saw in a set
`u1 -> < p1, p4, p7, p1024, p1056 >`
(user) -> (post he/she saw)

People watch 100s of reels every single day.
Over time the set of all posts watched by a user will be huge!!

So what?
To check the existence of a key in a set, we have to load the entire set in memory and then check.
-> **Super expensive and time consuming**

So, can we do better?

**Key Insight**: Once something is "watched", you cannot take it back.
-> once a post/reel is watched by you, we do not take it out of the set!!

This means, storing actual data is not worth it.
This is the concept over which Bloom Filters are setup.

![](/diagrams/bloomfilter.png)

## Space Efficiency
- Bloom Filters take significantly less space to hold the information (because it does not store keys).
- and is very efficient in checking existence (just an array lookup).

## False Positivity Rate
- As the number of keys we put in a Bloom Filter increases...
- the false positivity rate increases.
  -> says key is present but in reality it is not.

Hence, when # of keys increases:
1. we have to re-create the bloom filter with larger size and populate the keys again.
2. estimate the max keys and provision a large one to start with.

## Practical Bloom Filter
- We do not have to re-implement Bloom Filters.
- There are libraries in every single language.
- **Redis** has it as one of its core features.
  -> nowadays, mostly people go for this.

## Practical Application

Use it whenever:
- you insert but not remove data
- you need a **NO** with 100% certainty
- having false positive is okay

Examples:
- Medium recommendation
- Feed generation
- Web crawler
- Tinder feed
