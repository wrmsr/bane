package container

import (
	"fmt"
	"strings"

	iou "github.com/wrmsr/bane/pkg/util/io"
)

//

type mapFormatter struct {
	f fmt.State
}

func (mf mapFormatter) WriteString(s string) {
	iou.WriteStringDiscard(mf.f, s)
}

func (mf mapFormatter) WriteBegin() {
	if mf.f.Flag('#') {
		mf.WriteString("{")
	} else {
		mf.WriteString("[")
	}
}

func (mf mapFormatter) WriteEntry(i int, k, v any) {
	if i > 0 {
		if mf.f.Flag('#') {
			mf.WriteString(", ")
		} else {
			mf.WriteString(" ")
		}
	}
	if mf.f.Flag('#') {
		iou.FprintfDiscard(mf.f, "%#v:%#v", k, v)
	} else {
		iou.FprintfDiscard(mf.f, "%v:%v", k, v)
	}
}

func (mf mapFormatter) WriteEnd() {
	if mf.f.Flag('#') {
		mf.WriteString("}")
	} else {
		mf.WriteString("]")
	}
}

//

type mapStringer struct {
	sb *strings.Builder
}

func (ms mapStringer) WriteBegin() {
	ms.sb.WriteString("[")
}

func (ms mapStringer) WriteEnd() {
	ms.sb.WriteString("]")
}

func (ms mapStringer) WriteEntry(i int, k, v any) {
	if i > 0 {
		ms.sb.WriteRune(' ')
	}
	iou.FprintfDiscard(ms.sb, "%v:%v", k, v)
}
