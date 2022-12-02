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

func (r *ModuleReader) readString() string {
	l := r.readU64()
	b := make([]byte, l)
	check.Must1(r.r.Read(b))
	return string(b)
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

	case consts.ImportSection:
		numImps := int(r.readU64())
		for i := 0; i < numImps; i++ {
			modName := r.readString()
			_ = modName
			fldName := r.readString()
			_ = fldName

			k := r.readByte()
			switch k {
			case consts.FuncImport:
				idx := r.readU64()
				_ = idx
			default:
				panic(k)
			}
		}

	case consts.FunctionSection:
		numFuncs := int(r.readU64())
		for i := 0; i < numFuncs; i++ {
			sigIdx := r.readU64()
			_ = sigIdx
		}

	case consts.TableSection:
		//numTables := int(r.readU64())
		//for i := 0; i < numTables; i++ {
		//}
		check.Must(r.r.Skip(int64(secSz)))

	case consts.CodeSection:
		numFuncBodies := int(r.readU64())

		for i := 0; i < numFuncBodies; i++ {
			bodySize := int(r.readU64())

			startPos := r.r.Tell()
			endPos := startPos + int64(bodySize)

			numLocals := int(r.readU64())
			for j := 0; j < numLocals; j++ {
				numTys := int(r.readU64())
				_ = numTys
				ty := r.readI64()
				_ = ty
			}

			for r.r.Tell() < endPos {
				o := int(r.readByte())

				if o == consts.MathPrefix || o == consts.SimdPrefix {
					o = (o << 8) | int(r.readByte())
				}

				switch o {
				case consts.Block:
					panic(o)

				default:
					panic(o)
				}
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

}
