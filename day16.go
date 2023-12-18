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

type direction int

const (
	UP direction = iota
	RIGHT
	DOWN
	LEFT
)

type beam struct {
	x, y int
	dir  direction
}

func checkEnergize(lines []string, in beam) int {
	energized := make([][][]bool, len(lines))
	for i := range energized {
		energized[i] = make([][]bool, len(lines[0]))
		for j := range energized[i] {
			energized[i][j] = make([]bool, 4)
		}
	}

	beams := []beam{in}

	for len(beams) > 0 {
		b := &beams[len(beams)-1]
		switch b.dir {
		case UP:
			b.y -= 1
		case RIGHT:
			b.x += 1
		case DOWN:
			b.y += 1
		case LEFT:
			b.x -= 1
		}

		if b.x < 0 || b.x >= len(lines[0]) || b.y < 0 || b.y >= len(lines) {
			beams = beams[:len(beams)-1]
			continue
		}

		if energized[b.y][b.x][b.dir] {
			beams = beams[:len(beams)-1]
			continue
		} else {
			energized[b.y][b.x][b.dir] = true
		}

		cur := lines[b.y][b.x]
		switch cur {
		case '/':
			switch b.dir {
			case UP:
				b.dir = RIGHT
			case RIGHT:
				b.dir = UP
			case DOWN:
				b.dir = LEFT
			case LEFT:
				b.dir = DOWN
			}
		case '\\':
			switch b.dir {
			case UP:
				b.dir = LEFT
			case RIGHT:
				b.dir = DOWN
			case DOWN:
				b.dir = RIGHT
			case LEFT:
				b.dir = UP
			}
		case '|':
			switch b.dir {
			case LEFT, RIGHT:
				b.dir = UP
				beams = append(beams, beam{x: b.x, y: b.y, dir: DOWN})
			}
		case '-':
			switch b.dir {
			case UP, DOWN:
				b.dir = RIGHT
				beams = append(beams, beam{x: b.x, y: b.y, dir: LEFT})
			}
		}
	}

	total := 0
	for _, row := range energized {
		for _, cell := range row {
			e := false
			for _, x := range cell {
				e = e || x
			}
			if e {
				total += 1
			}
		}
	}

	return total
}

func part1() {
	fmt.Println("Part 1")

	lines := readFileLines("input.txt")

	total := checkEnergize(lines, beam{-1, 0, RIGHT})

	fmt.Println("Answer:", total)
}

func part2() {
	fmt.Println("Part 2")

	lines := readFileLines("input.txt")

	bestTotal := 0

	for y := range lines {
		totalR := checkEnergize(lines, beam{-1, y, RIGHT})
		if totalR > bestTotal {
			bestTotal = totalR
		}
		totalL := checkEnergize(lines, beam{len(lines[0]), y, LEFT})
		if totalL > bestTotal {
			bestTotal = totalL
		}
	}

	for x := range lines[0] {
		totalD := checkEnergize(lines, beam{x, -1, DOWN})
		if totalD > bestTotal {
			bestTotal = totalD
		}
		totalU := checkEnergize(lines, beam{x, len(lines), UP})
		if totalU > bestTotal {
			bestTotal = totalU
		}
	}

	fmt.Println("Answer:", bestTotal)
}

func main() {
	part1()
	part2()
}
