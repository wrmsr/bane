package codecs

import (
	"bytes"
	"compress/bzip2"
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

type Bzip2Codec struct{}

var _ Codec[[]byte, []byte] = Bzip2Codec{}

func (c Bzip2Codec) Encode(f []byte) []byte {
	//b := bytes.NewBuffer(nil)
	//w := bzip2.(b)
	//check.Must1(w.Write(f))
	//check.Must(w.Close())
	//return b.Bytes()
	panic("???")
}

func (c Bzip2Codec) Decode(t []byte) []byte {
	r := bzip2.NewReader(bytes.NewReader(t))
	return check.Must1(io.ReadAll(r))
}
