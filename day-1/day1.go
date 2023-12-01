package day1

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func getDigit(character byte) (int, bool) {
	return int(character - '0'), character >= '0' && character <= '9'
}

func getDigitFromWord(line string) (int, bool) {
	digits := []string{
		"zero", "one", "two", "three", "four",
		"five", "six", "seven", "eight", "nine",
	}

	for i, digit := range digits {
		if strings.Contains(line, digit) {
			return i, true
		}
	}
	return 0, false
}

type digitDetector func(line string) (int, bool)

func detectDigitTrivial(line string) (int, bool) {
	for i := 0; i < len(line); i++ {
		if digit, ok := getDigit(line[i]); ok {
			return digit, true
		}
	}
	return 0, false
}

func detectDigitComplex(line string) (int, bool) {
	for i := 0; i < len(line); i++ {
		if digit, ok := getDigit(line[i]); ok {
			return digit, true
		}
		if digit, ok := getDigitFromWord(line[i:]); ok {
			return digit, true
		}
	}
	return 0, false
}

func getFirstDigit(line string, isDigit digitDetector) (int, bool) {
	for i := 0; i < len(line); i++ {
		if digit, ok := isDigit(line[0 : i+1]); ok {
			return digit, true
		}
	}
	return 0, false
}

func getLastDigit(line string, isDigit digitDetector) (int, bool) {
	for i := len(line) - 1; i >= 0; i-- {
		if digit, ok := isDigit(line[i:]); ok {
			return digit, true
		}
	}
	return 0, false
}

// The calibration value can be found by combining
// the first digit and the last digit (in that order)
// to form a single two-digit number.
func calculateCalibrationValue(line string, isDigit digitDetector) (int, bool) {
	first, ok := getFirstDigit(line, isDigit)
	if !ok {
		return 0, false
	}
	last, ok := getLastDigit(line, isDigit)
	if !ok {
		return 0, false
	}
	return first*10 + last, true
}

func calculateSumOfCalibrationValues(reader io.Reader, isDigit digitDetector) int {
	sum := 0

	scanner := bufio.NewScanner(reader)

	for scanner.Scan() {
		line := scanner.Text()
		calibrationValue, ok := calculateCalibrationValue(line, isDigit)
		if !ok {
			continue
		}
		sum += calibrationValue
	}
	return sum
}

func SolutionStage1() {
	dataFile, err := os.Open("data.txt")
	if err != nil {
		log.Fatalf("could not open data file: %v", err)
	}
	defer dataFile.Close()

	sum := calculateSumOfCalibrationValues(dataFile, detectDigitTrivial)
	fmt.Printf("Stage 1: %d\n", sum)
}

func SolutionStage2() {
	dataFile, err := os.Open("data.txt")
	if err != nil {
		log.Fatalf("could not open data file: %v", err)
	}
	defer dataFile.Close()

	sum := calculateSumOfCalibrationValues(dataFile, detectDigitComplex)
	fmt.Printf("Stage 2: %d", sum)
}
