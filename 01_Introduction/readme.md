## What is System Design
- We have a set of requirements and we want to build a product out of it.

From a **Set of requirements**, we:
- **Decide architecture**
- **Decide components**
- **Decide modules**

These define **how they interact with each other** to **solve the problem** (which is almost **"Product Development"**).

## Why is it so popular?
- Every single "tech" product is a "system" that has been "designed".
- Companies are building products and need people to design them.

## Why understanding system design is important?
- This is what people do at work.
- Everything is practical.
- You can see it from Day 1 of joining the company.
- Once you grow in your career, you will spend 80% of time doing this.

## How to approach system?
- System Design is extremely practical and there is a structured way to tackle the situations.
- Take baby steps, no matter what! Start with high-level design and incrementally deep-dive into individual components.
- **Understand the problem statement**: Without having a thorough understanding of the problem at hand, we would easily digress.
  - Always clarify requirements (both functional and non-functional) and constraints (scale, latency) before designing.
- **Break it down into components (essential)**:
  - *Note (to start with)*: Do not create components for the sake of it. Only create components that you know are a must.
  - *Example (Design Facebook)*: Identify core components/features when stated/requested, such as Auth, Notification, Feed, and Gamification.
  - Start with a lean MVP architecture and introduce additional components only as scaling or functional requirements demand them.
- **For each sub-component, look into (repeat for each subcomponent one by one)**:
  1. **Database and Caching**: Determine data storage models, read/write patterns, and caching layers to optimize performance.
  2. **Scaling & Fault Tolerance**: Plan for horizontal scaling, redundancy, replication, and failover strategies.
  3. **Async processing (Delegation)**: Delegate non-blocking, heavy operations to message queues or background processors.
  4. **Communication**: Define interaction models between components (e.g., REST, gRPC, WebSockets, pub/sub).

![](/diagrams/feedgenerator.png)
Dependency on other services

## How do you know you have built a good system?
- Every system is "infinitely" buildable, and hence knowing when to stop the evolution is important.
- Here are some pointers that will help you:
  1. You broke your system into components
  2. Every component has a clear set of responsibilities (exclusive)
     - **Feed webserver**: Serves feed over HTTP.
     - **Feed Generator**: Pulls data from multiple services (posts, friends, recommendation) and puts them in DB (candidate feed items).
     - **Feed Aggregator**: Combines candidate items fetched by generator, filters out redundant ones, ranks them, and creates a final consumable feed.
  3. For each component, you've slight technical details figured out
     1. Database and Caching
     2. Scaling & Fault Tolerance
     3. Async processing (Delegation)
     4. Communication
  4. Each component (in isolation) is:
     - **Scalable**: Horizontally scalable.
     - **Fault tolerant**: Plan for recovery in case of a failure to a stable state (mostly data*).
     - **Available**: Component functions even when some component "fails".

This is precisely how we would tackle every single system in a structured and detailed way.
