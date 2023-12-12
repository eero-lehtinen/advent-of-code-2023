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

type cacheItem struct {
	row    string
	groups [50]int
}

func makeCacheKey(row string, groups []int) cacheItem {
	g := [50]int{}
	copy(g[:], groups)
	return cacheItem{row, g}
}

var cache = map[cacheItem]int{}

func cachedCheck(row string, groups []int) int {
	cacheKey := makeCacheKey(row, groups)
	if val, ok := cache[cacheKey]; ok {
		return val
	}

	val := check(row, groups)
	cache[cacheKey] = val
	return val
}

func check(row string, groups []int) int {
	if len(groups) == 0 {
		for _, char := range row {
			if char == '#' {
				return 0
			}
		}
		return 1
	}

	if row == "" {
		return 0
	}

	if groups[0] > len(row) {
		return 0
	}

	if row[0] == '?' {
		return cachedCheck("#"+row[1:], groups) +
			cachedCheck("."+row[1:], groups)
	}

	if row[0] == '.' {
		return cachedCheck(row[1:], groups)
	}

	valid := true
	for i := 0; i < groups[0]; i++ {
		if row[i] == '.' {
			valid = false
			break
		}
	}

	if valid {
		if groups[0] == len(row) {
			if len(groups) == 1 {
				return 1
			}
			return 0
		}
		if row[groups[0]] == '#' {
			return 0
		}

		return cachedCheck(row[groups[0]+1:], groups[1:])
	}

	return 0
}

func part(partNum int) {
	fmt.Println("Part", partNum)

	lines := readFileLines("input.txt")

	rows := []string{}
	damagedGroups := [][]int{}
	for _, line := range lines {
		parts := strings.Split(line, " ")
		strGroups := strings.Split(parts[1], ",")
		group := []int{}
		for _, strGroup := range strGroups {
			num, _ := strconv.Atoi(strGroup)
			group = append(group, num)
		}

		if partNum == 2 {
			oldRow := parts[0]
			oldGroup := group
			for i := 0; i < 4; i++ {
				parts[0] = parts[0] + "?" + oldRow
				group = append(group, oldGroup...)
			}
		}

		rows = append(rows, parts[0])
		damagedGroups = append(damagedGroups, group)
	}

	totalArrangements := 0

	for i, row := range rows {
		arrangements := cachedCheck(row, damagedGroups[i])
		totalArrangements += arrangements
	}

	fmt.Println("Answer:", totalArrangements)
}

func main() {
	part(1)
	part(2)
}
