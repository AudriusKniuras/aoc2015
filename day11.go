package main

import (
	"fmt"
)

func strToAscii(s string) []int {
	var arr = make([]int, len(s))
	for i := 0; i < len(s); i++ {
		// 97, 98, 99 (abc)
		arr[i] = int(s[i])
	}
	return arr
}

func incrementString(s string) string {
	arr := strToAscii(s)
	newString := ""

	for i := len(arr) - 1; i >= 0; i-- {
		arr[i]++
		// break immediately if we can increment
		if arr[i] <= 122 {
			break
		}
		// if last number and we couldn't increment, set it to "a", append "a" to the end and break
		if i == 0 {
			arr = append(arr, 97)
			arr[i] = 97
			break
		}
		// if we couldn't increment but it's not the last number, just set it to "a"
		arr[i] = 97
	}
	for i := 0; i < len(arr); i++ {
		newString += string(arr[i])
	}
	return newString
}

func checkStraight(s string) bool {
	arr := strToAscii(s)
	lastNum := 0
	streak := 1
	for _, num := range arr {
		if num-1 == lastNum {
			streak++
		} else {
			streak = 1
		}
		if streak == 3 {
			return true
		}
		lastNum = num
	}
	return false
}

func checkIllegal(s string) bool {
	// i - 105; l - 108; o - 111
	arr := strToAscii(s)
	for _, num := range arr {
		if num == 105 || num == 108 || num == 111 {
			return false
		}
	}
	return true
}

func checkPairs(s string) bool {
	// check aacc - 2 pairs
	// aaa - 1 pair
	arr := strToAscii(s)
	currNum := arr[0]
	pairs := 0

	for i := 1; i < len(arr); i++ {
		if arr[i] == currNum {
			pairs++
			currNum = -1
		} else {
			currNum = arr[i]
		}
	}
	if pairs == 2 {
		return true
	} else {
		return false
	}
}

func day11() {
	input := "hxbxxyzz"
	for {
		input = incrementString(input)
		if checkIllegal(input) && checkStraight(input) && checkPairs(input) {
			fmt.Println(input)
			break
		}
	}
}
