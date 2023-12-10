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

func startPos(lines []string) pos {
	for y, line := range lines {
		for x, char := range line {
			if char == 'S' {
				return pos{x, y}
			}
		}
	}
	panic("no start position")
}

func sliceContains[T comparable](slice []T, val T) bool {
	for _, v := range slice {
		if v == val {
			return true
		}
	}
	return false
}

type stackItem struct {
	pos    pos
	parent pos
}

func findLoop(lines []string) []pos {
	N := pos{0, -1}
	S := pos{0, 1}
	E := pos{1, 0}
	W := pos{-1, 0}

	tileDirs := map[byte][]pos{
		'|': {N, S},
		'-': {E, W},
		'L': {N, E},
		'J': {N, W},
		'7': {S, W},
		'F': {S, E},
		'.': {},
	}

	inBounds := func(pos pos) bool {
		return pos.x >= 0 && pos.x < len(lines[0]) && pos.y >= 0 && pos.y < len(lines)
	}

	startPos := startPos(lines)
	startDirs := []pos{}
	candidates := [][]pos{{N, S}, {S, N}, {E, W}, {W, E}}
	for _, dirs := range candidates {
		pos := startPos.add(dirs[0])
		if !inBounds(pos) {
			continue
		}
		if sliceContains(tileDirs[lines[pos.y][pos.x]], dirs[1]) {
			startDirs = append(startDirs, dirs[0])
		}
	}
	tileDirs['S'] = startDirs

	stack := []stackItem{{startPos, pos{-1, -1}}}
	visited := map[pos]struct{}{startPos: {}}
	parents := map[pos]pos{startPos: {-1, -1}}

	for len(stack) > 0 {
		item := stack[len(stack)-1]
		cur := item.pos

		stack = stack[:len(stack)-1]

		visited[cur] = struct{}{}
		parents[cur] = item.parent

		tile := lines[cur.y][cur.x]

		dirs := tileDirs[tile]
		for _, dir := range dirs {
			neighbor := cur.add(dir)
			if !inBounds(neighbor) {
				continue
			}
			if _, ok := visited[neighbor]; !ok {
				stack = append(stack, stackItem{neighbor, cur})
			} else if parent, ok := parents[cur]; ok && parent != neighbor {
				loop := []pos{cur}
				for cur != neighbor {
					cur = parents[cur]
					loop = append(loop, cur)
				}
				return loop
			}
		}
	}
	panic("no loop")
}

func part1() {
	fmt.Println("Part 1")

	lines := readFileLines("input.txt")
	loop := findLoop(lines)
	fmt.Println("Answer:", len(loop)/2)
}

func pointInPolygon(p pos, polygon []pos) bool {
	inside := false
	for i := 0; i < len(polygon); i++ {
		j := (i + 1) % len(polygon)
		p1 := polygon[i]
		p2 := polygon[j]
		if (p1.y > p.y) != (p2.y > p.y) &&
			p.x < (p2.x-p1.x)*(p.y-p1.y)/(p2.y-p1.y)+p1.x {
			inside = !inside
		}
	}
	return inside
}

func part2() {
	fmt.Println("Part 2")

	lines := readFileLines("input.txt")
	loop := findLoop(lines)

	width := len(lines[0])
	height := len(lines)

	pointsInside := 0
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if sliceContains(loop, pos{x, y}) {
				continue
			}
			if pointInPolygon(pos{x, y}, loop) {
				pointsInside++
			}
		}
	}

	fmt.Println("Answer:", pointsInside)
}

func main() {
	part1()
	part2()
}
