package main

import "fmt"

func main() {
	numbers := make([]int, 0, 5)
	printSlice(numbers)
	ints := append(numbers, 3, 4, 5, 6, 7, 8)

	for _, num := range ints {
		fmt.Println(num)
	}

	for i, i2 := range ints[3:] {
		fmt.Println(i, i2)
	}

}

func printSlice(x []int) {
	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
}
