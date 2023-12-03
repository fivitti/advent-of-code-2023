package day3

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if b > a {
		return a
	}
	return b
}

type coordinate [2]int

func (c coordinate) getRow() int {
	return c[0]
}

func (c coordinate) getColumn() int {
	return c[1]
}

type number struct {
	value int
	begin coordinate
}

func newNumberFromCharactersAndEnd(characters []byte, end coordinate) number {
	value := 0
	for _, character := range characters {
		value = value*10 + int(character-'0')
	}
	return number{value, coordinate{end.getRow(), end.getColumn() - len(characters) + 1}}
}

func (n number) getWidth() int {
	width := 0
	value := n.value
	for value != 0 {
		value /= 10
		width++
	}
	return width
}

func (n number) getBegin() coordinate {
	return n.begin
}

func (n number) getEnd() coordinate {
	return coordinate{n.begin.getRow(), n.begin.getColumn() + n.getWidth() - 1}
}

type symbol struct {
	character  byte
	coordinate coordinate
}

func (s symbol) isGear() bool {
	return s.character == '*'
}

type bound [2]coordinate

func (b bound) getLeftUpCorner() coordinate {
	return b[0]
}

func (b bound) getRightDownCorner() coordinate {
	return b[1]
}

func isDigit(c byte) bool {
	return c >= '0' && c <= '9'
}

type schematic [][]byte

func (s schematic) getNumberOfRows() int {
	return len(s)
}

func (s schematic) getNumberOfColumns() int {
	return len(s[0])
}

func (s schematic) findNumbers() []number {
	numbers := make([]number, 0)
	for row := 0; row < s.getNumberOfRows(); row++ {
		currentDigits := make([]byte, 0)
		for column := 0; column < s.getNumberOfColumns(); column++ {
			character := s[row][column]
			if isDigit(character) {
				currentDigits = append(currentDigits, character)
			} else if len(currentDigits) > 0 {
				numbers = append(
					numbers,
					newNumberFromCharactersAndEnd(
						currentDigits,
						coordinate{row, column - 1},
					),
				)
				currentDigits = make([]byte, 0)
			}
		}
		if len(currentDigits) > 0 {
			numbers = append(
				numbers,
				newNumberFromCharactersAndEnd(
					currentDigits,
					coordinate{row, s.getNumberOfColumns() - 1},
				),
			)
		}
	}
	return numbers
}

func (s schematic) getBounds(n number) bound {
	return bound{
		coordinate{
			max(n.getBegin().getRow()-1, 0),
			max(n.getBegin().getColumn()-1, 0),
		},
		coordinate{
			min(n.getEnd().getRow()+1, s.getNumberOfRows()-1),
			min(n.getEnd().getColumn()+1, s.getNumberOfColumns()-1),
		},
	}
}

func isSymbol(b byte) bool {
	if isDigit(b) {
		return false
	}
	if b == '.' {
		return false
	}
	return true
}

func (s schematic) findSymbolsInBounds(b bound) []symbol {
	symbols := make([]symbol, 0)
	for row := b.getLeftUpCorner().getRow(); row <= b.getRightDownCorner().getRow(); row++ {
		for column := b.getLeftUpCorner().getColumn(); column <= b.getRightDownCorner().getColumn(); column++ {
			character := s[row][column]
			if isSymbol(character) {
				symbols = append(symbols, symbol{character, coordinate{row, column}})
			}
		}
	}
	return symbols
}

func (s schematic) findPartNumbers() []number {
	numbers := make([]number, 0)
	for _, number := range s.findNumbers() {
		bounds := s.getBounds(number)
		symbols := s.findSymbolsInBounds(bounds)
		if len(symbols) != 0 {
			numbers = append(numbers, number)
		}
	}
	return numbers
}

type gear [2]number

func (g gear) calculateRatio() int {
	return g[0].value * g[1].value
}

func (s schematic) findGearParts() []gear {
	gears := map[coordinate][]number{}

	for _, number := range s.findNumbers() {
		bounds := s.getBounds(number)
		symbols := s.findSymbolsInBounds(bounds)
		for _, symbol := range symbols {
			if !symbol.isGear() {
				continue
			}
			gears[symbol.coordinate] = append(gears[symbol.coordinate], number)
		}
	}

	gearParts := make([]gear, 0, len(gears))
	for _, numbers := range gears {
		if len(numbers) != 2 {
			continue
		}
		gearParts = append(gearParts, gear{numbers[0], numbers[1]})
	}
	return gearParts
}

func sumNumbers(numbers []number) int {
	sum := 0
	for _, number := range numbers {
		sum += number.value
	}
	return sum
}

func sumGearRatios(gears []gear) int {
	sum := 0
	for _, gear := range gears {
		sum += gear.calculateRatio()
	}
	return sum
}

func newSchematic(reader io.Reader) schematic {
	input, _ := io.ReadAll(reader)
	schematicRaw := bytes.Split(input, []byte("\n"))
	return schematicRaw
}

func SolutionStage1() {
	dataFile, err := os.Open("data.txt")
	if err != nil {
		log.Fatalf("could not open data file: %v", err)
	}
	defer dataFile.Close()

	schematic := newSchematic(dataFile)
	numbers := schematic.findPartNumbers()
	sum := sumNumbers(numbers)

	fmt.Printf("Stage 1: %d\n", sum)
}

func SolutionStage2() {
	dataFile, err := os.Open("data.txt")
	if err != nil {
		log.Fatalf("could not open data file: %v", err)
	}
	defer dataFile.Close()

	schematic := newSchematic(dataFile)
	gears := schematic.findGearParts()
	sum := sumGearRatios(gears)

	fmt.Printf("Stage 2: %d\n", sum)
}
