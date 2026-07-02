# Consistent Hashing

One of the most amazing and popular algorithm out there.
and the only problem it solves is **Data Ownership**

## Hash Based Ownership

Say, we have a load balancer
and when a request comes in, it uses hash of access token
to decide which backend server to forward it to
![](/diagrams/hashbasedrouting.png)

**Note**: Hashing logic is NOT a 'service', but just a simple code running in the load balancer's code
-> more of a function


If one of the servers is taken down, then the routing function changes.
Now the requests will be evenly distributed between remaining two servers.

So, no hiccups!!
![](/diagrams/hashbasedrouting1.png)

## Load Balancers are stateless

Because the API servers are stateless,
which means every server is equally capable of handling requests
it does not matter if
request that was handled by server 1
now starts going to server 0.

This is precisely why we see 'Hash Based Routing'
as one of the most common ways of routing for
Stateless backends

eg. Load Balancer + API servers

## Hash Based 'routing' (ownership) for distributed storage

Instead of Stateless API request, say we are having a
stateful distributed storage

1. nodes store the data
2. proxy forwards the request to a node
3. end user / client talks to proxy
![](/diagrams/nodeshashbasedownership.png)

Say we want to store 6 keys K1, K2, K3, K4, K5 and K6:

K1 -> fn(k1) % 3 = 2
K2 -> fn(k2) % 3 = 0
K3 -> fn(k3) % 3 = 1
K4 -> fn(k4) % 3 = 2
K5 -> fn(k5) % 3 = 1
K6 -> fn(k6) % 3 = 0

*okay to do it when workload is stateless like API load balancer*

**Challenge**: if a storage node is removed or added,
the proxy cannot just forward request to any arbitrary node because it won't have the data.

### Repartitioning

When the number of nodes change, the proxy will change the
routing function and it would now become `fn(k) % 2`

Now, all the keys would need to be re-evaluated and moved to the correct node.
-> involves a lot of data transfer

**Before (3 nodes):**
K1 -> fn(k1) % 3 = 2       K4 -> fn(k4) % 3 = 2
K2 -> fn(k2) % 3 = 0       K5 -> fn(k5) % 3 = 1
K3 -> fn(k3) % 3 = 1       K6 -> fn(k6) % 3 = 0

**After (2 nodes):**
K1 -> fn(k1) % 2 = 0       K4 -> fn(k4) % 2 = 0
K2 -> fn(k2) % 2 = 1       K5 -> fn(k5) % 2 = 1
K3 -> fn(k3) % 2 = 1       K6 -> fn(k6) % 2 = 0

## Enter Consistent Hashing

Can we minimize the data movement?
This is where **consistent hashing** comes in.

**What consistent hashing is?**
An algorithm that helps in determining **data ownership** (who owns this data)

**What consistent hashing is not?**
*   it will not do data transfer for us
*   it is not a "service" in itself


### The Ring
![](/diagrams/consistenthashing.png)
**Hash function (SHA128)**
-> range of [0, 2^128)

Given hash functions are cyclic we can
visualize it as a ring of integers

Every node occupies one slot in
the ring, the slot is calculated
by passing node's IP to hash func^n

The ring can be modelled as a simple array and be part of proxy.

Here's how the ownership is determined
![](/diagrams/consistenthashing1.png)

K1 -> hash -> 0 -> node to right -> Node 0
K2 -> hash -> 10 -> node to right -> Node 3

Consistent Hashing ring will only tell you the node that "should" own it.

### Scaling up

when we add a new node to the "ring"
Say node 3 hashes to slot 1

The keys that hashed between slot 12
and slot 1 will now be "owned"
by node 3 instead of node 0

Other keys continue to remain at the respective nodes
-> Minimal Data Movement

Operationally, you just have to
1. snapshot node 0
2. create node 3
3. delete unwanted keys
![](/diagrams/consistenthashing2.png)

### Scaling Down

Say we scale down and remove node 0.
All the keys that were owned by Node 0
will now be owned by Node 2 (next in the ring).

Minimal Data Transfer

Operationally:
1. copy everything from Node 0 to Node 2
2. remove node 0 from the ring
3. delete node 0