package sql

import (
	"context"
	"database/sql"
	"fmt"
	"testing"

	"github.com/wrmsr/bane/pkg/util/check"
	"github.com/wrmsr/bane/pkg/util/dev"
	inj "github.com/wrmsr/bane/pkg/util/inject"
	"github.com/wrmsr/bane/pkg/util/log"
	rfl "github.com/wrmsr/bane/pkg/util/reflect"
	sqa "github.com/wrmsr/bane/pkg/util/sql/adapters"
	sqb "github.com/wrmsr/bane/pkg/util/sql/base"
	bt "github.com/wrmsr/bane/pkg/util/types"
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
