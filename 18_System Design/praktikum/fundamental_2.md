## 1. Redis
- Make SET
`SADD users user1 user2`
- Retrieve data
`SMEMBERS users`

## 2. Neo4j
- Create relationship
```
MATCH (u1:User {id: "user1"}), (u2:User {id: "user2"})
CREATE (u1)-[:FOLLOWS]->(u2)
```

- Retrieve data
```
MATCH (u:User)
RETURN u
```

## 3. Cassandra
- Create table
`CREATE TABLE users ( id UUID PRIMARY KEY, fullname text, birthday timestamp, email text );`

- Retrieve data
`SELECT * FROM users;`