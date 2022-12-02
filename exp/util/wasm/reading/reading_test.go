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

func (r *ByteReader) Tell() int64 { return r.i }

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

func (r *ByteReader) Skip(n int64) error {
	if (r.i + n) > int64(len(r.s)) {
		return io.EOF
	}
	r.i += n
	return nil
}

func NewByteReader(b []byte) *ByteReader {
	return &ByteReader{b, 0}
}

//

type ModuleReader struct {
	r *ByteReader
}

func (r *ModuleReader) readByte() byte {
	return check.Must1(r.r.ReadByte())
}

func (r *ModuleReader) readI64() int64 {
	return leb128.DecodeI64(r.readByte)
}

func (r *ModuleReader) readU64() uint64 {
	return leb128.DecodeU64(r.readByte)
}

func (r *ModuleReader) readI32le() int32 {
	var v int32
	check.Must(binary.Read(r.r, binary.LittleEndian, &v))
	return v
}

func (r *ModuleReader) ReadModule() {
	magic := r.readI32le()
	check.Equal(magic, consts.Magic)

	version := r.readI32le()
	check.Equal(version, consts.Version)

	for r.r.Len() > 0 {
		r.readSection()
	}

	check.Equal(r.r.Len(), 0)
}

func (r *ModuleReader) readSection() {
	secTy := r.readByte()
	secSz := r.readU64()

	switch secTy {

	case consts.TypeSection:
		numSigs := int(r.readU64())
		for i := 0; i < numSigs; i++ {
			ty := r.readByte()
			check.Equal(ty, consts.FuncType)

			numParam := int(r.readU64())
			for j := 0; j < numParam; j++ {
				pty := r.readI64()
				_ = pty
			}

			numResult := int(r.readU64())
			for j := 0; j < numResult; j++ {
				rty := r.readI64()
				_ = rty
			}
		}

	default:
		fmt.Printf("unhandled section: %v %v\n", secTy, secSz)
		check.Must(r.r.Skip(int64(secSz)))

	}
}

//

