package main

type SomeStruct struct {
	S string
	I int
}

func main() {
	var s *SomeStruct
	_ = &s.S
}
