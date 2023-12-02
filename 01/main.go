package main

import (
	"aoc/utils"
	"fmt"
	"strconv"
	"strings"
)

func number(input string) int {
	var num string
	for _, first := range input {
		if first >= '0' && first <= '9' {
			num += string(first)
			break
		}
	}
	for i := len(input) - 1; i >= 0; i-- {
		last := input[i]
		if last > '0' && last <= '9' {
			num += string(last)
			break
		}
	}
	res, _ := strconv.Atoi(num)
	return res
}

func replaceNumber(input string) string {
	input = strings.ReplaceAll(input, "oneight", "18")
	input = strings.ReplaceAll(input, "twone", "21")
	input = strings.ReplaceAll(input, "threeight", "38")
	input = strings.ReplaceAll(input, "fiveight", "58")
	input = strings.ReplaceAll(input, "eighthree", "83")
	input = strings.ReplaceAll(input, "nineight", "98")
	input = strings.ReplaceAll(input, "eightwo", "82")
	input = strings.ReplaceAll(input, "sevenine", "79")
	input = strings.ReplaceAll(input, "one", "1")
	input = strings.ReplaceAll(input, "two", "2")
	input = strings.ReplaceAll(input, "three", "3")
	input = strings.ReplaceAll(input, "four", "4")
	input = strings.ReplaceAll(input, "five", "5")
	input = strings.ReplaceAll(input, "six", "6")
	input = strings.ReplaceAll(input, "seven", "7")
	input = strings.ReplaceAll(input, "eight", "8")
	input = strings.ReplaceAll(input, "nine", "9")
	return input
}

func main() {
	input := utils.ReadFile("./input")

	var ansone int
	var anstwo int
	for _, line := range input {
		ansone += number(line)
		line = replaceNumber(line)
		anstwo += number(line)
	}
	fmt.Println(ansone)
	fmt.Println(anstwo)
}
