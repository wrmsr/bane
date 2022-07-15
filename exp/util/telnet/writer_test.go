package telnet

import (
	"bytes"
	"io/ioutil"
	"testing"

	"github.com/wrmsr/bane/pkg/util/check"
	tu "github.com/wrmsr/bane/pkg/util/dev/testing"
)

func TestWriter(t *testing.T) {
	b := bytes.NewBuffer(nil)
	w := &Writer{w: b}
	check.Must1(w.Write([]byte{41, 42, 43, IAC, 44}))
	r := check.Must1(ioutil.ReadAll(b))
	tu.AssertDeepEqual(t, r, []byte{41, 42, 43, IAC, IAC, 44})
}
