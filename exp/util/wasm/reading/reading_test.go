package reading

import (
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"testing"

	"github.com/wrmsr/bane/exp/util/wasm/reading/consts"
	"github.com/wrmsr/bane/exp/util/wasm/reading/leb128"
	"github.com/wrmsr/bane/pkg/util/check"
	"github.com/wrmsr/bane/pkg/util/dev/paths"
)

//

type ByteReader struct {
	s []byte
	i int64
}

var _ io.Reader = &ByteReader{}

func (r *ByteReader) Len() int {
	if r.i >= int64(len(r.s)) {
		return 0
	}
	return int(int64(len(r.s)) - r.i)
}

func (r *ByteReader) Size() int64 { return int64(len(r.s)) }

func (r *ByteReader) Tell() int64 {
	return r.i
}

func (r *ByteReader) Read(b []byte) (n int, err error) {
	if r.i >= int64(len(r.s)) {
		return 0, io.EOF
	}
	n = copy(b, r.s[r.i:])
	r.i += int64(n)
	return
}

func (r *ByteReader) ReadAt(b []byte, off int64) (n int, err error) {
	if off < 0 {
		return 0, errors.New("negative offset")
	}
	if off >= int64(len(r.s)) {
		return 0, io.EOF
	}
	n = copy(b, r.s[off:])
	if n < len(b) {
		err = io.EOF
	}
	return
}

func (r *ByteReader) ReadByte() (byte, error) {
	if r.i >= int64(len(r.s)) {
		return 0, io.EOF
	}
	b := r.s[r.i]
	r.i++
	return b, nil
}

func NewByteReader(b []byte) *ByteReader {
	return &ByteReader{b, 0}
}

//

func TestReading(t *testing.T) {
	src := check.Must1(os.ReadFile(filepath.Join(paths.FindProjectRoot(), "exp", "util", "wasm", "test", "go", "main_tiny.wasm")))
	r := NewByteReader(src)

	var magic int32
	check.Must(binary.Read(r, binary.LittleEndian, &magic))
	check.Equal(magic, consts.Magic)

	var version int32
	check.Must(binary.Read(r, binary.LittleEndian, &version))
	check.Equal(version, consts.Version)

	rf := func() byte {
		return check.Must1(r.ReadByte())
	}

	secTy := rf()
	fmt.Println(secTy)
	check.Equal(secTy, consts.TypeSection)

	secSz := leb128.DecodeU64(rf)
	fmt.Println(secSz)

	numSigs := int(leb128.DecodeU64(rf))
	fmt.Println(numSigs)

	for i := 0; i < numSigs; i++ {
		ty := check.Must1(r.ReadByte())
		check.Equal(ty, consts.FuncType)

		numParam := int(leb128.DecodeU64(rf))
		fmt.Println(numParam)

		for j := 0; j < numParam; j++ {
			pty := leb128.DecodeI64(rf)
			fmt.Println(pty)
		}

		numResult := int(leb128.DecodeU64(rf))
		fmt.Println(numResult)

		for j := 0; j < numResult; j++ {
			rty := leb128.DecodeI64(rf)
			fmt.Println(rty)
		}
	}

	secTy = rf()
	fmt.Println(secTy)
	check.Equal(secTy, consts.ImportSection)

	secSz = leb128.DecodeU64(rf)
	fmt.Println(secSz)

	numImps := int(leb128.DecodeU64(rf))
	fmt.Println(numImps)

	readStr := func() string {
		l := int(leb128.DecodeU64(rf))
		b := make([]byte, l)
		check.Must1(r.Read(b))
		return string(b)
	}

	for i := 0; i < numImps; i++ {
		modName := readStr()
		fmt.Println(modName)
		fldName := readStr()
		fmt.Println(fldName)

		k := check.Must1(r.ReadByte())
		fmt.Println(k)

		switch k {
		case consts.FuncImport:
			idx := leb128.DecodeU64(rf)
			fmt.Println(idx)
		default:
			panic(k)
		}
	}

	secTy = rf()
	fmt.Println(secTy)
	check.Equal(secTy, consts.FunctionSection)

	secSz = leb128.DecodeU64(rf)
	fmt.Println(secSz)

	numFuncs := int(leb128.DecodeU64(rf))
	fmt.Println(numFuncs)

	for i := 0; i < numFuncs; i++ {
		sigIdx := leb128.DecodeU64(rf)
		fmt.Println(sigIdx)
	}

	secTy = rf()
	fmt.Println(secTy)
	check.Equal(secTy, consts.TableSection)

	secSz = leb128.DecodeU64(rf)
	fmt.Println(secSz)

	numTables := int(leb128.DecodeU64(rf))
	fmt.Println(numTables)

	for i := 0; i < numTables; i++ {

	}
}
