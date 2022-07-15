package telnet

import (
	"bufio"
	"errors"
	"io"
)

type Reader struct {
	b *bufio.Reader
}

var _ io.Reader = &Reader{}

func (r *Reader) skipSb() error {
	for {
		b, err := r.b.ReadByte()
		if err != nil {
			return err
		}

		if b == IAC {
			pb, err := r.b.Peek(1)
			if err != nil {
				return err
			}

			switch pb[0] {
			case IAC:
				if _, err := r.b.Discard(1); err != nil {
					return err
				}

			case SE:
				if _, err := r.b.Discard(1); err != nil {
					return err
				}
				break
			}
		}
	}
}

func (r *Reader) Read(data []byte) (int, error) {
	n := 0
	push := func(b byte) {
		data[0] = b
		n++
		data = data[1:]
	}

	for len(data) > 0 {
		b, err := r.b.ReadByte()
		if err != nil {
			return n, err
		}

		if b != IAC {
			push(b)
			continue
		}

		pb, err := r.b.Peek(1)
		if err != nil {
			return n, err
		}

		switch pb[0] {
		case WILL, WONT, DO, DONT:
			if _, err := r.b.Discard(2); err != nil {
				return n, err
			}

		case IAC:
			push(IAC)
			if _, err := r.b.Discard(1); err != nil {
				return n, err
			}

		case SB:
			if err := r.skipSb(); err != nil {
				return n, err
			}

		case SE:
			if _, err := r.b.Discard(1); err != nil {
				return n, err
			}

		default:
			return n, errors.New("telnet stream error")
		}
	}

	return n, nil
}
