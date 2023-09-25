package main

import (
	"fmt"
	"math"
	"strings"
)

func base2ToBase10(num string) (float64, error) {
	var err error
	if strings.ContainsAny(num, "-23456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ") {
		err = fmt.Errorf("invalid binary number")
		return 0, err
	}
	var nSum, mSum float64
	n := strings.IndexRune(num, '.')
	if n != -1 {
		exp := float64(n - 1)
		for _, i := range num[:n] {
			if i == 49 {
				nSum += 1.0 * math.Pow(2, float64(exp))
			}
			exp--
		}
		for m, j := range num[n+1:] {
			if j == 49 {
				mSum += 1.0 * math.Pow(2, float64(-(m+1)))
			}
		}
	} else {
		exp := len(num) - 1
		for _, i := range num {
			if i == 49 {
				nSum += 1.0 * math.Pow(2, float64(exp))
			}
			exp--
		}
	}
	return nSum + mSum, nil
}

func base10ToBase2(num float64) (string, error) {
	return "", nil
}

func main() {
	a, err := base2ToBase10("101101")
	fmt.Println(a, err)
}
