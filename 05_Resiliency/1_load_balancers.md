# Load Balancers

![](/diagrams/loadbalancer.png)

* **One of the most important components in a distributed system that makes it easy to scale the load horizontally.**
* **Load Balancer is the only point of contact.**

### Every load balancer has either:
1. **Static IP**
2. **Static DNS name**

*(allowing clients (users / servers) to talk to it)*

* **Load balancer hides the # of servers that are "behind" it.**
  * This allows us to add as many servers as possible without the client knowing about it (**Horizontal Scalability**).

---

## Request Response Flow

1. **Client already has IP/domain of the load balancer**
   * *eg: auth.example.com*
2. **Client makes API call and it comes to the load balancer**
   * *eg: GET auth.example.com/login*
3. **Load balancer picks one server and makes the same request**
4. **Load balancer gets the response from the server**
5. **Load balancer responds back to the client**

---

* **Job of the load balancer is to "balance" the load.**
  * It depends on how it picks the server to forward the request (which is **configurable**).

---

## Load Balancing Algorithms

### 1. Round Robin
* **Uniform infrastructure**
* Distribute the load iteratively.
![](/diagrams/roundrobin.png)

### 2. Weighted Round Robin
* **Non-uniform infrastructure**
* Distribute the load iteratively but as per weights.
![](/diagrams/roundrobinweighted.png)

### 3. Least Connections
* **Pick the server having the least connections from the load balancer**
* **Used when response times have a big variance**
  * *eg: Analytics*

### 4. Hash Based Routing
* **Hash of some attribute (IP, user ID, URL) determines which server to pick.**
  * *Provides a "random enough" distribution.*

---

## Key Advantages of Load Balancers

### 1. Scalability
* **With more servers behind the load balancer, we can now handle more requests.**
  * *eg: Two servers (100 RPM each) → We can handle 200 RPM.*
  * *eg: Three servers (100 RPM each) → We can handle 300 RPM.*

### 2. Availability
* **Even if one of the servers crashes, it does not take down our entire system.**
* **Load balancer will forward requests to other healthy servers, improving the availability.**
  * *eg: If server `s2` is down, the load balancer will forward any new requests to `s1` and `s3`.*
