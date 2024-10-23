package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"sync"
	"testing"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/jackc/pgx/v5/stdlib"
)

const connStr = "postgres://user_rw:admin@localhost/postgres?sslmode=disable"

func SimulatedWork(db *sql.DB, wg *sync.WaitGroup) {
	defer wg.Done()
	_, err := db.ExecContext(context.Background(), "SELECT generate_series(1, 1000)")
	if err != nil {
		log.Fatalf("Error executing query: %v", err)
	}
	// log.Println("db.Stats().OpenConnections", db.Stats().InUse)
}

func SimulatedWorkWithPool(pool *pgxpool.Pool, wg *sync.WaitGroup) {
	defer wg.Done()
	_, err := pool.Exec(context.Background(), "SELECT generate_series(1, 1000)")
	if err != nil {
		log.Fatalf("Error executing query: %v", err)
	}
	// log.Println("pool", pool.Stat().TotalConns())
}

func BenchmarkWithoutPool(b *testing.B) {
	db, err := sql.Open("pgx", connStr)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v", err)
	}
	defer db.Close()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var wg sync.WaitGroup
		for j := 0; j < 100; j++ {
			wg.Add(1)
			go SimulatedWork(db, &wg)
		}

		wg.Wait()
	}
}

func BenchmarkWithPool(b *testing.B) {
	pool := getPool()

	defer pool.Close()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var wg sync.WaitGroup
		for j := 0; j < 100; j++ {
			wg.Add(1)
			go SimulatedWorkWithPool(pool, &wg)
		}
		wg.Wait()
	}

}

func BenchmarkBulkInsertCopyFrom(b *testing.B) {
	pool := getPool()

	defer pool.Close()
	pool.Exec(context.Background(), "TRUNCATE records")
	createTable(pool)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		records := make([]Record, 1000)
		for j := 0; j < 1000; j++ {
			records[j] = Record{ID: j, Name: "Item", Value: 100 + j}
		}

		err := bulkInsertCopyFrom(pool, records)
		if err != nil {
			b.Fatalf("Bulk insert failed: %v\n", err)
		}
		b.StopTimer()
		pool.Exec(context.Background(), "TRUNCATE records")
		createTable(pool)
		b.StartTimer()
	}
}

func BenchmarkBulkInsertBatch(b *testing.B) {
	pool := getPool()

	defer pool.Close()
	pool.Exec(context.Background(), "TRUNCATE records")
	createTable(pool)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		records := make([]Record, 1000)
		for j := 0; j < 1000; j++ {
			records[j] = Record{ID: j, Name: "Item", Value: 100 + j}
		}

		err := bulkInsertBatch(pool, records)
		if err != nil {
			b.Fatalf("Bulk insert failed: %v\n", err)
		}
		b.StopTimer()
		pool.Exec(context.Background(), "TRUNCATE records")
		createTable(pool)
		b.StartTimer()
	}
}

func BenchmarkBulkRead(b *testing.B) {
	pool := getPool()

	defer pool.Close()
	pool.Exec(context.Background(), "TRUNCATE records")
	createTable(pool)

	records := make([]Record, 1000)
	for j := 0; j < 1000; j++ {
		records[j] = Record{ID: j, Name: "Item", Value: 100 + j}
	}

	err := bulkInsertCopyFrom(pool, records)
	if err != nil {
		b.Fatalf("Bulk insert failed: %v\n", err)
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		readBatch(pool)
	}
}

// insert
type Record struct {
	ID    int
	Name  string
	Value int
}

func createTable(pool *pgxpool.Pool) {
	_, err := pool.Exec(context.Background(), `
        CREATE TABLE IF NOT EXISTS records (
            id INT PRIMARY KEY,
            name VARCHAR(100),
            value INT
        )
    `)
	if err != nil {
		log.Fatalf("Failed to create table: %v\n", err)
	}
}

func bulkInsertCopyFrom(pool *pgxpool.Pool, records []Record) error {
	ctx := context.Background()

	rows := make([][]interface{}, len(records))
	for i, record := range records {
		rows[i] = []interface{}{record.ID, record.Name, record.Value}
	}

	_, err := pool.CopyFrom(
		ctx,
		pgx.Identifier{"records"},
		[]string{"id", "name", "value"},
		pgx.CopyFromRows(rows),
	)
	return err
}

func bulkInsertBatch(pool *pgxpool.Pool, records []Record) error {
	ctx := context.Background()

	batch := &pgx.Batch{}
	for _, record := range records {
		batch.Queue("INSERT INTO records (id, name, value) VALUES ($1, $2, $3)", record.ID, record.Name, record.Value)
	}

	br := pool.SendBatch(ctx, batch)
	defer br.Close()

	_, err := br.Exec()
	return err
}

func readBatch(pool *pgxpool.Pool) ([]Record, error) {
	ctx := context.Background()

	query := "SELECT id, name, value FROM records"

	rows, err := pool.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("query failed: %w", err)
	}
	defer rows.Close()

	// return pgx.CollectRows(rows, pgx.RowToStructByName[Record])

	var records []Record
	// Fetch rows in batches
	for rows.Next() {
		var record Record
		err := rows.Scan(&record.ID, &record.Name, &record.Value)
		if err != nil {
			return nil, fmt.Errorf("row scan failed: %w", err)
		}
		records = append(records, record)
	}

	if rows.Err() != nil {
		return nil, fmt.Errorf("rows iteration failed: %w", rows.Err())
	}

	return records, nil
}

func getPool() *pgxpool.Pool {
	config, err := pgxpool.ParseConfig(connStr)
	if err != nil {
		log.Fatalf("Unable to parse connection string: %v", err)
	}

	config.MaxConns = 90                       // Maximum number of connections in the pool
	config.MinConns = 10                       // Minimum number of connections in the pool
	config.MaxConnIdleTime = 5 * time.Minute   // Maximum time a connection can be idle
	config.MaxConnLifetime = 30 * time.Minute  // Maximum lifetime of a connection
	config.HealthCheckPeriod = 1 * time.Minute // How often to perform health checks of idle connections

	pool, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		log.Fatalf("connect to database: %v", err)
	}
	return pool
}

func main() {
	fmt.Println("Running Benchmarks")
	result := testing.Benchmark(BenchmarkWithoutPool)
	fmt.Printf("Without Pool: %v\n", result)

	result = testing.Benchmark(BenchmarkWithPool)
	fmt.Printf("With Pool: %v\n", result)

	// result := testing.Benchmark(BenchmarkBulkInsertBatch)
	// fmt.Printf("BenchmarkBulkInsertBatch: %v\n", result)
	// result = testing.Benchmark(BenchmarkBulkInsertCopyFrom)
	// fmt.Printf("BenchmarkBulkInsertCopyFrom: %v\n", result)
}
