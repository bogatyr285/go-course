



### Step 1: Prepare the Environment

1. **Create the table and insert initial data:**

   Execute the following SQL commands:

   ```sql
    create database tx_test;
    CREATE DATABASE
    postgres=# \c test

   CREATE TABLE accounts (
       id SERIAL PRIMARY KEY,
       balance DECIMAL
   );

   INSERT INTO accounts (balance) VALUES (100), (200);
   ```

### Demonstrating "Read Committed" Isolation Level

### Step 2: Session 1 starts a transaction
   - Open Terminal 1 and connect to the PostgreSQL database.
   - Start a transaction and set it to the default "Read Committed" isolation level (it is the default level).

   ```sql
   BEGIN;
   -- Read the balance
   SELECT * FROM accounts WHERE id = 1;
   ```

   You should see:
   ```
   id | balance
   ----+--------
   1  | 100
   ```

### Step 3: Session 2 starts a transaction and updates the balance
   - Open Terminal 2 and connect to the PostgreSQL database.
   - Start a transaction and update one of the rows.

   ```sql
   BEGIN;
   UPDATE accounts SET balance = balance + 50 WHERE id = 1;
   COMMIT;
   ```

### Step 4: Session 1 re-reads the balance
   - Back in Terminal 1, read the balance again within the same transaction.

   ```sql
   SELECT * FROM accounts WHERE id = 1;
   ```

   You should see the updated balance due to the "Read Committed" isolation level allowing non-repeatable reads:
   ```
   id | balance
   ----+--------
   1  | 150
   ```

### Illustrating "Repeatable Read" Isolation Level

### Step 5: Session 1 starts a new transaction with "Repeatable Read" isolation
   - In Terminal 1, commit the previous transaction, start a new one, and set isolation level to "Repeatable Read".

   ```sql
   COMMIT;
   BEGIN;
   SET TRANSACTION ISOLATION LEVEL REPEATABLE READ;
   ```

   Check the initial balance again:
   ```sql
   SELECT * FROM accounts WHERE id = 1;
   ```

   You should see:
   ```
   id | balance
   ----+--------
   1  | 150
   ```

### Step 6: Session 2 updates the balance again
   - In Terminal 2, update the balance again.

   ```sql
   BEGIN;
   UPDATE accounts SET balance = balance - 25 WHERE id = 1;
   COMMIT;
   ```

### Step 7: Session 1 re-reads the balance
   - Back in Terminal 1, read the balance again.

   ```sql
   SELECT * FROM accounts WHERE id = 1;
   ```

   You should see the same balance as before the update by Session 2, because "Repeatable Read" prevents seeing changes made by other transactions after the transaction has started:
   ```
   id | balance
   ----+--------
   1  | 150
   ```

### Step 8: Clean Up (optional)
   - You can commit or rollback the transaction in Terminal 1 and close both sessions if desired.

   ```sql
   COMMIT;
   ```

## LEVEL SERIALIZABLE

#### Session 1: Set Serializable Isolation Level

1. In Terminal 1, commit the previous transaction, start a new one, and set isolation level to "Serializable".

   ```sql
   COMMIT;
   BEGIN;
   SET TRANSACTION ISOLATION LEVEL SERIALIZABLE;
   ```

2. Query the balance to see the initial state of accounts:

   ```sql
   SELECT * FROM accounts WHERE id = 1;
   ```

   You should see:
   ```
   id | balance
   ----+--------
   1  | 150
   ```

#### Session 2: Update Balance

1. In Terminal 2, try to update the balance as before.

   ```sql
   BEGIN;
   UPDATE accounts SET balance = balance + 100 WHERE id = 1;
   COMMIT;
   ```

#### Session 1: Attempt another update

1. Back in Terminal 1, try to update the balance as well.

   ```sql
   UPDATE accounts SET balance = balance - 50 WHERE id = 1;
   ```

   At this point, you might encounter a serialization error instead of updating successfully. The error will look something like this:

   ```
   ERROR:  could not serialize access due to concurrent update
   ```

   This happens because PostgreSQL's Serializable isolation maintains a logical sequence of transactions that could have been executed serially, which didn't happen here due to conflicting updates.

### Step 10: Handling Serialization Errors

Serialization errors require transaction rollback and replay:

1. In the event of a serialization error, you need to rollback:

   ```sql
   ROLLBACK;
   ```

2. You can retry the transaction:

   ```sql
   BEGIN;
   SET TRANSACTION ISOLATION LEVEL SERIALIZABLE;
   UPDATE accounts SET balance = balance - 50 WHERE id = 1;
   COMMIT;
   ```




## REPEATABLE READ with new records

#### Scenario Setup

1. **Modify the Table:**

   We adjust the table by adding more rows for clearer conditions:
   
   ```sql
   INSERT INTO accounts (balance) VALUES (3000), (4000), (5000);
   ```

2. **Initial Select Setup**

   Let's create a scenario where we select a range of balances and modify them:

#### Session 1: Start a Repeatable Read Transaction

```sql
BEGIN TRANSACTION ISOLATION LEVEL REPEATABLE READ;

-- Execute a range query
SELECT * FROM accounts WHERE balance > 2000;
```

#### Session 2: Insert/Update affecting the same range

In another session, we will perform an operation that would conceptually change the outcome of Session 1’s 

```sql
BEGIN;

-- Insert a new row that would fall into the range queried by Session 1
INSERT INTO accounts (balance) VALUES (2500);

-- Commit the transaction
COMMIT;
```

#### Session 1: Repeat the Read

Go back to Session 1 and repeat the range query:

```sql
-- Repeat the range query
SELECT * FROM accounts WHERE balance > 2000;
```