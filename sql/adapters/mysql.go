package adapters

import (
	"strconv"
	"strings"

	sqb "github.com/wrmsr/bane/sql/base"
)

func FormatMysqlDsn(dsn sqb.Dsn) string {
	var b strings.Builder
	if dsn.User != "" {
		b.WriteString(dsn.User)
		if dsn.Pass.Value() != "" {
			b.WriteString(":")
			b.WriteString(dsn.Pass.Value())
		}
		b.WriteString("@")
	}
	b.WriteString("tcp(")
	b.WriteString(dsn.Host)
	if dsn.Port != 0 {
		b.WriteString(":")
		b.WriteString(strconv.Itoa(dsn.Port))
	}
	b.WriteString(")/")
	return b.String()
}
