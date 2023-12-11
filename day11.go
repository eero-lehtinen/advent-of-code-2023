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

type pos struct {
	x int
	y int
}

func (p pos) add(p2 pos) pos {
	return pos{p.x + p2.x, p.y + p2.y}
}

func printLines(lines []string) {
	for _, line := range lines {
		fmt.Println(line)
	}
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func manhattan(p1 pos, p2 pos) int {
	return abs(p1.x-p2.x) + abs(p1.y-p2.y)
}

func part(partNum int) {
	fmt.Println("Part", partNum)

	lines := readFileLines("input.txt")

	emptyRows := []int{}
	for row, line := range lines {
		empty := true
		for _, r := range line {
			if r != '.' {
				empty = false
				break
			}
		}
		if empty {
			emptyRows = append(emptyRows, row)
		}
	}

	emptyCols := []int{}
	for x := 0; x < len(lines[0]); x++ {
		empty := true
		for y := 0; y < len(lines); y++ {
			if lines[y][x] != '.' {
				empty = false
				break
			}
		}
		if empty {
			emptyCols = append(emptyCols, x)
		}
	}

	galaxies := []pos{}
	for y := 0; y < len(lines); y++ {
		for x := 0; x < len(lines[0]); x++ {
			if lines[y][x] == '#' {
				galaxies = append(galaxies, pos{x, y})
			}
		}
	}

	total := 0

	expansion := 0
	if partNum == 1 {
		expansion = 2
	} else {
		expansion = 1_000_000
	}

	for i := 0; i < len(galaxies); i++ {
		for j := i + 1; j < len(galaxies); j++ {
			p1 := galaxies[i]
			p2 := galaxies[j]

			extraCols := 0
			for _, col := range emptyCols {
				if col > p1.x && col < p2.x || col > p2.x && col < p1.x {
					extraCols += expansion - 1
				}
			}

			extraRows := 0
			for _, row := range emptyRows {
				if row > p1.y && row < p2.y || row > p2.y && row < p1.y {
					extraRows += expansion - 1
				}
			}

			dist := manhattan(p1, p2) + extraCols + extraRows
			total += dist
		}
	}

	fmt.Println("Answer:", total)
}

func main() {
	part(1)
	part(2)
}
