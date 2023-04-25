package telnet

import (
	"bufio"
	"bytes"
	"io/ioutil"
	"testing"

	"github.com/wrmsr/bane/core/check"
	tu "github.com/wrmsr/bane/core/dev/testing"
)

func TestReader(t *testing.T) {
	b := []byte{40, 41, 42, IAC, IAC, 43}
	r := &Reader{b: bufio.NewReader(bytes.NewReader(b))}
	b2 := check.Must1(ioutil.ReadAll(r))
	tu.AssertDeepEqual(t, b2, []byte{40, 41, 42, IAC, 43})
}
