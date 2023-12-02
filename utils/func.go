package utils

import (
	"bufio"
	"fmt"
	"os"
)

func CkeckErr(msg interface{}) {
	if msg != nil {
		fmt.Fprintln(os.Stderr, "Error:", msg)
		os.Exit(1)
	}
}

func ReadFile(filename string) []string {
	f, err := os.Open(filename)
	CkeckErr(err)
	defer f.Close()

	var result []string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		result = append(result, line)
	}
	err = scanner.Err()
	CkeckErr(err)

	return result
}
