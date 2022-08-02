package main

//go:noinline
func add2(x, y int) int {
	return x + y
}

func main() {
	println(add2(2, 3))
}
