package main

import (
	"aoc/utils"
	"fmt"
	"testing"
)

func TestOne(t *testing.T) {
	input := utils.ReadFile("./test_input01")

	var ans int
	for _, line := range input {
		ans += number(line)
	}

	if ans != 142 {
		t.Errorf("expected 142, got %d", ans)
	}
}

func TestTwo(t *testing.T) {
	input := utils.ReadFile("./test_input02")

	var ans int
	for _, line := range input {
		line = replaceNumber(line)
		fmt.Println(line)
		ans += number(line)
	}

	if ans != 281 {
		t.Errorf("expected 281, got %d", ans)
	}
}
