package main

import (
	"strconv"
)

func maxLen(arr []int) int {
	var mx int
	for i := 0; i < len(arr); i++ {
		strLen := len(strconv.Itoa(arr[i]))
		if strLen > mx {
			mx = strLen
		}
	}

	return mx
}

func getDigit(x, d int) int {
	a := []int{1, 10, 100, 1000, 10000, 100000, 1000000, 10000000}
	return (x / a[d]) % 10
}

func core(arr []int, digit, maxl int) []int {
	if digit >= maxl {
		return arr
	}

	radix := 10
	count := make([]int, radix)
	bucket := make([]int, len(arr))

	for i := 0; i < len(arr); i++ {
		count[getDigit(arr[i], digit)]++
	}
	for i := 1; i < radix; i++ {
		count[i] += count[i-1]
	}
	for i := len(arr) - 1; i >= 0; i-- {
		d := getDigit(arr[i], digit)
		bucket[count[d]-1] = arr[i]
		count[d]--
	}
	return core(bucket, digit+1, maxl)
}

// LSD ss
func LSD(arr []int) []int {
	return core(arr, 0, maxLen(arr))
}

func getDigit1(x int, digit int) int {
	a := []int{1, 10, 100, 1000, 10000, 100000, 1000000}
	return (x / a[digit]) % 10
}

func maxLen1(arr []int) int {
	temp := arr[0]
	for i := 0; i < len(arr); i++ {
		if arr[i] > temp {
			temp = arr[i]
		}
	}

	return len(strconv.Itoa(temp))
}

func core1(arr []int, digit, maxl int) []int {
	if digit > maxl {
		return arr
	}

	radix := 10
	count := make([]int, radix)
	bucket := make([]int, len(arr))
	for i := 0; i < radix; i++ {
		count[getDigit(arr[i], digit)]++
	}
	for i := 1; i < len(arr); i++ {
		d := getDigit(arr[i], digit)
		bucket[count[d]-1] = arr[i]
		count[d]--
	}

	return core1(bucket, digit+1, maxl)
}
