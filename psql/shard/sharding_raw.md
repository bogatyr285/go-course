

```sh
docker exec -it postgres_master psql -U master_user -d master_db
```

```sql
CREATE TABLE sales (
    id serial PRIMARY KEY,
    sales_date date NOT NULL,
    amount numeric NOT NULL
);

CREATE EXTENSION postgres_fdw;
```


```sql
CREATE SERVER shard_server FOREIGN DATA WRAPPER postgres_fdw
OPTIONS (host 'postgres_shard', dbname 'shard_db', port '5432');

CREATE USER MAPPING FOR master_user SERVER shard_server
OPTIONS (user 'shard_user', password 'shard_password');
```


```sql
CREATE FOREIGN TABLE shard_sales (
    id serial,
    sales_date date,
    amount numeric
) SERVER shard_server
OPTIONS (table_name 'sales');
```

`postgres_shard` Container

```sh
docker exec -it postgres_shard psql -U shard_user -d shard_db
```



```sql
CREATE TABLE sales (
    id serial PRIMARY KEY,
    sales_date date NOT NULL,
    amount numeric NOT NULL
);

CREATE EXTENSION postgres_fdw;
```


 Insert data into `shard_sales` from `postgres_master`:

```sh
docker exec -it postgres_master psql -U master_user -d master_db
```



```sql
INSERT INTO shard_sales (sales_date, amount) VALUES ('2023-05-12', 150.00);
INSERT INTO shard_sales (sales_date, amount) VALUES ('2023-05-13', 200.00);
```


```sh
docker exec -it postgres_shard psql -U shard_user -d shard_db
```



```sql
SELECT * FROM sales;
```


## Additional 
1. Set up partitioning logic to distribute data between the master and shard nodes:

   ```sql
   CREATE OR REPLACE FUNCTION insert_sales_trigger()
   RETURNS TRIGGER AS $$
   BEGIN
       IF (NEW.amount < 500) THEN
           -- Insert into local master table
           INSERT INTO sales (sales_date, amount) VALUES (NEW.sales_date, NEW.amount);
       ELSE
           -- Insert into shard table
           INSERT INTO shard_sales (sales_date, amount) VALUES (NEW.sales_date, NEW.amount);
       END IF;
       RETURN NULL;
   END;
   $$ LANGUAGE plpgsql;

   CREATE TRIGGER sales_insert_trigger
   BEFORE INSERT ON sales
   FOR EACH ROW EXECUTE FUNCTION insert_sales_trigger();
   ```