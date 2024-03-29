package sql

import (
	"context"
	"database/sql"
	"fmt"
	"testing"

	"github.com/wrmsr/bane/core/check"
	"github.com/wrmsr/bane/core/dev"
	inj "github.com/wrmsr/bane/core/inject"
	rfl "github.com/wrmsr/bane/core/reflect"
	log "github.com/wrmsr/bane/core/slog"
	bt "github.com/wrmsr/bane/core/types"
	sqa "github.com/wrmsr/bane/sql/adapters"
	sqb "github.com/wrmsr/bane/sql/base"
)

func TestAll(t *testing.T) {
	dsn := dev.Provide(inj.Tag(rfl.TypeOf[bt.Optional[sqb.Dsn]](), "mysql")).(bt.Optional[sqb.Dsn]).Value()

	sdb := check.Must1(sql.Open("mysql", sqa.FormatMysqlDsn(dsn)))
	defer log.OrError(sdb.Close)

	db := sqa.NewStdDb(sdb)
	_ = db

	ctx := context.Background()
	fmt.Println(check.Must1(All(ctx, db, "select version() union select 'foo'")))

	fmt.Println(check.Must1(Scalar(ctx, db, "select 420")))
}
