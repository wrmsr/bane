package sql

import (
	"database/sql"
	"testing"

	"github.com/wrmsr/bane/pkg/util/check"
	"github.com/wrmsr/bane/pkg/util/dev"
	inj "github.com/wrmsr/bane/pkg/util/inject"
	"github.com/wrmsr/bane/pkg/util/log"
	opt "github.com/wrmsr/bane/pkg/util/optional"
	rfl "github.com/wrmsr/bane/pkg/util/reflect"
	sqa "github.com/wrmsr/bane/pkg/util/sql/adapters"
	sqb "github.com/wrmsr/bane/pkg/util/sql/base"
)

func TestAll(t *testing.T) {
	dsn := dev.Provide(inj.Tag(rfl.TypeOf[opt.Optional[sqb.Dsn]](), "mysql")).(opt.Optional[sqb.Dsn]).Value()

	sdb := check.Must1(sql.Open("mysql", sqa.FormatMysqlDsn(dsn)))
	defer log.OrError(sdb.Close)

	db := sqa.NewStdDb(sdb)
	_ = db
}
