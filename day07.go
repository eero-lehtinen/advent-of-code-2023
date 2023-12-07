//go:build ignore

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

func indexOf[T comparable](slice []T, fn func(T) bool) int {
	for i, v := range slice {
		if fn(v) {
			return i
		}
	}
	return -1
}

type group struct {
	card  byte
	count int
}

func handScore(hand [5]byte) int {
	grouped := []group{}
	for _, card := range hand {
		index := indexOf(grouped, func(g group) bool {
			return g.card == card
		})

		if index == -1 {
			grouped = append(grouped, group{card, 1})
		} else {
			grouped[index].count++
		}
	}

	if len(grouped) == 1 {
		// Five of a kind
		return 7
	} else if len(grouped) == 2 {
		if grouped[0].count == 4 || grouped[1].count == 4 {
			// Four of a kind
			return 6
		} else {
			// Full house
			return 5
		}
	} else if len(grouped) == 3 {
		index := indexOf(grouped, func(g group) bool {
			return g.count == 3
		})
		if index != -1 {
			// Three of a kind
			return 3
		} else {
			// Two pair
			return 2
		}
	} else if len(grouped) == 4 {
		// One pair
		return 1
	} else {
		// High card
		return 0
	}
}

type hand struct {
	cards [5]byte
	score int
	bid   int
}

func makeCardScore() map[byte]int {
	cardScore := map[byte]int{
		'A': 14,
		'K': 13,
		'Q': 12,
		'J': 11,
		'T': 10,
		'9': 9,
		'8': 8,
		'7': 7,
		'6': 6,
		'5': 5,
		'4': 4,
		'3': 3,
		'2': 2,
	}
	return cardScore
}

func makeHand(line string) hand {
	parts := strings.Split(line, " ")
	cards := [5]byte{}
	for i, byte := range []byte(parts[0]) {
		cards[i] = byte
	}
	bid, _ := strconv.Atoi(parts[1])
	return hand{cards, handScore(cards), bid}
}

func part1() {
	fmt.Println("Part1")

	lines := readFileLines("input.txt")

	cardScore := makeCardScore()

	hands := []hand{}

	for _, line := range lines {
		hands = append(hands, makeHand(line))
	}

	sort.Slice(hands, func(i, j int) bool {
		if hands[i].score == hands[j].score {
			for k := 0; k < 5; k++ {
				if hands[i].cards[k] == hands[j].cards[k] {
					continue
				}
				return cardScore[hands[i].cards[k]] < cardScore[hands[j].cards[k]]
			}
		}
		return hands[i].score < hands[j].score
	})

	total := 0
	for i, hand := range hands {
		total += hand.bid * (i + 1)
	}

	fmt.Println("Answer:", total)
}

func handScore2(hand [5]byte) int {
	grouped := []group{}
	for _, card := range hand {
		index := indexOf(grouped, func(g group) bool {
			return g.card == card
		})

		if index == -1 {
			grouped = append(grouped, group{card, 1})
		} else {
			grouped[index].count++
		}
	}

	jokerIndex := indexOf(grouped, func(g group) bool {
		return g.card == 'J'
	})

	if jokerIndex != -1 {
		jokerCount := grouped[jokerIndex].count
		if jokerCount == 5 {
			// Five of a kind
			return 7
		}

		largestGroup := &group{0, 0}
		for i := range grouped {
			if grouped[i].card != 'J' && grouped[i].count > largestGroup.count {
				largestGroup = &grouped[i]
			}
		}

		largestGroup.count += jokerCount

		grouped[jokerIndex] = grouped[len(grouped)-1]
		grouped = grouped[:len(grouped)-1]
	}

	if len(grouped) == 1 {
		// Five of a kind
		return 7
	} else if len(grouped) == 2 {
		if grouped[0].count == 4 || grouped[1].count == 4 {
			// Four of a kind
			return 6
		} else {
			// Full house
			return 5
		}
	} else if len(grouped) == 3 {
		index := indexOf(grouped, func(g group) bool {
			return g.count == 3
		})
		if index != -1 {
			// Three of a kind
			return 3
		} else {
			// Two pair
			return 2
		}
	} else if len(grouped) == 4 {
		// One pair
		return 1
	} else {
		// High card
		return 0
	}
}

func makeHand2(line string) hand {
	parts := strings.Split(line, " ")
	cards := [5]byte{}
	for i, byte := range []byte(parts[0]) {
		cards[i] = byte
	}
	bid, _ := strconv.Atoi(parts[1])
	return hand{cards, handScore2(cards), bid}
}

func makeCardScore2() map[byte]int {
	cardScore := map[byte]int{
		'A': 14,
		'K': 13,
		'Q': 12,
		'T': 10,
		'9': 9,
		'8': 8,
		'7': 7,
		'6': 6,
		'5': 5,
		'4': 4,
		'3': 3,
		'2': 2,
		'J': 1,
	}
	return cardScore
}

func part2() {
	fmt.Println("Part2")

	lines := readFileLines("input.txt")

	cardScore := makeCardScore2()

	hands := []hand{}

	for _, line := range lines {
		hands = append(hands, makeHand2(line))
	}

	sort.Slice(hands, func(i, j int) bool {
		if hands[i].score == hands[j].score {
			for k := 0; k < 5; k++ {
				if hands[i].cards[k] == hands[j].cards[k] {
					continue
				}
				return cardScore[hands[i].cards[k]] < cardScore[hands[j].cards[k]]
			}
		}
		return hands[i].score < hands[j].score
	})

	total := 0
	for i, hand := range hands {
		total += hand.bid * (i + 1)
	}

	fmt.Println("Answer:", total)
}

func main() {
	part1()
	part2()
}
