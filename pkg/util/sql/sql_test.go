package sql

import (
	"context"
	"database/sql"
	"fmt"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jackc/pgx/v4/pgxpool"
	_ "github.com/mattn/go-sqlite3"

	"github.com/wrmsr/bane/pkg/util/check"
	"github.com/wrmsr/bane/pkg/util/log"
)

func TestPgx(t *testing.T) {
	databaseUrl := "postgres://bane:bane@localhost:33222/postgres"

	dbPool := check.Must1(pgxpool.Connect(context.Background(), databaseUrl))
	defer dbPool.Close()

	rows := check.Must1(dbPool.Query(context.Background(), "select 1"))
	for rows.Next() {
		values := check.Must1(rows.Values())
		log.Printf("%v\n", values)
	}
}

func TestMysql(t *testing.T) {
	db := check.Must1(sql.Open("mysql", "bane:bane@tcp(127.0.0.1:33221)/"))
	defer log.OrError(db.Close)

	var version string
	check.Must(db.QueryRow("SELECT VERSION()").Scan(&version))

	fmt.Println(version)
}

func TestSqlite(t *testing.T) {
	db := check.Must1(sql.Open("sqlite3", "file::memory:?cache=shared"))
	defer log.OrError(db.Close)

	var version string
	check.Must(db.QueryRow("SELECT sqlite_version()").Scan(&version))

	fmt.Println(version)
}
