package main

import "fmt"

func main() {
	x := 3
	y := 4
	z := 5

	//a1, b1 := max(x, y)
	a := max(x, z)

	fmt.Printf("max(%d,%d) = %d\n", x, y, a)
	//fmt.Printf("max(%d,%d) = %d\n", x, z, max_xz)
	//fmt.Printf("max(%d,%d) = %d\n", y, z, max(y, z))
}

func max(a, b int) int {
	if a > b && a > 4 {
		return a
	}
	return b
}
