package container

import (
	"fmt"
	"strings"

	iou "github.com/wrmsr/bane/pkg/util/io"
)

//

func (m SliceMap[K, V]) format(tn string, f fmt.State, c rune) {
	iou.WriteStringDiscard(f, tn)
	mf := mapFormatter{f: f}
	mf.WriteBegin()
	for i, kv := range m.s {
		mf.WriteEntry(i, kv.K, kv.V)
	}
	mf.WriteEnd()
}

func (m SliceMap[K, V]) Format(f fmt.State, c rune) {
	m.format("ordMap", f, c)
}

func (m SliceMap[K, V]) string(tn string) string {
	var sb strings.Builder
	sb.WriteString(tn)
	ms := mapStringer{sb: &sb}
	ms.WriteBegin()
	for i, kv := range m.s {
		ms.WriteEntry(i, kv.K, kv.V)
	}
	ms.WriteEnd()
	return sb.String()
}

func (m SliceMap[K, V]) String() string {
	return m.string("ordMap")
}

//

func (m MutSliceMap[K, V]) String() string {
	return m.m.string("mutOrdMap")
}

func (m MutSliceMap[K, V]) Format(f fmt.State, c rune) {
	m.m.format("mutOrdMap", f, c)
}
