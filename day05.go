//go:build ignore

package main

import (
	"bufio"
	"fmt"
	"math"
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

type seed struct {
	src  int
	size int
}

type rule struct {
	src  int
	dest int
	size int
}

func parseRule(line string) rule {
	parts := strings.Split(line, " ")
	from, _ := strconv.Atoi(parts[1])
	to, _ := strconv.Atoi(parts[0])
	size, _ := strconv.Atoi(parts[2])

	return rule{
		src:  from,
		dest: to,
		size: size,
	}
}

func (r rule) Convert(n int) (int, bool) {
	if n < r.src || n > r.src+r.size {
		return 0, false
	}
	pos := n - r.src
	return r.dest + pos, true
}

func (r rule) Convert2(seeds []seed) ([]seed, []seed) {
	unprocessed := []seed{}
	result := []seed{}
	for _, s := range seeds {
		ruleLeft := r.src
		ruleRight := r.src + r.size
		seedLeft := s.src
		seedRight := s.src + s.size
		if seedRight < ruleLeft || seedLeft > ruleRight {
			// Seed is outside of rule
			unprocessed = append(unprocessed, s)
		} else if seedLeft >= ruleLeft && seedRight <= ruleRight {
			// Seed is inside of rule
			newSrc, _ := r.Convert(s.src)
			result = append(result, seed{src: newSrc, size: s.size})
		} else if seedLeft <= ruleLeft {
			// Seed is to the left of rule, needs to be split
			seedSize := ruleLeft - seedLeft
			unprocessed = append(unprocessed, seed{src: seedLeft, size: seedSize})

			seedLeft2 := ruleLeft
			seedSize2 := s.size - seedSize
			newSrc, _ := r.Convert(seedLeft2)
			result = append(result, seed{src: newSrc, size: seedSize2})
		} else if seedRight >= ruleRight {
			// Seed is to the right of rule, needs to be split
			seedSize := seedRight - ruleRight - 1
			unprocessed = append(unprocessed, seed{src: ruleRight, size: seedSize})

			seedLeft2 := seedLeft
			seedSize2 := s.size - seedSize
			newSrc, _ := r.Convert(seedLeft2)
			result = append(result, seed{src: newSrc, size: seedSize2})
		} else {
			panic("unreachable")
		}
	}

	return result, unprocessed
}

type mapping []rule

func parseMappings(lines []string) []mapping {
	mappings := []mapping{}

	i := 0
	for _, line := range lines[2:] {
		if line == "" {
			i += 1
			continue
		}
		if !isDigit(line[0]) {
			continue
		}
		if i >= len(mappings) {
			mappings = append(mappings, mapping{})
		}
		mappings[i] = append(mappings[i], parseRule(line))
	}

	return mappings
}

func part1() {
	fmt.Println("Part1")

	lines := readFileLines("input.txt")

	seedRegexp := regexp.MustCompile(`\d+`)
	seeds := []int{}
	matches := seedRegexp.FindAllStringSubmatch(lines[0], -1)
	for _, match := range matches {
		n, _ := strconv.Atoi(match[0])
		seeds = append(seeds, n)
	}

	mappings := parseMappings(lines[2:])

	for i := range seeds {
		for _, mapping := range mappings {
			for _, rule := range mapping {
				newNum, ok := rule.Convert(seeds[i])
				if ok {
					seeds[i] = newNum
					break
				}
			}
		}
	}

	lowest := math.MaxInt
	for _, num := range seeds {
		if num < lowest {
			lowest = num
		}
	}

	fmt.Println("Answer:", lowest)
}

func part2() {
	fmt.Println("Part2")

	lines := readFileLines("input.txt")

	seedRegexp := regexp.MustCompile(`(\d+) (\d+)`)
	seeds := [][]seed{}
	matches := seedRegexp.FindAllStringSubmatch(lines[0], -1)
	for _, match := range matches {
		start, _ := strconv.Atoi(match[1])
		size, _ := strconv.Atoi(match[2])
		seeds = append(seeds, []seed{{src: start, size: size}})
	}

	mappings := parseMappings(lines[2:])

	for i := range seeds {
		for _, mapping := range mappings {
			result := []seed{}
			for _, rule := range mapping {
				r, u := rule.Convert2(seeds[i])
				seeds[i] = u
				result = append(result, r...)
			}

			seeds[i] = append(seeds[i], result...)
		}
	}

	lowest := math.MaxInt
	for _, arr := range seeds {
		for _, seed := range arr {
			if seed.src < lowest {
				lowest = seed.src
			}
		}
	}

	fmt.Println("Answer:", lowest)
}

func main() {
	part1()
	part2()
}