func TestReading(t *testing.T) {
	src := check.Must1(os.ReadFile(filepath.Join(paths.FindProjectRoot(), "exp", "util", "wasm", "test", "go", "main_tiny.wasm")))
	r := &ModuleReader{r: NewByteReader(src)}
	r.ReadModule()

	////
	//
	//secTy = rf()
	//fmt.Println(secTy)
	//check.Equal(secTy, consts.ImportSection)
	//
	//secSz = leb128.DecodeU64(rf)
	//fmt.Println(secSz)
	//
	//numImps := int(leb128.DecodeU64(rf))
	//fmt.Println(numImps)
	//
	//readStr := func() string {
	//	l := int(leb128.DecodeU64(rf))
	//	b := make([]byte, l)
	//	check.Must1(r.Read(b))
	//	return string(b)
	//}
	//
	//for i := 0; i < numImps; i++ {
	//	modName := readStr()
	//	fmt.Println(modName)
	//	fldName := readStr()
	//	fmt.Println(fldName)
	//
	//	k := check.Must1(r.ReadByte())
	//	fmt.Println(k)
	//
	//	switch k {
	//	case consts.FuncImport:
	//		idx := leb128.DecodeU64(rf)
	//		fmt.Println(idx)
	//	default:
	//		panic(k)
	//	}
	//}
	//
	////
	//
	//secTy = rf()
	//fmt.Println(secTy)
	//check.Equal(secTy, consts.FunctionSection)
	//
	//secSz = leb128.DecodeU64(rf)
	//fmt.Println(secSz)
	//
	//numFuncs := int(leb128.DecodeU64(rf))
	//fmt.Println(numFuncs)
	//
	//for i := 0; i < numFuncs; i++ {
	//	sigIdx := leb128.DecodeU64(rf)
	//	fmt.Println(sigIdx)
	//}
	//
	////
	//
	//secTy = rf()
	//fmt.Println(secTy)
	//check.Equal(secTy, consts.TableSection)
	//
	//secSz = leb128.DecodeU64(rf)
	//fmt.Println(secSz)
	//
	////numTables := int(leb128.DecodeU64(rf))
	////fmt.Println(numTables)
	////
	////for i := 0; i < numTables; i++ {
	////
	////}
	//
	//check.Must(r.Skip(int64(secSz)))
	//
	////
	//
	//secTy = rf()
	//fmt.Println(secTy)
	//check.Equal(secTy, consts.MemorySection)
	//
	//secSz = leb128.DecodeU64(rf)
	//fmt.Println(secSz)
	//check.Must(r.Skip(int64(secSz)))
	//
	////
	//
	//secTy = rf()
	//fmt.Println(secTy)
	//check.Equal(secTy, consts.GlobalSection)
	//
	//secSz = leb128.DecodeU64(rf)
	//fmt.Println(secSz)
	//check.Must(r.Skip(int64(secSz)))
	//
	////
	//
	//secTy = rf()
	//fmt.Println(secTy)
	//check.Equal(secTy, consts.ExportSection)
	//
	//secSz = leb128.DecodeU64(rf)
	//fmt.Println(secSz)
	//check.Must(r.Skip(int64(secSz)))
	//
	////
	//
	//secTy = rf()
	//fmt.Println(secTy)
	//check.Equal(secTy, consts.ElementSection)
	//
	//secSz = leb128.DecodeU64(rf)
	//fmt.Println(secSz)
	//check.Must(r.Skip(int64(secSz)))
	//
	////
	//
	//secTy = rf()
	//fmt.Println(secTy)
	//check.Equal(secTy, consts.DataCountSection)
	//
	//secSz = leb128.DecodeU64(rf)
	//fmt.Println(secSz)
	//check.Must(r.Skip(int64(secSz)))
	//
	////
	//
	//secTy = rf()
	//fmt.Println(secTy)
	//check.Equal(secTy, consts.CodeSection)
	//
	//secSz = leb128.DecodeU64(rf)
	//fmt.Println(secSz)
	////check.Must(r.Skip(int64(secSz)))
	//
	//numFuncBodies := int(leb128.DecodeU64(rf))
	//fmt.Println(numFuncBodies)
	//
	//readTy := func() int {
	//	lty := int(leb128.DecodeI64(rf))
	//	fmt.Println(lty)
	//
	//	return lty
	//}
	//
	//for i := 0; i < numFuncBodies; i++ {
	//	bodySize := int(leb128.DecodeU64(rf))
	//	fmt.Println(bodySize)
	//
	//	startPos := r.Tell()
	//	endPos := startPos + int64(bodySize)
	//
	//	numLocals := int(leb128.DecodeU64(rf))
	//	fmt.Println(numLocals)
	//
	//	for j := 0; j < numLocals; j++ {
	//		numLocalTys := int(leb128.DecodeU64(rf))
	//		fmt.Println(numLocalTys)
	//
	//		readTy()
	//	}
	//
	//	for r.Tell() < endPos {
	//		o := int(check.Must1(r.ReadByte()))
	//
	//		if o == consts.MathPrefix || o == consts.SimdPrefix {
	//			o = (o << 8) | int(check.Must1(r.ReadByte()))
	//		}
	//
	//		fmt.Println(o)
	//
	//		switch o {
	//		case consts.Block:
	//			panic(o)
	//
	//		default:
	//			panic(o)
	//		}
	//	}
	//}
	//
	////
	//
	//secTy = rf()
	//fmt.Println(secTy)
	//check.Equal(secTy, consts.DataSection)
	//
	//secSz = leb128.DecodeU64(rf)
	//fmt.Println(secSz)
	//check.Must(r.Skip(int64(secSz)))
	//
	////
	//
	//for r.Len() > 0 {
	//	secTy = rf()
	//	fmt.Println(secTy)
	//	check.Equal(secTy, consts.CustomSection)
	//
	//	secSz = leb128.DecodeU64(rf)
	//	fmt.Println(secSz)
	//	check.Must(r.Skip(int64(secSz)))
	//}
}
