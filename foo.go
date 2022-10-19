package main

/*
#include "foo.h"
*/
import "C"

import "fmt"

func main() {
	p := C.get_point()
	fmt.Println(p.x)
	fmt.Println(p.y)
}
