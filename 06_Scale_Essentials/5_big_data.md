# Big Data Processing

When one machine is not enough to process the data
we **divide and conquer** -> This essentially is Big Data Processing!

Companies use this to process massive amount of data and extract insights
out of it, train ML models, move data across databases, and much more.

- when too much data needs to be processed quickly, we use Big Data tools.

*all these fancy processing on commodity hardware* (not specialised hardwares which you can find anywhere)

## Counting word frequency

Given a 1 TB text dataset find frequency of each word

**Approach 1: Simple**

1. load data on one machine (disk)
2. read it character by character
3. when space is encountered,       `word_freq[word] += 1`
   update an in-mem hash table
   do count++

This is a simple approach that runs in O(n) time approx.

But because only one machine is doing it,
it will take a long time. So, can we parallelize it?
Yes... add threads....

**Approach 2: Threads**

You can easily parallelize the code
each thread can handle a chunk
of the file/dataset and can do `word_freq[word] += 1`

But, what if the dataset is not 1 TB but 100 TB?
-> big data on smaller hardware

-> something that does not fit on a single machine, or
-> even if it did, it is slow to compute

-> threads are bounded by cpu cores
-> limited computational capabilities of underlying hardware

**Approach 3: Distributed Computing**

Instead of one machine, can we distribute the workload
across multiple (often smaller) machines
-> and leverage parallelism

More computers, more CPU, more processing

**idea**:
1. split the file into 'partitions'
2. distribute the partitions across all servers
3. let each server compute word freq independently
4. send word freq to one server (coordinator)
5. merge the word freq
6. return the result

![](/diagrams/distributedjob.png)

- user submits the job to 'coordinator'
- coordinator distributes the job across multiple machines
- machines compute and send result to coordinator
- coordinator merges and returns

**Challenges**:
1. what about failures          3. what about recovery
2. what about completion        4. what about scaling & distribution

## Big Data Tools

Although, we can always do it on our own,
but if there tools to manage this for us... adopt it.

Big Data Tools manages these complexities for us
-> we just write the business logic

- distribute across machines
- knowing which machines are doing what
- retrying in case of failures
- reprocessing in case of crashes or corruption
- cleaning up the resources once job is complete

### Spark and Flink

- large scale data processing on commodity hardware
- it has connectors to a lot of databases and infra components
- eg: combine user, order, payments and logistics DBs and put the result in AWS Redshift


Lets say we have a lot of events being ingested into kafka, some user facing events. We want to enrich those events for example if a blog is published we want to enrich it with 'who published the blog' or 'is this a paid user' and then we want to send the result to the elasticsearch so that the stakeholders can visualize it. Spark will take care of making the call, merging it (enriching it). One machine will not be able to make 2 db calls and then make a 3rd db call. It needs to be done in a distributed way, reading from kafka make a call to the db, enriching the info etc. Big Data computing is not only about doing average, percentage etc. Its about computing anything that you want to in a ditributed fashion. 
![](/diagrams/sparkafka.png)