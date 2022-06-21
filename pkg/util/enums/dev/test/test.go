//go:build !nodev

package main

//go:generate go run github.com/wrmsr/bane/pkg/utils/enums/dev/gen -type=RoundingMode

type RoundingMode byte

const (
	ToNearestEven RoundingMode = iota
	ToNearestAway
	ToZero
	AwayFromZero
	ToNegativeInf
	ToPositiveInf
)

func main() {

}
