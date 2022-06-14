package sql

import (
	"context"
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/jackc/pgx/v4/pgxpool"
)

func TestSql(t *testing.T) {
	databaseUrl := "postgres://bane:bane@localhost:32221/postgres"

	dbPool, err := pgxpool.Connect(context.Background(), databaseUrl)

	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	defer dbPool.Close()

	rows, err := dbPool.Query(context.Background(), "select 1")
	if err != nil {
		log.Fatal("error while executing query")
	}

	for rows.Next() {
		values, err := rows.Values()
		if err != nil {
			log.Fatal("error while iterating dataset")
		}

		log.Printf("%v\n", values)
	}
}
