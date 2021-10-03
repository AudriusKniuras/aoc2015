package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
)

func day12() {
	file, _ := ioutil.ReadFile("day12_in.txt")
	re := regexp.MustCompile(`[-0-9]+`)
	numbers := re.FindAll(file, -1)
	sum := 0
	for _, i := range numbers {
		x, _ := strconv.Atoi(string(i))
		sum += x
	}
	fmt.Println(sum)
}
