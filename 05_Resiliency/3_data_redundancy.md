# Data Redundancy & Recovery

API servers are "stateless" but databases are "stateful".

* **API servers going down is fine** because a new one will be spun up almost instantly.
  * API server gets a request and it does not matter which one handles it.

* **Databases going down is catastrophic**; almost always an outage!!
  * Worst: Disk crash leading to loss of data!

A good system always takes care of such catastrophic situations.
* The only way to protect ourselves against loss of data is to create multiple copies of it -> **Data Redundancy**

Redundancy can be implemented at row/document level, table level or DB level.
Redundant data can be stored on different table, different DB or different region.

## Backup and Restore

- Daily backup of data (incremental)
- Weekly complete backup
- storing one copy across region -> Disaster recovery
- when something goes wrong, just restore the last backup
- almost always the easiest thing to do

## Continuous Redundancy

Setup replica of the database
and writes go to both DB (sync / async)

1. API server writing to both Databases
2. API writes to one and is copied to other asynchronously

* if the main database goes down, replica can take its place (almost instantly)
  * note: this replica may just be a stand-by and not serve any production traffic
![](/diagrams/dataredundancy.png)
