# Leader Election for Auto Recovery

Say, we have a bunch of servers, serving HTTP requests.

When one of them goes down, we have to spin up a new one.
Doing this is responsibility of another module, say **orchestrator**.

## Orchestrator
- Keeps an eye on the servers.
- When one goes down, Orchestrator spins one and add
  * No human intervention
  * minimal time - outage!

But what if orchestrator is down? Who monitors it?
- we need orchestrator for orchestrator
- who monitors that?
- another orchestrator??

This goes out of hand quickly!

Hence, we need an automated way to recover the system.
- when one orchestrator is down, somehow other orchestrator comes back up and takes responsibilities

## Leader Follower Setup

We run orchestrator in leader-follower mode. multiple nodes running the orchestrator code. One leader while others are workers / follower.
![](/diagrams/leaderfollower.png)


- leader keeps an eye on workers. if worker dies, leader spins up new one
- Workers ping the servers and checks if they are healthy.
- if orch worker finds a backend server is unhealthy, it spins new one
- if orch leader finds that orch worker is unhealthy, it spins new one
- if orch worker finds orch leader is dead, they trigger a leader election and system is auto-recovered.

How workers will choose a leader depends on the Leader Election algo.
