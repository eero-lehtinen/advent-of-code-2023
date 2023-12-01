//go:build ignore

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func parseDigit(r byte) (int, bool) {
	if r >= '0' && r <= '9' {
		return int(r - '0'), true
	}
	return 0, false
}

func findDigits1(s string) (int, int) {
	var ok bool
	var firstDigit int
	var lastDigit int

	for i := range s {
		firstDigit, ok = parseDigit(s[i])
		if ok {
			break
		}
	}

	for i := range s {
		lastDigit, ok = parseDigit(s[len(s)-1-i])
		if ok {
			break
		}
	}

	return firstDigit, lastDigit
}

var wordDigits = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

func findDigits2(s string) (int, int) {
	var firstDigit int
	var lastDigit int

out1:
	for i := range s {
		d, ok := parseDigit(s[i])
		if ok {
			firstDigit = d
			break out1
		}

		for word, digit := range wordDigits {
			if strings.HasPrefix(s[i:], word) {
				firstDigit = digit
				break out1
			}
		}
	}

out2:
	for i := range s {
		d, ok := parseDigit(s[len(s)-1-i])
		if ok {
			lastDigit = d
			break out2
		}

		for word, digit := range wordDigits {
			if strings.HasSuffix(s[:len(s)-i], word) {
				lastDigit = digit
				break out2
			}
		}
	}

	return firstDigit, lastDigit
}

func readFileLines(filename string) []string {
	file, err := os.Open(filename)
	if err != nil {
		panic("Error reading file")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	lines := []string{}
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}

func part1() {
	fmt.Println("Part1")

	lines := readFileLines("input.txt")
	score := 0

	for _, line := range lines {
		firstDigit, lastDigit := findDigits1(line)
		score += firstDigit*10 + lastDigit
	}

	fmt.Println("Answer:", score)
}

func part2() {
	fmt.Println("Part2")

	lines := readFileLines("input.txt")
	score := 0

	for _, line := range lines {
		firstDigit, lastDigit := findDigits2(line)
		score += firstDigit*10 + lastDigit
	}

	fmt.Println("Answer:", score)

}

func main() {
	part1()
	part2()
}
