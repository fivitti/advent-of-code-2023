package day1

import (
	"strings"
	"testing"
)

func TestGetDigit(t *testing.T) {
	t.Run("digits", func(t *testing.T) {
		for i := 0; i < 10; i++ {
			digit, ok := getDigit(byte(i) + '0')
			if !ok {
				t.Errorf("Expected %d to be a digit", i)
			}
			if digit != i {
				t.Errorf("Expected %d, got %d", i, digit)
			}
		}
	})

	t.Run("letters", func(t *testing.T) {
		for i := 'a'; i <= 'z'; i++ {
			_, ok := getDigit(byte(i))
			if ok {
				t.Errorf("Expected %c to not be a digit", i)
			}
		}
	})
}

func TestGetDigitFromWord(t *testing.T) {
	testCases := []string{
		"zero", "one", "two", "three", "four",
		"five", "six", "seven", "eight", "nine",
		"foozero", "fooone", "footwo", "foothree", "foofour",
		"foofive", "foosix", "fooseven", "fooeight", "foonine",
		"zerobar", "onebar", "twobar", "threebar", "fourbar",
		"fivebar", "sixbar", "sevenbar", "eightbar", "ninebar",
	}

	for i, testCase := range testCases {
		expected := i % 10
		actual, ok := getDigitFromWord(testCase)
		if !ok {
			t.Errorf("Expected %s to be a digit", testCase)
		}
		if actual != expected {
			t.Errorf("Expected %d, got %d", expected, actual)
		}
	}
}

func TestGetDigitFromWordNoDigit(t *testing.T) {
	input := "foo"
	_, ok := getDigitFromWord(input)
	if ok {
		t.Errorf("Expected ok to be false")
	}
}

func TestGetFirstDigit(t *testing.T) {
	testCases := map[string]int{
		"1abc2":       1,
		"pqr3stu8vwx": 3,
		"a1b2c3d4e5f": 1,
		"treb7uchet":  7,
	}

	for input, expected := range testCases {
		t.Run(input, func(t *testing.T) {
			actual, ok := getFirstDigit(input, detectDigitTrivial)
			if !ok {
				t.Errorf("Expected %d, got %d", expected, actual)
			}
			if actual != expected {
				t.Errorf("Expected %d, got %d", expected, actual)
			}
		})
	}
}

func TestGetFirstDigitComplex(t *testing.T) {
	testCases := []string{
		"zero", "fooone", "two3", "barthree4", "4six",
	}

	for i, testCase := range testCases {
		digit, ok := getFirstDigit(testCase, detectDigitComplex)
		if !ok {
			t.Errorf("Expected %s to be a digit", testCase)
		}
		if digit != i {
			t.Errorf("Expected %d, got %d", i, digit)
		}
	}
}

func TestGetFirstDigitNoDigit(t *testing.T) {
	input := "abc"
	_, ok := getFirstDigit(input, detectDigitTrivial)
	if ok {
		t.Errorf("Expected ok to be false")
	}
}

func TestGetLastDigitTrivial(t *testing.T) {
	testCases := map[string]int{
		"1abc2":       2,
		"pqr3stu8vwx": 8,
		"a1b2c3d4e5f": 5,
		"treb7uchet":  7,
	}

	for input, expected := range testCases {
		t.Run(input, func(t *testing.T) {
			actual, ok := getLastDigit(input, detectDigitTrivial)
			if !ok {
				t.Errorf("Expected %d, got %d", expected, actual)
			}
			if actual != expected {
				t.Errorf("Expected %d, got %d", expected, actual)
			}
		})
	}
}

func TestGetLastDigitComplex(t *testing.T) {
	testCases := []string{
		"zero", "onebar", "three2abc", "bar4three", "six4",
	}

	for i, testCase := range testCases {
		digit, ok := getLastDigit(testCase, detectDigitComplex)
		if !ok {
			t.Errorf("Expected %s to be a digit", testCase)
		}
		if digit != i {
			t.Errorf("Expected %d, got %d", i, digit)
		}
	}
}

func TestGetLastDigitNoDigit(t *testing.T) {
	input := "abc"
	_, ok := getLastDigit(input, detectDigitTrivial)
	if ok {
		t.Errorf("Expected ok to be false")
	}
}

func TestCalculateCalibrationTrivialValue(t *testing.T) {
	testCases := map[string]int{
		"1abc2":       12,
		"pqr3stu8vwx": 38,
		"a1b2c3d4e5f": 15,
		"treb7uchet":  77,
	}

	for input, expected := range testCases {
		t.Run(input, func(t *testing.T) {
			actual, ok := calculateCalibrationValue(input, detectDigitTrivial)
			if !ok {
				t.Errorf("Expected %d, got %d", expected, actual)
			}
			if actual != expected {
				t.Errorf("Expected %d, got %d", expected, actual)
			}
		})
	}
}

func TestCalculateCalibrationComplexValue(t *testing.T) {
	testCases := map[string]int{
		"two1nine":         29,
		"eightwothree":     83,
		"abcone2threexyz":  13,
		"xtwone3four":      24,
		"4nineeightseven2": 42,
		"zoneight234":      14,
		"7pqrstsixteen":    76,
	}

	for input, expected := range testCases {
		t.Run(input, func(t *testing.T) {
			actual, ok := calculateCalibrationValue(input, detectDigitComplex)
			if !ok {
				t.Errorf("Expected %d, got %d", expected, actual)
			}
			if actual != expected {
				t.Errorf("Expected %d, got %d", expected, actual)
			}
		})
	}
}

func TestCalculateSumOfCalibrationTrivialValues(t *testing.T) {
	lines := []string{
		"1abc2",
		"pqr3stu8vwx",
		"a1b2c3d4e5f",
		"treb7uchet",
	}
	content := strings.Join(lines, "\n")
	reader := strings.NewReader(content)

	expected := 142
	actual := calculateSumOfCalibrationValues(reader, detectDigitTrivial)

	if actual != expected {
		t.Errorf("Expected %d, got %d", expected, actual)
	}
}

func TestCalculateSumOfCalibrationComplexValues(t *testing.T) {
	lines := []string{
		"two1nine",
		"eightwothree",
		"abcone2threexyz",
		"xtwone3four",
		"4nineeightseven2",
		"zoneight234",
		"7pqrstsixteen",
	}
	content := strings.Join(lines, "\n")
	reader := strings.NewReader(content)

	expected := 281
	actual := calculateSumOfCalibrationValues(reader, detectDigitComplex)

	if actual != expected {
		t.Errorf("Expected %d, got %d", expected, actual)
	}
}

func ExampleSolutionStage1() {
	SolutionStage1()
	// Output:
	// Stage 1: 53334
}

func ExampleSolutionStage2() {
	SolutionStage2()
	// Output:
	// Stage 2: 52834
}
