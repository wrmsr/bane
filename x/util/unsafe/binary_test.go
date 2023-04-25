package unsafe

import (
	"testing"

	tu "github.com/wrmsr/bane/pkg/util/dev/testing"
	"github.com/wrmsr/bane/pkg/util/platform"
)

func TestNativeEndian(t *testing.T) {
	pne := platform.NativeEndian()
	une := NativeEndian()
	var pnb [2]byte
	var unb [2]byte
	pne.PutUint16(pnb[:], 0xABCD)
	une.PutUint16(unb[:], 0xABCD)
	tu.AssertDeepEqual(t, pnb, unb)
}
