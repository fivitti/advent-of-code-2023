package day2

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

type round struct {
	blue  int
	red   int
	green int
}

func (r round) power() int {
	return r.blue * r.red * r.green
}

type game struct {
	id     int
	rounds []round
}

func (g game) isPossible(bag round) bool {
	for _, round := range g.rounds {
		if round.blue > bag.blue || round.red > bag.red || round.green > bag.green {
			return false
		}
	}
	return true
}

func (g game) getFewerNumberOfCubesForPossible() round {
	bag := round{}
	for _, round := range g.rounds {
		bag.blue = max(bag.blue, round.blue)
		bag.red = max(bag.red, round.red)
		bag.green = max(bag.green, round.green)
	}
	return bag
}

// Each game is listed with its ID number (like the 11 in Game 11: ...)
// followed by a semicolon-separated list of subsets of cubes that were
// revealed from the bag (like 3 red, 5 green, 4 blue).
// Example: Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
func parseGame(input string) game {
	gamePart, roundsPart, _ := strings.Cut(input, ":")

	// Game ID
	gameIDRaw := strings.TrimPrefix(gamePart, "Game ")
	gameID, _ := strconv.Atoi(gameIDRaw)

	// Rounds
	roundParts := strings.Split(roundsPart, ";")
	rounds := make([]round, 0, len(roundParts))
	for _, roundPart := range roundParts {
		roundPart := strings.TrimSpace(roundPart)
		cubeParts := strings.Split(roundPart, ",")
		round := round{}

		for _, cubePart := range cubeParts {
			cubePart := strings.TrimSpace(cubePart)
			countRaw, color, _ := strings.Cut(cubePart, " ")
			count, _ := strconv.Atoi(countRaw)

			switch color {
			case "blue":
				round.blue = count
			case "red":
				round.red = count
			case "green":
				round.green = count
			}
		}

		rounds = append(rounds, round)
	}

	return game{
		id:     gameID,
		rounds: rounds,
	}
}

func parseGames(reader io.Reader) []game {
	games := make([]game, 0)

	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		games = append(games, parseGame(scanner.Text()))
	}

	return games
}

func sumPossibleGames(reader io.Reader) int {
	games := parseGames(reader)
	sum := 0

	for _, game := range games {
		bag := round{blue: 14, red: 12, green: 13}
		if game.isPossible(bag) {
			sum += game.id
		}
	}

	return sum
}

func sumPowersOfFewestNumberOfCubes(reader io.Reader) int {
	games := parseGames(reader)
	sum := 0

	for _, game := range games {
		bag := game.getFewerNumberOfCubesForPossible()
		power := bag.power()
		sum += power
	}

	return sum
}

func SolutionStage1() {
	data, err := os.Open("data.txt")
	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}
	defer data.Close()

	sum := sumPossibleGames(data)
	fmt.Printf("Stage 1: %d\n", sum)
}

func SolutionStage2() {
	data, err := os.Open("data.txt")
	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}
	defer data.Close()

	sum := sumPowersOfFewestNumberOfCubes(data)
	fmt.Printf("Stage 2: %d\n", sum)
}
