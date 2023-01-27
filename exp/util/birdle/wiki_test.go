package birdle

import (
	"context"
	"fmt"
	"testing"

	"github.com/wrmsr/bane/pkg/util/check"
)

func TestWiki(t *testing.T) {
	wp := check.Must1(FetchWikiLinks(context.Background(), DefaultWikiBirdsPage))
	fmt.Println(wp)
}
