package container

import (
	"fmt"
	"strings"

	iou "github.com/wrmsr/bane/pkg/util/io"
)

//

func (m orderedMapImpl[K, V]) format(tn string, f fmt.State, c rune) {
	iou.WriteStringDiscard(f, tn)
	if f.Flag('#') {
		iou.WriteStringDiscard(f, "{")
	} else {
		iou.WriteStringDiscard(f, "[")
	}
	for i, kv := range m.s {
		if i > 0 {
			if f.Flag('#') {
				iou.WriteStringDiscard(f, ", ")
			} else {
				iou.WriteStringDiscard(f, " ")
			}
		}
		if f.Flag('#') {
			iou.FprintfDiscard(f, "%#v:%#v", kv.K, kv.V)
		} else {
			iou.FprintfDiscard(f, "%v:%v", kv.K, kv.V)
		}
	}
	if f.Flag('#') {
		iou.WriteStringDiscard(f, "}")
	} else {
		iou.WriteStringDiscard(f, "]")
	}
}

func (m orderedMapImpl[K, V]) Format(f fmt.State, c rune) {
	m.format("ordMap", f, c)
}

func (m orderedMapImpl[K, V]) string(tn string) string {
	var sb strings.Builder
	sb.WriteString("ordMap")
	for i, kv := range m.s {
		if i > 0 {
			sb.WriteRune(' ')
		}
		iou.FprintfDiscard(&sb, "%v:%v", kv.K, kv.V)
	}
	return sb.String()
}

func (m orderedMapImpl[K, V]) String() string {
	return m.string("ordMap")
}

//

func (m mutOrderedMapImpl[K, V]) String() string {
	return m.string("mutOrdMap")
}

func (m mutOrderedMapImpl[K, V]) Format(f fmt.State, c rune) {
	m.format("mutOrdMap", f, c)
}
