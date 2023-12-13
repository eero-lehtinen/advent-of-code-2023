//go:build ignore

package main

import (
	"bufio"
	"fmt"
	"os"
)

func readFileLines(filename string) []string {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	lines := []string{}
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}

func main() {
	fmt.Println("Part 1")

	lines := readFileLines("sample.txt")

	fmt.Println("Answer:", len(lines))
}
