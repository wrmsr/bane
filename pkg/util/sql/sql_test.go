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
	rfl "github.com/wrmsr/bane/pkg/util/reflect"
	sqa "github.com/wrmsr/bane/pkg/util/sql/adapters"
	sqb "github.com/wrmsr/bane/pkg/util/sql/base"
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

	m := check.Must1(ScanMap(
		[]sqb.Column{{Name: "foo", Type: rfl.TypeOf[string]()}},
		db.QueryRow("SELECT VERSION()").Scan,
	))
	fmt.Println(m)

	m = check.Must1(ScanMapAny(
		[]string{"foo"},
		db.QueryRow("SELECT VERSION()").Scan,
	))
	fmt.Println(m)

	ctx := context.Background()

	d := sqa.NewStdDb(db)

	c := check.Must1(d.Connect(ctx))
	defer log.OrError(c.Close)

	r := check.Must1(c.Query(sqb.Query{
		Ctx:  ctx,
		Mode: sqb.QueryQuery,
		Text: "select version() union select 'foo'",
	}))
	defer log.OrError(r.Close)

	//for r.Err() == nil && r.Next() {
	//	var s string
	//	check.Must(r.Scan(&s))
	//	fmt.Println(s)
	//}
	//check.Must(r.Err())

	for ri := Iter(r); ri.HasNext(); {
		re := ri.Next()
		check.Must(re.Err)
		fmt.Println(re.Val)
	}

	//Exec(context.Background(), )
	fmt.Println(version)
}

func TestSqlite(t *testing.T) {
	db := check.Must1(sql.Open("sqlite3", "file::memory:?cache=shared"))
	defer log.OrError(db.Close)

	var version string
	check.Must(db.QueryRow("SELECT sqlite_version()").Scan(&version))

	fmt.Println(version)
}
