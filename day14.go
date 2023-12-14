//go:build ignore

package main

import (
	"bufio"
	"bytes"
	"crypto/sha256"
	"fmt"
	"os"
)

func readFileLines(filename string) [][]byte {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	lines := [][]byte{}
	for scanner.Scan() {
		lines = append(lines, []byte(scanner.Text()))
	}

	return lines
}

func printLines(lines [][]byte) {
	for _, line := range lines {
		fmt.Println(string(line))
	}
}

type coord struct {
	x int
	y int
}

func part1() {
	fmt.Println("Part 1")

	lines := readFileLines("input.txt")

	freeRocks := []coord{}
	for y, line := range lines {
		for x, char := range line {
			if char == 'O' {
				freeRocks = append(freeRocks, coord{x, y})
			}
		}
	}

	for len(freeRocks) > 0 {
		canMove := func(c coord) bool {
			if c.y < 0 {
				return false
			}
			if lines[c.y][c.x] != '.' {
				return false
			}
			return true
		}
		doneRocks := []int{}
		for i, rock := range freeRocks {
			if canMove(coord{rock.x, rock.y - 1}) {
				lines[rock.y][rock.x] = '.'
				lines[rock.y-1][rock.x] = 'O'
				freeRocks[i].y--
			} else {
				doneRocks = append(doneRocks, i)
			}
		}
		for i := len(doneRocks) - 1; i >= 0; i-- {
			freeRocks = append(freeRocks[:doneRocks[i]], freeRocks[doneRocks[i]+1:]...)
		}
	}

	// printLines(lines)

	score := 0
	for y, line := range lines {
		s := len(line) - y
		for _, char := range line {
			if char == 'O' {
				score += s
			}
		}
	}

	fmt.Println("Answer:", score)
}

type direction int

const (
	north direction = iota
	west
	south
	east
)

func sha256lines(lines [][]byte) [32]byte {
	return sha256.Sum256(bytes.Join(lines, []byte{'\n'}))
}

func part2() {
	fmt.Println("Part 2")

	lines := readFileLines("input.txt")

	rocks := []coord{}
	for y, line := range lines {
		for x, char := range line {
			if char == 'O' {
				rocks = append(rocks, coord{x, y})
			}
		}
	}

	hashes := [][32]byte{}

	targetIters := 1000000000

	jumped := false
	i := 0
	for i < targetIters {
		for dir := north; dir <= east; dir++ {
			var offset func(coord) coord
			switch dir {
			case north:
				offset = func(c coord) coord {
					return coord{c.x, c.y - 1}
				}
			case west:
				offset = func(c coord) coord {
					return coord{c.x - 1, c.y}
				}
			case south:
				offset = func(c coord) coord {
					return coord{c.x, c.y + 1}
				}
			case east:
				offset = func(c coord) coord {
					return coord{c.x + 1, c.y}
				}
			}

			moved := true

			for moved {
				moved = false
				canMove := func(c coord) bool {
					if c.y < 0 || c.y >= len(lines) || c.x < 0 || c.x >= len(lines[c.y]) {
						return false
					}
					if lines[c.y][c.x] != '.' {
						return false
					}
					return true
				}
				for i, rock := range rocks {
					newCoord := offset(rock)
					if canMove(newCoord) {
						lines[rock.y][rock.x] = '.'
						lines[newCoord.y][newCoord.x] = 'O'
						rocks[i] = newCoord
						moved = true
					}
				}
			}
		}

		score := 0
		for y, line := range lines {
			s := len(line) - y
			for _, char := range line {
				if char == 'O' {
					score += s
				}
			}
		}
		// fmt.Println(i, score)

		hash := sha256lines(lines)
		// fmt.Printf("Hash %x\n", hash[:8])
		hashes = append(hashes, hash)

		i++

		if !jumped {
			for j := 0; j < i-1; j++ {
				if hashes[j] == hash {
					// fmt.Println("Found loop", j, i)
					jumpSize := i - j - 1
					i += (targetIters - i) / jumpSize * jumpSize
					// fmt.Println("Jump size", jumpSize)
					// fmt.Println("New i", i)
					jumped = true
					break
				}
			}
		}

	}

	score := 0
	for y, line := range lines {
		s := len(line) - y
		for _, char := range line {
			if char == 'O' {
				score += s
			}
		}
	}

	fmt.Println("Answer:", score)
}
func main() {
	part1()
	part2()
}
