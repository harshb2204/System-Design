# Relational Databases

Databases are the most critical component of any system. They make or break a system.

Data is stored & represented in rows and columns.

## History of relational databases

Everything "revolutionary" (e.g., Computers, Internet, Blockchain) starts with Financial Applications.

- Computers first did "accounting" &rarr; ledgers &rarr; Rows & Columns
- Databases were developed to support accounting.

Hence, key properties were:
1. Data consistency
2. Data durability
3. Data integrity
4. Constraints
5. Everything in one place

Because of these reasons, relational databases provide "Transactions".

### ACID Properties
- **A**tomicity
- **C**onsistency
- **I**solation
- **D**urability

#### Atomicity
All statements within a transaction take effect or none.

**Example (Publish a post and increase total posts count):**
```sql
START TRANSACTION;

-- Publish a post
INSERT INTO posts VALUES (...);

-- Increase total posts count
UPDATE stats SET total_posts = total_posts + 1 WHERE user_id = 100;

COMMIT;
```

#### Consistency
Data will never go incorrect, no matter what.
- Maintained using tools like constraints, cascades, and triggers to ensure data never goes inconsistent (e.g., `total_posts` must equal the total entries in the `posts` table for a user).
- **Example:** Foreign key checks do not allow you to delete a parent if a child exists (this behavior can be tuned).

#### Isolation
When multiple transactions are executing parallelly, the isolation level determines how much changes of one transaction are visible to other.

#### Durability
When a transaction commits, the changes outlive outages (survive system crashes, power failures, etc.).

### Remember
You pick relational databases for relations and acid.

