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

func part1() {
	fmt.Println("Part1")

	lines := readFileLines("input.txt")

	bagCubes := map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}

	possibleGames := []int{}

	for _, line := range lines {
		possibleGame := true
		parts := strings.Split(line, ": ")
		gameNum, _ := strconv.Atoi(strings.Split(parts[0], " ")[1])
		cubeGroups := strings.Split(parts[1], "; ")
	game:
		for _, cubeGroup := range cubeGroups {
			cubeStrings := strings.Split(cubeGroup, ", ")
			for _, cubeString := range cubeStrings {
				cubeStringParts := strings.Split(cubeString, " ")
				cubeCount, _ := strconv.Atoi(cubeStringParts[0])
				cubeColor := cubeStringParts[1]
				if bagCubes[cubeColor] < cubeCount {
					possibleGame = false
					break game
				}
			}
		}
		if possibleGame {
			possibleGames = append(possibleGames, gameNum)
		}
	}

	gameSum := 0
	for _, game := range possibleGames {
		gameSum += game
	}

	fmt.Println("Answer:", gameSum)
}

func part2() {
	fmt.Println("Part2")

	lines := readFileLines("input.txt")

	gamePowers := []int{}

	for _, line := range lines {
		fewestPossibleCubes := map[string]int{
			"red":   0,
			"green": 0,
			"blue":  0,
		}

		parts := strings.Split(line, ": ")
		cubeGroups := strings.Split(parts[1], "; ")
		for _, cubeGroup := range cubeGroups {
			cubeStrings := strings.Split(cubeGroup, ", ")
			for _, cubeString := range cubeStrings {
				cubeStringParts := strings.Split(cubeString, " ")
				cubeCount, _ := strconv.Atoi(cubeStringParts[0])
				cubeColor := cubeStringParts[1]
				if fewestPossibleCubes[cubeColor] < cubeCount {
					fewestPossibleCubes[cubeColor] = cubeCount
				}
			}
		}

		cubePower := 1
		for _, cubeCount := range fewestPossibleCubes {
			cubePower *= cubeCount
		}

		gamePowers = append(gamePowers, cubePower)
	}

	gameSum := 0
	for _, game := range gamePowers {
		gameSum += game
	}

	fmt.Println("Answer:", gameSum)
}

func main() {
	part1()
	part2()
}
