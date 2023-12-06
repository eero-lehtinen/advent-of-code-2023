//go:build ignore

package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
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

func isDigit(r byte) bool {
	return r >= '0' && r <= '9'
}

func toInts(strs []string) []int {
	ints := []int{}
	for _, str := range strs {
		n, _ := strconv.Atoi(str)
		ints = append(ints, n)
	}
	return ints
}

func calcDistance(time int, heldTime int) int {
	time -= heldTime
	return time * heldTime
}

func part1() {
	fmt.Println("Part1")

	lines := readFileLines("input.txt")

	numberRegexp := regexp.MustCompile(`\d+`)
	timeMatches := numberRegexp.FindAllString(lines[0], -1)
	distanceRecordMatches := numberRegexp.FindAllString(lines[1], -1)

	times := toInts(timeMatches)
	distanceRecords := toInts(distanceRecordMatches)

	if len(times) != len(distanceRecords) {
		panic("Times and distances are not equal")
	}

	answer := 1

	for i, time := range times {
		beatCount := 0
		prevRecord := distanceRecords[i]

		for held := 0; held <= time; held++ {
			distance := calcDistance(time, held)

			if distance <= prevRecord {
				continue
			}
			beatCount++

		}
		answer *= beatCount
	}

	fmt.Println("Answer:", answer)
}

func part2() {
	fmt.Println("Part2")

	lines := readFileLines("input.txt")

	numberRegexp := regexp.MustCompile(`\d+`)
	timeMatches := numberRegexp.FindAllString(lines[0], -1)
	distanceRecordMatches := numberRegexp.FindAllString(lines[1], -1)

	time, _ := strconv.Atoi(strings.Join(timeMatches, ""))
	prevRecord, _ := strconv.Atoi(strings.Join(distanceRecordMatches, ""))

	beatCount := 0

	for held := 0; held <= time; held++ {
		distance := calcDistance(time, held)
		if distance <= prevRecord {
			continue
		}
		beatCount++
	}

	fmt.Println("Answer:", beatCount)
}

func main() {
	part1()
	part2()
}
