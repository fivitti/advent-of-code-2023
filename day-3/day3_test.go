package day3

import (
	"advent-of-code-2023/utils/require"
	"testing"
)

func TestNewNumberFromCharactersAndEnd(t *testing.T) {
	// Arrange
	characters := []byte{'1', '2', '3'}
	end := coordinate{1, 3}

	// Act
	number := newNumberFromCharactersAndEnd(characters, end)

	// Assert
	require.Equal(t, 123, number.value)
	require.Equal(t, coordinate{1, 1}, number.begin)
}

func TestNumberGetWidth(t *testing.T) {
	// Arrange
	number := number{123, coordinate{1, 1}}

	// Act
	width := number.getWidth()

	// Assert
	require.Equal(t, 3, width)
}

func TestNumberGetBegin(t *testing.T) {
	// Arrange
	number := number{123, coordinate{1, 1}}

	// Act
	begin := number.getBegin()

	// Assert
	require.Equal(t, coordinate{1, 1}, begin)
}

func TestNumberGetEnd(t *testing.T) {
	// Arrange
	number := number{123, coordinate{1, 1}}

	// Act
	end := number.getEnd()

	// Assert
	require.Equal(t, coordinate{1, 3}, end)
}

func TestIsDigit(t *testing.T) {
	for c := '0'; c <= '9'; c++ {
		require.True(t, isDigit(byte(c)))
	}
}

func TestSchematicGetNumberOfRows(t *testing.T) {
	// Arrange
	schematic := schematic([][]byte{
		[]byte("123"),
		[]byte("456"),
	})

	// Act
	numberOfRows := schematic.getNumberOfRows()

	// Assert
	require.Equal(t, 2, numberOfRows)
}

func TestSchematicGetNumberOfColumns(t *testing.T) {
	// Arrange
	schematic := schematic([][]byte{
		[]byte("123"),
		[]byte("456"),
		[]byte("789"),
	})

	// Act
	numberOfColumns := schematic.getNumberOfColumns()

	// Assert
	require.Equal(t, 3, numberOfColumns)
}

func TestSchematicFindNumbers(t *testing.T) {
	// Arrange
	schematic := schematic([][]byte{
		[]byte("12345"),
		[]byte(".456."),
		[]byte("1.789"),
		[]byte("....."),
	})

	// Act
	numbers := schematic.findNumbers()

	// Assert
	require.Len(t, numbers, 4)
	require.Equal(t, number{12345, coordinate{0, 0}}, numbers[0])
	require.Equal(t, number{456, coordinate{1, 1}}, numbers[1])
	require.Equal(t, number{1, coordinate{2, 0}}, numbers[2])
	require.Equal(t, number{789, coordinate{2, 2}}, numbers[3])
}

func TestSchematicGetBounds(t *testing.T) {
	// Arrange
	schematic := schematic([][]byte{
		[]byte("12345"),
		[]byte(".456."),
		[]byte("1.789"),
		[]byte("....."),
	})

	t.Run("top left corner", func(t *testing.T) {
		number := number{42, coordinate{0, 0}}

		// Act
		bounds := schematic.getBounds(number)

		// Assert
		require.Equal(t, coordinate{0, 0}, bounds.getLeftUpCorner())
		require.Equal(t, coordinate{1, 2}, bounds.getRightDownCorner())
	})

	t.Run("top edge", func(t *testing.T) {
		number := number{12345, coordinate{0, 0}}

		// Act
		bounds := schematic.getBounds(number)

		// Assert
		require.Equal(t, coordinate{0, 0}, bounds.getLeftUpCorner())
		require.Equal(t, coordinate{1, 4}, bounds.getRightDownCorner())
	})

	t.Run("inside", func(t *testing.T) {
		number := number{5, coordinate{1, 2}}

		// Act
		bounds := schematic.getBounds(number)

		// Assert
		require.Equal(t, coordinate{0, 1}, bounds.getLeftUpCorner())
		require.Equal(t, coordinate{2, 3}, bounds.getRightDownCorner())
	})

	t.Run("top right corner", func(t *testing.T) {
		number := number{42, coordinate{0, 3}}

		// Act
		bounds := schematic.getBounds(number)

		// Assert
		require.Equal(t, coordinate{0, 2}, bounds.getLeftUpCorner())
		require.Equal(t, coordinate{1, 4}, bounds.getRightDownCorner())
	})

	t.Run("bottom left corner", func(t *testing.T) {
		number := number{42, coordinate{3, 0}}

		// Act
		bounds := schematic.getBounds(number)

		// Assert
		require.Equal(t, coordinate{2, 0}, bounds.getLeftUpCorner())
		require.Equal(t, coordinate{3, 2}, bounds.getRightDownCorner())
	})

	t.Run("bottom right corner", func(t *testing.T) {
		number := number{42, coordinate{3, 3}}

		// Act
		bounds := schematic.getBounds(number)

		// Assert
		require.Equal(t, coordinate{2, 2}, bounds.getLeftUpCorner())
		require.Equal(t, coordinate{3, 4}, bounds.getRightDownCorner())
	})
}

