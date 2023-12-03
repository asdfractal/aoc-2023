package main

import (
	"aoc/utils"
	"fmt"
	"testing"
)

func TestOne(t *testing.T) {
	var ansone int
	var anstwo int
	input := utils.ReadFile("./test_input01")
	for gameNum, line := range input {
		one, two := processLine((gameNum + 1), line)
		ansone += one
		anstwo += two
	}
	fmt.Println(ansone)
	if ansone != 8 {
		t.Errorf("expected 8, got %d", ansone)
	}
	fmt.Println(anstwo)
	if anstwo != 2286 {
		t.Errorf("expected 2286, got %d", anstwo)
	}
}
