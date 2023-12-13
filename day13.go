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

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func printLines(lines []string) {
	for _, line := range lines {
		fmt.Println(line)
	}
}

func part1() {
	fmt.Println("Part 1")

	lines := readFileLines("input.txt")

	patterns := [][]string{{}}
	for _, line := range lines {
		if line == "" {
			patterns = append(patterns, []string{})
			continue
		}
		patterns[len(patterns)-1] = append(patterns[len(patterns)-1], line)
	}

	total := 0

	for _, pattern := range patterns {
		// printLines(pattern)
		for col := 1; col < len(pattern[0]); col++ {
			maxOffset := min(col, len(pattern[0])-col)
			symmetric := true

			// fmt.Println("col", col, "maxOffset", maxOffset, "len(pattern[0])", len(pattern[0]))

		colout:
			for offset := 1; offset <= maxOffset; offset++ {
				// fmt.Println("Checking col", col-offset, "vs", col+offset-1)
				for _, row := range pattern {
					// fmt.Print(string(row[col-offset]), string(row[col+offset-1]), "\n")
					if row[col-offset] != row[col+offset-1] {
						symmetric = false
						break colout
					}
				}
			}

			if symmetric {
				// fmt.Println("Symmetric pattern at col", col, "in pattern", i)
				total += col
				break
			}
		}

		for row := 1; row < len(pattern); row++ {
			maxOffset := min(row, len(pattern)-row)
			symmetric := true

			// fmt.Println("row", row, "maxOffset", maxOffset, "len(pattern)", len(pattern))
		rowout:
			for offset := 1; offset <= maxOffset; offset++ {
				// fmt.Println("Checking row", row-offset, "vs row", row+offset-1, pattern[row-offset], pattern[row+offset-1])
				if pattern[row-offset] != pattern[row+offset-1] {
					symmetric = false
					break rowout
				}
			}

			if symmetric {
				// fmt.Println("Symmetric pattern at row", row, "in pattern", i)
				total += 100 * row
				break
			}
		}
	}

	fmt.Println("Answer:", total)
}

func part2() {
	fmt.Println("Part 2")

	lines := readFileLines("input.txt")

	patterns := [][]string{{}}
	for _, line := range lines {
		if line == "" {
			patterns = append(patterns, []string{})
			continue
		}
		patterns[len(patterns)-1] = append(patterns[len(patterns)-1], line)
	}

	total := 0

	for _, pattern := range patterns {
		// printLines(pattern)
		for col := 1; col < len(pattern[0]); col++ {
			maxOffset := min(col, len(pattern[0])-col)
			symmetric := true
			distance := 0

			// fmt.Println("col", col, "maxOffset", maxOffset, "len(pattern[0])", len(pattern[0]))
		colout:
			for offset := 1; offset <= maxOffset; offset++ {
				// fmt.Println("Checking col", col-offset, "vs", col+offset-1)
				for _, row := range pattern {
					// fmt.Print(string(row[col-offset]), string(row[col+offset-1]), "\n")
					if row[col-offset] != row[col+offset-1] {
						symmetric = false
						distance += 1
						if distance > 1 {
							break colout
						}
					}
				}
			}

			if !symmetric && distance == 1 {
				// fmt.Println("Symmetric with smudge at col", col, "in pattern", i)
				total += col
				break
			}
		}

		for row := 1; row < len(pattern); row++ {
			maxOffset := min(row, len(pattern)-row)
			symmetric := true
			distance := 0

			// fmt.Println("row", row, "maxOffset", maxOffset, "len(pattern)", len(pattern))
		rowout:
			for offset := 1; offset <= maxOffset; offset++ {
				// fmt.Println("Checking row", row-offset, "vs row", row+offset-1, pattern[row-offset], pattern[row+offset-1])
				for col := 0; col < len(pattern[0]); col++ {
					if pattern[row-offset][col] != pattern[row+offset-1][col] {
						symmetric = false
						distance += 1
						if distance > 1 {
							break rowout
						}
					}
				}
			}

			if !symmetric && distance == 1 {
				// fmt.Println("Symmetric with smudge at row", row, "in pattern", i)
				total += 100 * row
				break
			}
		}
	}

	fmt.Println("Answer:", total)
}

func main() {
	part1()
	part2()
}
