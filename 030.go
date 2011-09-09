package main

import (
	"math"
	"strconv"
)

func Euler030() {
	const (
		power float64 = 5
	)
	var (
		can []int
		sum float64
		one int
		str string
		sumStr string
	)
	for i := 10; i < 200000; i++ {
		sum = 0
		str = strconv.Itoa(i)
		for _, v := range str {
			one, _ = strconv.Atoi(string(v))
			sum += math.Pow(float64(one), power)
		}
		if sumStr = strconv.Ftoa64(sum, 'f', 0); str == sumStr {
			can = append(can, i)
		}
	}
	result(Sum(can))
}
