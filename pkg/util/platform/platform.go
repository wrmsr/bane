package platform

import "encoding/binary"

var nativeEndian binary.ByteOrder

func NativeEndian() binary.ByteOrder {
	return nativeEndian
}
