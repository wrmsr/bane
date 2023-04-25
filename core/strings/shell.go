package strings

import (
	"fmt"
	"strings"
)

func ShellQuote(s string) string {
	return fmt.Sprintf("'%s'", strings.ReplaceAll(s, "'", "'\"'\"'"))
}
