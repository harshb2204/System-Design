## Scaling

Ability to handle large number of concurrent requests

Two scaling strategies

- Vertical scaling → Make infra bulky, add more CPU,Bulk Disk

- Horizontal scaling -> Linear amplification: Given that you know your `UNIT TECH economics` (how many requests one server can handle), how many servers would we need. Only way to know how many requests a server can handle is through a load test.
Why isn't there a formula for this -> Even if there was one it would be making some assumption. (how do you know if the req id cpu intensive or memory intensive, multiple queries or single queries...). Given that are so many variables you cannot come up with a formula here.
Fault Tolerance


* Horizontal scaling ≈ ∞ scale, but there is a catch tot his. 

Your stateful components like DB and cache should be able to handle those concurrent requests.
Hence whenever you scale, always do it bottom up.
Scale DB first -> then API servers.


### Scaling DB
![](/diagrams/scalingdb.png)
Most people say -> "writes will go to master, reads will go to replica."
This does not mean reads will not go to master. Your writes and critical reads -> reads where consistency is important (account balance) will go to master and all other reads which are okay with staleness will go to replica. 
a -> b != ~a->~b

Either API servers knows the DB topology or you add a proxy that is aware of topology and takes care of routing.
![](/diagrams/proxysql.png)



#### ProxySQL Primer
A hostgroup is simply a logical group of database servers.
create hostgroups representing different clusters
For example:
Hostgroup 10 → Primary (write database)
Hostgroup 20 → Replicas (read databases)
```sql
INSERT INTO mysql_servers (hostgroup_id, hostname, port)
VALUES (10, 'db-primary.example.com', 3306),
       (20, 'db-replica.example.com', 3306);

LOAD MYSQL SERVERS;
SAVE MYSQL SERVERS;
```
LOAD → Apply the configuration immediately (runtime).
SAVE → Persist it so it survives a ProxySQL restart.

add multiple replicas in same hostgroup


```sql
INSERT INTO mysql_servers (hostgroup_id, hostname, port)
VALUES (20, 'replica1.example.com', 3306),
       (20, 'replica2.example.com', 3306),
       (20, 'replica3.example.com', 3306);

LOAD MYSQL SERVERS;
SAVE MYSQL SERVERS;
```
least connection based routing


Every SQL query starting with SELECT should go to Hostgroup 20.
```sql
INSERT INTO mysql_query_rules (rule_id, active, match_pattern, destination_hostgroup)
VALUES (1, 1, '^SELECT', 20, 1);

LOAD MYSQL QUERY RULES;
SAVE MYSQL QUERY RULES;
```

Instead of routing based on SQL commands, you can route based on the database (schema).
```sql
INSERT INTO mysql_query_rules (rule_id, active, schemaname, destination_hostgroup)
VALUES (1, 1, 'analytics_db', 20, 1),
       (2, 1, 'transactions_db', 10, 1);

LOAD MYSQL QUERY RULES;
SAVE MYSQL QUERY RULES;
```
If the application is connected to
```sql

USE analytics_db;
```
then every query goes to Hostgroup 20.

#### Scaling ProxySQL
![](/diagrams/proxysqlscaling.png)
But now our API servers need to know the proxy server topology.
To handle this add a load balancer. API server will connect to a proxy server instance 



