package main

import "fmt"

func main() {
	x := 3
	fmt.Println("x=", x)

	//值copy
	x1 := add(x)
	fmt.Println("x1=", x1)
	//指针
	x2 := add1(&x)
	fmt.Println("x2=", x2)

	fmt.Println("x=", x)
}

// 值copy
func add(a int) int {
	a = a + 1
	return a
}

// 传递指针
func add1(a *int) int {
	*a = *a + 1
	return *a
}
