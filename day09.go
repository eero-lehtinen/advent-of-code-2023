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

func main() {
	lines := readFileLines("input.txt")

	histories := [][]int{}
	for _, line := range lines {
		history := []int{}
		for _, part := range strings.Split(line, " ") {
			num, _ := strconv.Atoi(part)
			history = append(history, num)
		}
		histories = append(histories, history)
	}

	answer1 := 0
	answer2 := 0

	for _, history := range histories {
		sequences := [][]int{}
		sequences = append(sequences, []int{})
		sequences[0] = append(sequences[0], history...)

		sequencesDone := func() bool {
			for _, val := range sequences[len(sequences)-1] {
				if val != 0 {
					return false
				}
			}
			return true
		}

		for !sequencesDone() {
			i := len(sequences) - 1
			prevSequence := sequences[i]
			sequence := make([]int, len(prevSequence)-1)
			for j := 0; j < len(sequence); j++ {
				sequence[j] = prevSequence[j+1] - prevSequence[j]
			}
			sequences = append(sequences, sequence)
		}

		forward := 0
		backward := 0
		for i := len(sequences) - 2; i >= 0; i-- {
			forward += sequences[i][len(sequences[i])-1]
			backward = sequences[i][0] - backward
		}
		answer1 += forward
		answer2 += backward
	}

	fmt.Println("Part1\nAnswer:", answer1)
	fmt.Println("Part2\nAnswer:", answer2)
}
