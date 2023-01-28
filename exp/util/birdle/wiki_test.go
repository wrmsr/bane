package birdle

import (
	"context"
	"fmt"
	"testing"

	"github.com/wrmsr/bane/pkg/util/check"
)

func TestWiki(t *testing.T) {
	ws := check.Must1(FetchWikiLinks(context.Background(), DefaultWikiBirdsPage))
	for _, s := range ws {
		fmt.Println(s)
	}
}
