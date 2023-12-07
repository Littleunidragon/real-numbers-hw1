package main

import (
	"fmt"
	"math"
	"strings"
)

func base2ToBase10(num string) (float64, error) {
	numbers := strings.Split(num, ".")
	var answer float64
	length := len([]rune(numbers[0])) - 1
	for i := range []rune(numbers[0]) {
		answer += math.Pow(2, float64(length-i))
	}
	for i := range []rune(numbers[1]) {
		answer += math.Pow(2, float64(-i-1))
	}
	return answer, nil
}

func base10ToBase2(num float64) (string, error) {
	return "", nil
}

func main() {
	a, err := base2ToBase10("101101")
	fmt.Println(a, err)
}
