package telnet

import (
	"bufio"
	"net"

	log "github.com/wrmsr/bane/core/slog"
)

type Server struct {
}

func (s *Server) ListenAndServe(addr string) error {
	l, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}
	return s.Serve(l)
}

func (s *Server) Serve(l net.Listener) error {
	defer log.OrError(l.Close)

	for {
		conn, err := l.Accept()
		if err != nil {
			return err
		}
		go s.handle(conn)
	}
}

func (s *Server) handle(c net.Conn) {
	defer log.OrError(c.Close)

	defer func() {
		if r := recover(); nil != r {
			log.Error("Recovered from telnet handler", log.P(r))
		}
	}()

	r := &Reader{b: bufio.NewReader(c)}
	w := &Writer{w: c}

	s.echo(w, r)
}

func (s *Server) echo(w *Writer, r *Reader) {
	var buffer [1]byte
	p := buffer[:]

	for {
		n, err := r.Read(p)

		if n > 0 {
			_, _ = w.Write(p[:n])
		}

		if nil != err {
			break
		}
	}
}
