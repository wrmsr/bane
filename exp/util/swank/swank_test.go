package swank

import (
	"fmt"
	"io"
	"net"
	"strconv"
	"testing"

	"github.com/wrmsr/bane/pkg/util/check"
	log "github.com/wrmsr/bane/pkg/util/slog"
	bt "github.com/wrmsr/bane/pkg/util/types"
)

// https://slime.common-lisp.dev/doc/html/Communication-style.html

type Connection struct {
	n net.Conn
}

func (c *Connection) readChunk(l int) []byte {
	buffer := make([]byte, l)
	n := check.Must1(c.n.Read(buffer))
	if n != l {
		panic(fmt.Errorf("short read: %d", n))
	}
	return buffer
}

func (c *Connection) readPacket() {
	header := c.readChunk(6)
	l := int(check.Must1(strconv.ParseUint(string(header), 16, 64)))
	payload := c.readChunk(l)

	// read_sexp_from_string(payload)
	// (:emacs-rex (swank:connection-info) "COMMON-LISP-USER" t 1)
	fmt.Println(string(payload))
}

type LispReader struct {
	r io.Reader
	p bt.Optional[byte]
}

//func (r LispReader) peek() byte {
//
//}
//
//func (r LispReader) getb() byte {
//	b := []byte{0}
//	check.Equal(check.Must1(r.r.Read(b)), 1)
//}

/*
func (r LispReader) read(allowConsingDot bool) any {
    skip_whitespace
    c = @io.getc
    case c
    when ?( then read_list(true)
    when ?" then read_string
    when ?' then read_quote
    when nil then raise EOFError.new("EOF during read")
    else
      @io.ungetc(c)
      obj = read_number_or_symbol
      if obj == :"." and not allow_consing_dot
        raise "Consing-dot in invalid context"
      end
      obj
    end
}
*/

func (r LispReader) skipWhitespace() {
	for {
		//c := r.r.Read()
		//c := @io.getc
		//  case c
		//  when ?\s, ?\n, ?\t then next
		//  when nil then break
		//  else @io.ungetc(c); break
		//  end
	}
}

func TestSwank(t *testing.T) {
	handle := func(n net.Conn) {
		defer log.OrError(n.Close)
		c := &Connection{n: n}
		c.readPacket()
	}

	l := check.Must1(net.Listen("tcp", ":4005"))
	for {
		conn := check.Must1(l.Accept())
		go handle(conn)
	}
}
