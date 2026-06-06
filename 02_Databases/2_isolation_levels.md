# Isolation Levels

Relational databases provide ACID guarantees, and the **I** in ACID is **Isolation**. Isolation levels help us tune them.

Isolation levels dictate how much one transaction knows about the other.

We look at each one of them and understand with examples.

## Repeatable Reads
Consistent reads within the same transaction. Even if another transaction committed, the 1st transaction would not see the changes (if the value was already read).

## Read Committed
Reads within the same transaction always read the fresh value.
- **Con:** Multiple reads within the same transaction are inconsistent.

## Read Uncommitted
Reads even uncommitted values from other transactions (txns).
- Known as a **"dirty read"**.

## Serializable
Every read is a locking read (depends on engine), and while one transaction reads, others will have to wait.


