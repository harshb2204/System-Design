## Multi User Blogging Platform (Medium.com)
One user multiple blogs
Multiple users

#### Database

users
- id
- name
- bio

blogs
- id
- author_id
- title
- body
- is_deleted
- published_at

Importance of is_deleted [soft delete]
when user invokes delete blog, instead of DELETE we UPDATE
Key reasons: Recoverability, Archival, Audit
+ Easy on the database engine [No tree re-balancing]
![](/diagrams/rebalancing.png)
So soft delete the rows, and then later you can batch delete. 

