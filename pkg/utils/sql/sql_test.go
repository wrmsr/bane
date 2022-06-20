package sql

import (
	"context"
	"database/sql"
	"fmt"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jackc/pgx/v4/pgxpool"
	_ "github.com/mattn/go-sqlite3"

	"github.com/wrmsr/bane/pkg/utils/check"
	"github.com/wrmsr/bane/pkg/utils/log"
)

func TestPgx(t *testing.T) {
	databaseUrl := "postgres://bane:bane@localhost:32221/postgres"

	dbPool := check.Must(pgxpool.Connect(context.Background(), databaseUrl))
	defer dbPool.Close()

	rows := check.Must(dbPool.Query(context.Background(), "select 1"))
	for rows.Next() {
		values := check.Must(rows.Values())
		log.Printf("%v\n", values)
	}
}

func TestMysql(t *testing.T) {
	db := check.Must(sql.Open("mysql", "bane:bane@tcp(127.0.0.1:32220)/"))
	defer log.OrError(db.Close)

	var version string
	check.NoErr(db.QueryRow("SELECT VERSION()").Scan(&version))

	fmt.Println(version)
}

func TestSqlite(t *testing.T) {
	db := check.Must(sql.Open("sqlite3", "file::memory:?cache=shared"))
	defer log.OrError(db.Close)

	var version string
	check.NoErr(db.QueryRow("SELECT sqlite_version()").Scan(&version))

	fmt.Println(version)
}
