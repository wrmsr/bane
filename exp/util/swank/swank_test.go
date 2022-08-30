package swank

import (
	"fmt"
	"net"
	"strconv"
	"testing"

	"github.com/wrmsr/bane/pkg/util/check"
	"github.com/wrmsr/bane/pkg/util/log"
)

// https://slime.common-lisp.dev/doc/html/Communication-style.html

func TestSwank(t *testing.T) {
	readChunk := func(c net.Conn, l int) []byte {
		buffer := make([]byte, l)
		n := check.Must1(c.Read(buffer))
		if n != l {
			panic(fmt.Errorf("short read: %d", n))
		}
		return buffer
	}

	readPacket := func(c net.Conn) {
		header := readChunk(c, 6)
		l := int(check.Must1(strconv.ParseUint(string(header), 16, 64)))
		payload := readChunk(c, l)
		// read_sexp_from_string(payload)
		fmt.Println(string(payload))
	}

	handle := func(c net.Conn) {
		defer log.OrError(c.Close)

		readPacket(c)
	}

	l := check.Must1(net.Listen("tcp", ":4005"))

	for {
		conn := check.Must1(l.Accept())
		go handle(conn)
	}
}
