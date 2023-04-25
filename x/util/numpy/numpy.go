package numpy

import (
	"encoding/binary"
	"io"
	"math"
	"os"

	"github.com/wrmsr/bane/pkg/util/check"
	"github.com/wrmsr/bane/pkg/util/slices"
	log "github.com/wrmsr/bane/pkg/util/slog"
	"github.com/wrmsr/bane/x/util/python"
)

func ShittyReadFloat32s(pth string) []float32 {
	f := check.Must1(os.Open(pth))
	defer log.OrError(f.Close)

	hdrsz := 0xFF
	ohdr := make([]byte, hdrsz)
	hdr := ohdr
	_ = check.Must1(f.Read(hdr))

	read := func(n int) []byte {
		ret := hdr[:n]
		hdr = hdr[n:]
		return ret
	}

	check.Ok(slices.Equal(read(6), []byte("\x93NUMPY")))

	check.Equal(read(1)[0], byte(1))
	check.Equal(read(1)[0], byte(0))

	hl := binary.LittleEndian.Uint16(read(2))

	pyhdr := string(read(int(hl)))
	_ = python.ParseNode(pyhdr)

	ofs := int(hl) + 10

	dl := int(check.Must1(os.Stat(pth)).Size()) - ofs
	check.Equal(dl%4, 0)

	check.Must1(f.Seek(int64(ofs), io.SeekStart))

	b := make([]byte, dl)
	check.Must1(f.Read(b))

	ret := make([]float32, dl/4)
	for i := 0; i < dl; i += 4 {
		b := binary.LittleEndian.Uint32(b[i : i+4])
		ret[i/4] = math.Float32frombits(b)
	}

	return ret
}
