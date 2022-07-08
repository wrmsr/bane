//go:build amd64

package platform

import "encoding/binary"

func init() {
	nativeEndian = binary.LittleEndian
}
