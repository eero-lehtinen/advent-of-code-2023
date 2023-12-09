//go:build ignore

package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
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

type node struct {
	left  string
	right string
}

func readNodes(lines []string) map[string]node {
	nodes := map[string]node{}
	nodeRegexp := regexp.MustCompile(`(\w+) = \((\w+), (\w+)\)`)
	for _, line := range lines {
		matches := nodeRegexp.FindStringSubmatch(line)
		nodes[matches[1]] = node{matches[2], matches[3]}
	}
	return nodes
}

func part1() {
	fmt.Println("Part1")

	lines := readFileLines("input.txt")

	instructions := lines[0]

	nodes := readNodes(lines[2:])

	count := 0
	i := 0
	n := "AAA"
	for n != "ZZZ" {
		inst := instructions[i]
		if inst == 'L' {
			n = nodes[n].left
		} else {
			n = nodes[n].right
		}
		count++
		i++
		if i >= len(instructions) {
			i = 0
		}
	}

	fmt.Println("Answer:", count)
}

func GCD(a int, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func LCM(ns []int) int {
	lcm := 1
	for _, n := range ns {
		lcm = lcm * n / GCD(lcm, n)
	}
	return lcm
}

func part2() {
	fmt.Println("Part2")

	lines := readFileLines("input.txt")

	instructions := lines[0]

	nodes := readNodes(lines[2:])

	ns := []string{}
	for key := range nodes {
		if key[2] == 'A' {
			ns = append(ns, key)
		}
	}

	done := make([]int, len(ns))
	allDone := func() bool {
		for _, d := range done {
			if d == 0 {
				return false
			}
		}
		return true
	}

	count := 0
	i := 0
	for !allDone() {
		inst := instructions[i]
		left := inst == 'L'
		count++
		for j := range ns {
			if done[j] > 0 {
				continue
			}
			if left {
				ns[j] = nodes[ns[j]].left
			} else {
				ns[j] = nodes[ns[j]].right
			}
			if ns[j][2] == 'Z' {
				done[j] = count
			}
		}
		i++
		if i >= len(instructions) {
			i = 0
		}
	}

	lcm := LCM(done)

	fmt.Println("Answer:", lcm)
}
func main() {
	part1()
	part2()
}
