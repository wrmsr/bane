package main

import (
	"bufio"
	"fmt"
	"os"
)

//type e = [10_000_000_000]int

//go:noinline
//func barf(c chan e) {
//	var z e
//	c <- z
//}

func main() {
	//c := make(chan e, 1)
	//barf(c)
	bufio.NewScanner(os.Stdin).Scan()
	a := make([]int64, 100_000_000_000)
	fmt.Println(len(a))
	bufio.NewScanner(os.Stdin).Scan()
	a[50_000_000_001] = int64(420)
	bufio.NewScanner(os.Stdin).Scan()
}
