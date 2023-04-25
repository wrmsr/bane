package platform

import (
	"encoding/binary"

	"golang.org/x/sys/cpu"
)

func NativeEndian() binary.ByteOrder {
	if cpu.IsBigEndian {
		return binary.BigEndian
	} else {
		return binary.LittleEndian
	}
}
