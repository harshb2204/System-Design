## Designing Twitter Trends
Twitter trends is a page where people see what is happening.
![](/diagrams/twitter_trends.png)

### Input to the system (almost realtime)

In order to build Twitter Trends, we need all the recent tweets.
Hence, we push tweets (as soon as they are published) in a persistent, high-throughput message stream - Kafka

Durability, Volume high (350,000 tweets per minute)
![](/diagrams/tweetsdb.png)


### News Clustering
We assume that we have a pre defined hierarchy. The hierarchy is of a particular set of topics that. At a very high level we would have this hierarchy, we would not have granular topics. Laliga and epl are very granular and might be happening at the moment, beyond this is what you would have defined.
![](/diagrams/hierarchy.png)
group similar news stories into clusters
There are various clustering algorithms that can be applied like k-means + tfidf
We have to prepare data for clustering
![](/diagrams/fetchdata.png)
