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

func readFileAsLine(filename string) string {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	line := ""
	for scanner.Scan() {
		line += scanner.Text()
	}

	return line
}

func hash(input string) int {
	cur := 0
	for _, c := range input {
		ascii := int(c)
		cur += ascii
		cur *= 17
		cur = cur % 256
	}
	return cur
}

func part1() {
	fmt.Println("Part 1")

	line := readFileAsLine("input.txt")

	parts := strings.Split(line, ",")

	total := 0
	for _, part := range parts {
		total += hash(part)
	}

	fmt.Println("Answer:", total)
}

type item struct {
	label string
	value int
}

func find(arr []item, label string) int {
	for i, item := range arr {
		if item.label == label {
			return i
		}
	}
	return -1
}

func part2() {
	fmt.Println("Part 2")

	line := readFileAsLine("input.txt")

	parts := strings.Split(line, ",")
	partRegex := regexp.MustCompile(`(.*)(=|-)(\d+)?`)

	boxes := make([][]item, 256)

	for _, part := range parts {
		matches := partRegex.FindStringSubmatch(part)
		label := matches[1]
		op := matches[2]
		value := -1
		if op == "=" {
			value, _ = strconv.Atoi(matches[3])
		}

		hash := hash(label)
		switch op {
		case "=":
			index := find(boxes[hash], label)
			if index == -1 {
				boxes[hash] = append(boxes[hash], item{label, value})
			} else {
				boxes[hash][index].value = value
			}
		case "-":
			index := find(boxes[hash], label)
			if index != -1 {
				boxes[hash] = append(boxes[hash][:index], boxes[hash][index+1:]...)
			}
		default:
			panic("Unknown op")
		}
	}

	power := 0
	for i, box := range boxes {
		for j, item := range box {
			power += (i + 1) * (j + 1) * item.value
		}
	}

	fmt.Println("Answer:", power)
}

func main() {
	part1()
	part2()
}
