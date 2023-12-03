package main

import (
	"aoc/utils"
	"fmt"
	"strings"
)

var maxCube = map[string]int{
	"red":   12,
	"green": 13,
	"blue":  14,
}

func processLine(gameNum int, input string) (int, int) {
	game := strings.Split(input, ": ")
	rolls := strings.Split(game[1], "; ")

	var currCube = make(map[string]int)

	for _, roll := range rolls {
		cubes := strings.Split(roll, ", ")
		for _, cube := range cubes {
			cubeSplit := strings.Split(cube, " ")
			cubeNum := utils.Strint(cubeSplit[0])
			cubeColor := cubeSplit[1]

			lastMax, _ := currCube[cubeColor]
			currCube[cubeColor] = max(lastMax, cubeNum)
		}
	}

	for k := range currCube {
		if currCube[k] > maxCube[k] {
			gameNum = 0
			break
		}
	}

	power := 1
	for _, v := range currCube {
		power *= v
	}
	return gameNum, power
}

func main() {
	input := utils.ReadFile("./input")
	var ansone int
	var anstwo int

	for gameNum, line := range input {
		one, two := processLine((gameNum + 1), line)
		ansone += one
		anstwo += two
	}
	fmt.Println(ansone)
	fmt.Println(anstwo)
}
