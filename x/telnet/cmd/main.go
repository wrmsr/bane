package main

import (
	"github.com/wrmsr/bane/x/telnet"
)

func main() {
	s := &telnet.Server{}
	_ = s.ListenAndServe(":5556")
}
