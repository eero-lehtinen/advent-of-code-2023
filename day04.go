//go:build ignore

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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

func toIntsSet(strs []string) map[int]struct{} {
	ints := map[int]struct{}{}
	for _, str := range strs {
		str = strings.TrimSpace(str)
		if str == "" {
			continue
		}
		n, _ := strconv.Atoi(str)
		ints[n] = struct{}{}
	}
	return ints
}

func toInts(strs []string) []int {
	ints := []int{}
	for _, str := range strs {
		str = strings.TrimSpace(str)
		if str == "" {
			continue
		}
		n, _ := strconv.Atoi(str)
		ints = append(ints, n)
	}
	return ints
}

func part1() {
	fmt.Println("Part1")

	lines := readFileLines("input.txt")

	total := 0

	for _, line := range lines {
		numberGroups := strings.Split(strings.Split(line, ": ")[1], " | ")
		winningNumbers := toIntsSet(strings.Split(numberGroups[0], " "))
		myNumbers := toInts(strings.Split(numberGroups[1], " "))

		points := 0
		for _, myNumber := range myNumbers {
			if _, ok := winningNumbers[myNumber]; ok {
				if points > 0 {
					points *= 2
				} else {
					points = 1
				}
			}
		}
		total += points
	}

	fmt.Println("Answer:", total)
}

func part2() {
	fmt.Println("Part2")

	lines := readFileLines("input.txt")

	cardMatches := []int{}

	for _, line := range lines {
		numberGroups := strings.Split(strings.Split(line, ": ")[1], " | ")
		winningNumbers := toIntsSet(strings.Split(numberGroups[0], " "))
		myNumbers := toInts(strings.Split(numberGroups[1], " "))

		matches := 0
		for _, myNumber := range myNumbers {
			if _, ok := winningNumbers[myNumber]; ok {
				matches += 1
			}
		}

		cardMatches = append(cardMatches, matches)
	}

	cardCounts := make([]int, len(lines))
	for i := range cardCounts {
		cardCounts[i] = 1
	}

	for i, count := range cardCounts {
		for j := i + 1; j <= i+cardMatches[i] && j < len(cardCounts); j++ {
			cardCounts[j] += count
		}
	}

	total := 0
	for _, count := range cardCounts {
		total += count
	}

	fmt.Println("Answer:", total)
}

func main() {
	part1()
	part2()
}
