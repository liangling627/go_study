package main

import (
	"fmt"
)

type testInt func(int) bool

func isOld(integer int) bool {
	if integer%2 == 0 {
		return false
	}
	return true
}

func isEven(integer int) bool {
	if integer%2 == 0 {
		return true
	}

	return false
}

func filter(silce []int, f testInt) []int {
	var result []int
	for _, value := range silce {
		if f(value) {
			result = append(result, value)
		}
		// 直接调用函数
		//if isEven(value) {
		//	result = append(result, value)
		//}
	}

	return result
}

func main() {
	slice := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println("slice = ", slice)
	old := filter(slice, isOld)
	fmt.Println("old element is slice", old)
	even := filter(slice, isEven)

	fmt.Println("even is ", even)

	// 异常
	//getenv := os.Getenv("use")
	//if getenv == "" {
	//	panic("no value for $USER ")
	//}

}
