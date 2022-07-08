//go:build arm64

package platform

import "encoding/binary"

func init() {
	nativeEndian = binary.LittleEndian
}
