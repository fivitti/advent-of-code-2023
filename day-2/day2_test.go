package day2

import (
	"advent-of-code-2023/utils/require"
	"strings"
	"testing"
)

func TestParseGame(t *testing.T) {
	t.Run("Game 1", func(t *testing.T) {
		// Arrange
		input := "Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green"

		// Act
		game := parseGame(input)

		// Assert
		require.Equal(t, 1, game.id)
		require.Len(t, game.rounds, 3)
		require.Equal(t, 3, game.rounds[0].blue)
		require.Equal(t, 4, game.rounds[0].red)
		require.Zero(t, game.rounds[0].green)
		require.Equal(t, 1, game.rounds[1].red)
		require.Equal(t, 2, game.rounds[1].green)
		require.Equal(t, 6, game.rounds[1].blue)
		require.Equal(t, 2, game.rounds[2].green)
		require.Zero(t, game.rounds[2].red)
		require.Zero(t, game.rounds[2].blue)
	})

	t.Run("Game 3", func(t *testing.T) {
		// Arrange
		input := "Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red"

		// Act
		game := parseGame(input)

		// Assert
		require.Equal(t, 3, game.id)
		require.Len(t, game.rounds, 3)
		require.Equal(t, 8, game.rounds[0].green)
		require.Equal(t, 20, game.rounds[0].red)
		require.Equal(t, 6, game.rounds[0].blue)
		require.Equal(t, 5, game.rounds[1].blue)
		require.Equal(t, 4, game.rounds[1].red)
		require.Equal(t, 13, game.rounds[1].green)
		require.Equal(t, 5, game.rounds[2].green)
		require.Equal(t, 1, game.rounds[2].red)
		require.Zero(t, game.rounds[2].blue)
	})
}

func TestGameIsPossible(t *testing.T) {
	t.Run("Game 1", func(t *testing.T) {
		// Arrange
		input := "Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green"
		bag := round{blue: 14, red: 12, green: 13}

		// Act
		game := parseGame(input)

		// Assert
		require.True(t, game.isPossible(bag))
	})

	t.Run("Game 3", func(t *testing.T) {
		// Arrange
		input := "Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red"
		bag := round{blue: 14, red: 12, green: 13}

		// Act
		game := parseGame(input)

		// Assert
		require.False(t, game.isPossible(bag))
	})
}

func TestParseGames(t *testing.T) {
	// Arrange
	input := strings.Join([]string{
		"Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green",
		"Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red",
	}, "\n")
	reader := strings.NewReader(input)

	// Act
	games := parseGames(reader)

	// Assert
	require.Len(t, games, 2)
}

func TestSumPossibleGames(t *testing.T) {
	// Arrange
	input := strings.Join([]string{
		"Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green",
		"Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue",
		"Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red",
		"Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red",
		"Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green",
	}, "\n")
	reader := strings.NewReader(input)

	// Act
	sum := sumPossibleGames(reader)

	// Assert
	require.Equal(t, 8, sum)
}

func TestGameGetFewerNumberOfCubesForPossible(t *testing.T) {
	t.Run("Game 1", func(t *testing.T) {
		// Arrange
		input := "Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green"
		game := parseGame(input)

		// Act
		bag := game.getFewerNumberOfCubesForPossible()

		// Assert
		require.Equal(t, 4, bag.red)
		require.Equal(t, 2, bag.green)
		require.Equal(t, 6, bag.blue)
	})
}

func TestRoundPower(t *testing.T) {
	round := round{blue: 6, red: 4, green: 2}
	require.Equal(t, 48, round.power())
}

func TestSumPowerOfFewestNumberOfCubes(t *testing.T) {
	// Arrange
	input := strings.Join([]string{
		"Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green",
		"Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue",
		"Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red",
		"Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red",
		"Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green",
	}, "\n")
	reader := strings.NewReader(input)

	// Act
	sum := sumPowersOfFewestNumberOfCubes(reader)

	// Assert
	require.Equal(t, 2286, sum)
}

func ExampleSolutionStage1() {
	SolutionStage1()
	// Output:
	// Stage 1: 2541
}

func ExampleSolutionStage2() {
	SolutionStage2()
	// Output:
	// Stage 2: 66016
}
