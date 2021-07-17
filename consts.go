package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func main() {
	str := "hello word"
	fmt.Println(strings.HasPrefix(str, "H"))
	fmt.Println(strings.Contains(str, "w"))
	fmt.Println(strings.Index(str, "h"))
	fmt.Println(strings.LastIndex(str, "ll"))
	fmt.Println(strings.Replace(str, "l", "k", 2))
	fmt.Println(strings.Count(str, "o"))
	fmt.Println(strings.Repeat(str, 4))
	fmt.Println(strings.ToLower(str))
	fmt.Println(strings.ToUpper(str))
	fmt.Println(strings.Fields(str))
	fmt.Println(strings.Join(strings.Fields(str), "';'"))
	fmt.Println(strconv.Atoi("3"))
	fmt.Println(strconv.Itoa(3))
}

func Uint8FromInt(n int) (uint8, error) {
	if 0 <= n && n <= math.MaxUint8 {
		return uint8(n), nil
	}
	return 0, fmt.Errorf("%d is out of the uint8 range", n)
}

func IntFromFloat64(x float64) int {
	if math.MinInt32 <= x && x <= math.MaxInt32 {
		whole, fraction := math.Modf(x)
		if fraction >= 0.5 {
			whole++
		}
		return int(whole)
	}
	panic(fmt.Sprintf("%g is out of the int32 range", x))
}
