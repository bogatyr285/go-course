### Basic Data Structures

#### 1. Strings
**Use Cases:** Used for caching frequently accessed data, maintaining counters, managing distributed locks, and keeping session data. Strings are the most straightforward data type in Redis.

**Examples:**
```sh
# Cache the user data under a key "cache:user:1000"
SET cache:user:1000 "John Doe"
# Retrieve the cached user data
GET cache:user:1000
# Increment the counter for homepage views
INCR counter:views:homepage
# Attempt to acquire a distributed lock for resource 123
# “setting the key if not exists”
SETNX lock:resource:123 "1" # Returns 1 if the lock is acquired, 0 if it is already locked
```

#### 2. Lists
**Use Cases:** Ideal for implementing message queues, task queues, or stacks. Lists maintain the order of the elements, enabling LIFO/FIFO operations.

**Examples:**
```sh
# Push tasks "task1" and "task2" to the left end of the list
LPUSH queue:tasks "task1" "task2"
# Pop the most recent task from the left end of the list
LPOP queue:tasks # Pops "task2"
# Get the length of the list
LLEN queue:tasks
# Move tasks between lists (common in task/queue management)
LMOVE queue:tasks pending-tasks LEFT LEFT
```

#### 3. Sets
**Use Cases:** Useful when you need to store and manage unique items, or perform set operations like intersections, unions, and differences.

**Examples:**
```sh
# Add users 1001, 1002, and 1003 to the set of online users
SADD users:online 1001 1002 1003
# Remove user 1002 from the set of online users
SREM users:online 1002
# Get the number of online users
SCARD users:online
# Check if user 1001 is online
SISMEMBER users:online 1001
# Get the intersection of online users and subscribed users
SINTER users:online users:subscribed
```

#### 4. Hashes
**Use Cases:** Optimal for representing objects and storing related information, such as user profiles or configuration settings.

**Examples:**
```sh
# Set user's name and age in a hash
HSET user:1000 name "John" age 30
# Get the user's name from the hash
HGET user:1000 name
# Get multiple fields (name and age) from the hash
HMGET user:1000 name age
# Increment the user's age by 1
HINCRBY user:1000 age 1
```

#### 5. Sorted Sets (ZSet)
**Use Cases:** Suitable for scenarios where ranking is required, such as leaderboards, priority queues, or time-series data.

**Examples:**
```sh
# Add users to a leaderboard with scores
ZADD leaderboard 1000 "user1" 2000 "user2"
# Get the entire leaderboard with scores in ascending order
ZRANGE leaderboard 0 -1 WITHSCORES
# Get the leaderboard in descending order
ZREVRANGE leaderboard 0 -1 WITHSCORES
# Get the users with scores between 1000 and 2000
ZRANGEBYSCORE leaderboard 1000 2000
# Remove users with scores between 0 and 500
ZREMRANGEBYSCORE leaderboard 0 500
# Get the rank of "user1" in the leaderboard
ZRANK leaderboard "user1"
```


#### 6. Streams
**Use Cases:** Well-suited for event sourcing, sensor data monitoring, logging, and other append-only log scenarios.

**Examples:**
```sh
# Add a new entry to the stream with automatic ID generation (*)
XADD stream:sensors * sensor-id 1234 temperature 30
# Read the first two entries from the stream
XREAD COUNT 2 STREAMS stream:sensors 0
```

#### 7. Geospatial
**Use Cases:** Designed for storing and querying geographical coordinates, making it possible to build location-based features like geofencing and proximity searches.

**Examples:**
```sh
# Add geographical coordinates for Palermo and Catania
GEOADD locations 13.361389 38.115556 "Palermo" 15.087269 37.502669 "Catania"
# Find locations within 200 km radius of the given coordinates
GEORADIUS locations 15 37 200 km
```

#### 8. Bitmaps
**Use Cases:** Excellent for tracking binary states efficiently, like daily user login status, feature activation flags, or A/B testing.

**Examples:**
```sh
# Set the bit at position 1 to 1 for user's login status
SETBIT user:login:1000 1 1
# Get the bit value at position 1 for user's login status
GETBIT user:login:1000 1
# Count the number of bits set to 1
BITCOUNT user:login:1000
```

### Added Later - Probabilistic Data Structures for Big Data

#### 9. HyperLogLog
**Use Cases:** Utilized for estimating the cardinality of very large datasets, such as unique visitors or unique search queries, with minimal memory overhead.

**Examples:**
```sh
# Add visitors to the unique visitors HyperLogLog
PFADD unique:visitors user1 user2 user3
# Get the estimated number of unique visitors
PFCOUNT unique:visitors
```

#### 10. Bloom Filter
**Use Cases:** Handy for determining if an element is possibly in a set, used in search engines, databases, and cache systems to reduce the need for expensive lookups.
**Examples:**
```sh
# Add an element to a Bloom filter
BF.ADD bloom:filter user1
# Check if the element possibly exists in the Bloom filter
BF.EXISTS bloom:filter user1
```

#### 11. Cuckoo Filter
**Use Cases:** Efficiently represents sets with low memory usage and allows insertions and lookups. Used in network devices, storage systems, and database engines.

**Examples:**
```sh
# Add an element to a Cuckoo filter
CF.ADD cuckoo:filter user1
# Check if the element possibly exists in the Cuckoo filter
CF.EXISTS cuckoo:filter user1
```

#### 13. Top-k
**Use Cases:** Aimed at identifying the most frequent items in a data stream, beneficial for trending topics, frequent queries, or popular items analysis.

**Examples:**
https://redis.io/docs/latest/develop/data-types/probabilistic/top-k/


# Full text-search & JSON manipulation

### 1. Full-Text Search

**Example: Indexing and searching text documents:**

1. **Create an index with schema:**

```sh
FT.CREATE myIndex ON HASH PREFIX 1 doc: SCHEMA title TEXT WEIGHT 5.0 body TEXT url TAG
```

2. **Add documents:**

```sh
HSET doc:1 title "Redis Stack" body "Redis Stack is a powerful combination of multiple Redis modules." url "https://redis.io/docs/stack/"
HSET doc:2 title "Redis Search" body "Full-text search capabilities of Redis Stack." url "https://redis.io/docs/stack/search/"
```

3. **Search documents:**

```sh
FT.SEARCH myIndex "Redis" RETURN 2 title url
```

### 2. JSON

**Example: Storing and querying JSON documents:**

1. **Add JSON document:**

```sh
JSON.SET user:1 $ '{"name": "Alice", "age": 30, "email": "alice@example.com"}'
JSON.SET user:2 $ '{"name": "Bob", "age": 25, "email": "bob@example.com"}'
```

2. **Retrieve JSON document:**

```sh
JSON.GET user:1
```

3. **Query specific fields:**

```sh
JSON.GET user:1 $.name
```

4. **Increment a numeric field in JSON:**

```sh
JSON.NUMINCRBY user:1 $.age 1
```

### 3. Time-Series

**Example: Managing time-series data:**

1. **Create a time-series:**

```sh
TS.CREATE temperature:room1 LABELS room room1 sensor temp
```

2. **Add data points:**

```sh
TS.ADD temperature:room1 * 21.5
TS.ADD temperature:room1 * 22.0
```

3. **Retrieve recent data points:**

```sh
TS.RANGE temperature:room1 - +
```

4. **Retrieve aggregated data:**

```sh
TS.RANGE temperature:room1 - + AGGREGATION avg 60000
```
