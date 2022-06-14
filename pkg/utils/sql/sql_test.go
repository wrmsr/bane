package sql

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jackc/pgx/v4/pgxpool"
	_ "github.com/mattn/go-sqlite3"
)

func TestPgx(t *testing.T) {
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

func TestMysql(t *testing.T) {
	db, err := sql.Open("mysql", "bane:bane@tcp(127.0.0.1:32220)/")
	defer db.Close()

	if err != nil {
		log.Fatal(err)
	}

	var version string

	err2 := db.QueryRow("SELECT VERSION()").Scan(&version)

	if err2 != nil {
		log.Fatal(err2)
	}

	fmt.Println(version)
}

func TestSqlite(t *testing.T) {
	db, err := sql.Open("sqlite3", "file::memory:?cache=shared")
	defer db.Close()

	if err != nil {
		log.Fatal(err)
	}

	var version string

	err2 := db.QueryRow("SELECT sqlite_version()").Scan(&version)

	if err2 != nil {
		log.Fatal(err2)
	}

	fmt.Println(version)
}
