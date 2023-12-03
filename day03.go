//go:build ignore

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func isDigit(r byte) bool {
	return r >= '0' && r <= '9'
}

func isSymbol(r byte) bool {
	return !isDigit(r) && r != '.'
}

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

type coord struct {
	x int
	y int
}

var directions = [...]coord{
	{-1, -1},
	{0, -1},
	{1, -1},
	{-1, 0},
	{1, 0},
	{-1, 1},
	{0, 1},
	{1, 1},
}

func searchNumber(lines []string, visited map[coord]bool, x int, y int) (int, bool) {
	visit := func(x int, y int) bool {
		if isDigit(lines[y][x]) && !visited[coord{x, y}] {
			visited[coord{x, y}] = true
			return true
		}
		return false
	}

	if y < 0 || y >= len(lines) || x < 0 || x >= len(lines[y]) {
		return 0, false
	}

	if !visit(x, y) {
		return 0, false
	}

	left := x
	for left-1 >= 0 && visit(left-1, y) {
		left -= 1
	}

	right := x
	for right+1 < len(lines[y]) && visit(right+1, y) {
		right += 1
	}

	number, err := strconv.Atoi(lines[y][left : right+1])
	if err != nil {
		panic(err)
	}

	return number, true
}

func part1() {
	fmt.Println("Part1")

	lines := readFileLines("input.txt")

	visited := map[coord]bool{}
	total := 0

	for y, line := range lines {
		for x, char := range line {
			if !isSymbol(byte(char)) {
				continue
			}
			for _, dir := range directions {
				number, ok := searchNumber(lines, visited, x+dir.x, y+dir.y)
				if ok {
					total += number
				}
			}
		}
	}

	fmt.Println("Answer:", total)
}

func part2() {
	fmt.Println("Part2")

	lines := readFileLines("input.txt")

	visited := map[coord]bool{}
	total := 0

	for y, line := range lines {
		for x, char := range line {
			if byte(char) != '*' {
				continue
			}
			numbers := []int{}
			for _, dir := range directions {
				number, ok := searchNumber(lines, visited, x+dir.x, y+dir.y)
				if ok {
					numbers = append(numbers, number)
				}
			}
			if len(numbers) == 2 {
				total += numbers[0] * numbers[1]
			}
		}
	}

	fmt.Println("Answer:", total)
}

func main() {
	part1()
	part2()
}
