package day4

import (
	"advent-of-code-2023/utils/require"
	"testing"
)

func TestCardCountHavingWinningNumber(t *testing.T) {
	// Arrange
	card := newCard(42, []int{1, 2, 3}, []int{1, 2, 3, 4, 5})

	// Act
	numbers := card.getHavingWinningNumbers()

	// Assert
	require.Len(t, numbers, 3)
	require.Equal(t, 1, numbers[0])
	require.Equal(t, 2, numbers[1])
	require.Equal(t, 3, numbers[2])
}

func TestParseCard(t *testing.T) {
	// Arrange
	input := []string{
		"Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53",
		"Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19",
		"Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1",
		"Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83",
		"Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36",
		"Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11",
	}

	expected := [][]int{
		{83, 86, 17, 48},
		{32, 61},
		{1, 21},
		{84},
		{},
		{},
	}

	for i := 0; i < len(input); i++ {
		// Act
		card := parseCard(input[i])

		// Assert
		require.Equal(t, i+1, card.id)
		numbers := card.getHavingWinningNumbers()
		require.Len(t, numbers, len(expected[i]))
		for j := 0; j < len(numbers); j++ {
			require.Contains(t, expected[i], numbers[j])
		}
	}
}

func TestCountPoints(t *testing.T) {
	// Arrange
	input := []string{
		"Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53",
		"Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19",
		"Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1",
		"Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83",
		"Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36",
		"Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11",
	}

	cards := make([]card, 0, len(input))
	for _, line := range input {
		cards = append(cards, parseCard(line))
	}

	// Act
	points := countPoints(cards)

	// Assert
	require.Equal(t, 13, points)
}

func TestSolveStack(t *testing.T) {
	// Arrange
	input := []string{
		"Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53",
		"Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19",
		"Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1",
		"Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83",
		"Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36",
		"Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11",
	}

	cards := make([]card, 0, len(input))
	for _, line := range input {
		cards = append(cards, parseCard(line))
	}

	// Act
	stack := solveStack(cards)

	// Assert
	require.Equal(t, 1, stack[1])
	require.Equal(t, 2, stack[2])
	require.Equal(t, 4, stack[3])
	require.Equal(t, 8, stack[4])
	require.Equal(t, 14, stack[5])
	require.Equal(t, 1, stack[6])
}

func TestSumScratchcards(t *testing.T) {
	// Arrange
	scratchcards := []int{
		1,
		2,
		4,
		8,
		14,
		1,
	}

	// Act
	sum := sumScratchcards(scratchcards)

	// Assert
	require.Equal(t, 30, sum)
}

func ExampleSolutionStage1() {
	SolutionStage1()
	// Output:
	// Stage 1: 26426
}

func ExampleSolutionStage2() {
	SolutionStage2()
	// Output:
	// Stage 2: 6227972
}