func TestIsSymbol(t *testing.T) {
	require.True(t, isSymbol('#'))
	require.True(t, isSymbol('$'))
	require.True(t, isSymbol('+'))
	require.True(t, isSymbol('*'))
	require.False(t, isSymbol('.'))
	require.False(t, isSymbol('1'))
}

func TestFindSymbolsInBounds(t *testing.T) {
	// Arrange
	schematic := schematic([][]byte{
		[]byte("1.+23"),
		[]byte("45..."),
		[]byte("...6$"),
		[]byte("...-*"),
	})

	t.Run("missing", func(t *testing.T) {
		bounds := bound{
			coordinate{0, 0},
			coordinate{0, 0},
		}

		// Act
		symbols := schematic.findSymbolsInBounds(bounds)

		// Assert
		require.Empty(t, symbols)
	})

	t.Run("diagonal", func(t *testing.T) {
		bounds := bound{
			coordinate{0, 0},
			coordinate{2, 2},
		}

		// Act
		symbols := schematic.findSymbolsInBounds(bounds)

		// Assert
		require.Len(t, symbols, 1)
		require.Equal(t, byte('+'), symbols[0])
	})

	t.Run("multiple", func(t *testing.T) {
		bounds := bound{
			coordinate{1, 2},
			coordinate{3, 4},
		}

		// Act
		symbols := schematic.findSymbolsInBounds(bounds)

		// Assert
		require.Len(t, symbols, 3)
		require.Equal(t, byte('$'), symbols[0])
		require.Equal(t, byte('-'), symbols[1])
		require.Equal(t, byte('*'), symbols[2])
	})
}

func TestDetectPartNumbers(t *testing.T) {
	// Arrange
	schematic := schematic([][]byte{
		[]byte("467..114.."),
		[]byte("...*......"),
		[]byte("..35..633."),
		[]byte("......#..."),
		[]byte("617*......"),
		[]byte(".....+.58."),
		[]byte("..592....."),
		[]byte("......755."),
		[]byte("...$.*...."),
		[]byte(".664.598.."),
	})

	// Act
	numbers := schematic.findPartNumbers()
	sum := sumNumbers(numbers)

	// Assert
	require.Len(t, numbers, 8)
	require.Equal(t, 467, numbers[0].value)
	require.Equal(t, 35, numbers[1].value)
	require.Equal(t, 633, numbers[2].value)
	require.Equal(t, 617, numbers[3].value)
	require.Equal(t, 592, numbers[4].value)
	require.Equal(t, 755, numbers[5].value)
	require.Equal(t, 664, numbers[6].value)
	require.Equal(t, 598, numbers[7].value)
	require.Equal(t, 4361, sum)
}

func TestCalculateRatio(t *testing.T) {
	// Arrange
	gear := gear{
		number{2, coordinate{0, 0}},
		number{3, coordinate{0, 0}},
	}

	// Act
	ratio := gear.calculateRatio()

	// Assert
	require.Equal(t, 6, ratio)
}

func TestSchematicFindGearParts(t *testing.T) {
	// Arrange
	schematic := schematic([][]byte{
		[]byte("467..114.."),
		[]byte("...*......"),
		[]byte("..35..633."),
		[]byte("......#..."),
		[]byte("617*......"),
		[]byte(".....+.58."),
		[]byte("..592....."),
		[]byte("......755."),
		[]byte("...$.*...."),
		[]byte(".664.598.."),
	})

	// Act
	gearParts := schematic.findGearParts()

	// Assert
	require.Len(t, gearParts, 2)
	require.Equal(t, gearParts[0][0].value, 467)
	require.Equal(t, gearParts[0][1].value, 35)
	require.Equal(t, gearParts[0].calculateRatio(), 16345)
	require.Equal(t, gearParts[1][0].value, 755)
	require.Equal(t, gearParts[1][1].value, 598)
	require.Equal(t, gearParts[1].calculateRatio(), 451490)
}

func ExampleSolutionStage1() {
	SolutionStage1()
	// Output:
	// Stage 1: 526404
}

func ExampleSolutionStage2() {
	SolutionStage2()
	// Output:
	// Stage 2: 84399773
}
