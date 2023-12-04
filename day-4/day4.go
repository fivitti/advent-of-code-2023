package day4

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func power(base, exponent int) int {
	result := 1
	for i := 0; i < exponent; i++ {
		result *= base
	}
	return result
}

type card struct {
	id             int
	winningNumbers map[int]any
	havingNumbers  []int
}

func (c *card) getHavingWinningNumbers() []int {
	var numbers []int
	for _, number := range c.havingNumbers {
		if _, ok := c.winningNumbers[number]; ok {
			numbers = append(numbers, number)
		}
	}
	return numbers
}

func (c *card) countHavingWinningNumbers() int {
	return len(c.getHavingWinningNumbers())
}

func (c *card) getPoints() int {
	count := c.countHavingWinningNumbers()
	if count == 0 {
		return 0
	}
	return power(2, count-1)
}

func newCard(id int, winningNumbers, havingNumbers []int) card {
	c := card{
		id:             id,
		winningNumbers: make(map[int]any),
		havingNumbers:  havingNumbers,
	}
	for _, number := range winningNumbers {
		c.winningNumbers[number] = nil
	}
	return c
}

// ID  Winning numbers  Having numbers
// Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53
func parseCard(line string) card {
	cardPart, numberPart, _ := strings.Cut(line, ": ")
	idRaw := strings.TrimPrefix(cardPart, "Card ")
	id, err := strconv.Atoi(strings.TrimSpace(idRaw))
	if err != nil {
		panic(err)
	}

	winningPart, havingPart, _ := strings.Cut(numberPart, " | ")
	winningNumbersRaw := strings.Split(winningPart, " ")
	winningNumbers := make([]int, 0, len(winningNumbersRaw))
	for _, numberRaw := range winningNumbersRaw {
		if numberRaw == "" {
			continue
		}
		number, _ := strconv.Atoi(numberRaw)
		winningNumbers = append(winningNumbers, number)
	}

	havingNumbersRaw := strings.Split(havingPart, " ")
	havingNumbers := make([]int, 0, len(havingNumbersRaw))
	for _, numberRaw := range havingNumbersRaw {
		if numberRaw == "" {
			continue
		}
		number, _ := strconv.Atoi(numberRaw)
		havingNumbers = append(havingNumbers, number)
	}

	return newCard(id, winningNumbers, havingNumbers)
}

func parseCards(reader io.Reader) []card {
	cards := make([]card, 0)
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		line := scanner.Text()
		cards = append(cards, parseCard(line))
	}
	return cards
}

func countPoints(cards []card) int {
	points := 0
	for _, card := range cards {
		points += card.getPoints()
	}
	return points
}

func solveStack(cards []card) []int {
	winningHavingCounts := make(map[int]int)
	for _, card := range cards {
		winningHavingCounts[card.id] = card.countHavingWinningNumbers()
	}

	stack := make([]int, len(cards))
	for i := range stack {
		stack[i] = 1
	}

	for i := range cards {
		n := cards[i].countHavingWinningNumbers()

		for j := 0; j < n; j++ {
			stack[i+j+1] += stack[i]
		}
	}

	return stack
}

func sumScratchcards(counts []int) int {
	sum := 0
	for _, count := range counts {
		sum += count
	}
	return sum
}

func SolutionStage1() {
	input, err := os.Open("data.txt")
	if err != nil {
		panic(err)
	}
	defer input.Close()

	cards := parseCards(input)
	points := countPoints(cards)

	fmt.Println("Stage 1:", points)
}

func SolutionStage2() {
	input, err := os.Open("data.txt")
	if err != nil {
		panic(err)
	}
	defer input.Close()

	cards := parseCards(input)
	stack := solveStack(cards)
	count := sumScratchcards(stack)

	fmt.Println("Stage 2:", count)
}
