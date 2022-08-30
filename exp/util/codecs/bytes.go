package codecs

import (
	"bytes"
	"compress/gzip"
	"io"

	"github.com/wrmsr/bane/pkg/util/check"
)

//

type GzipCodec struct{}

var _ Codec[[]byte, []byte] = GzipCodec{}

func (c GzipCodec) Encode(f []byte) []byte {
	b := bytes.NewBuffer(nil)
	w := gzip.NewWriter(b)
	check.Must1(w.Write(f))
	check.Must(w.Close())
	return b.Bytes()
}

func (c GzipCodec) Decode(t []byte) []byte {
	r := check.Must1(gzip.NewReader(bytes.NewReader(t)))
	return check.Must1(io.ReadAll(r))
}

//
