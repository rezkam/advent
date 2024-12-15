package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Regex pattern to match `mul(X,Y)` where X and Y are 1-3 digit numbers
	// submatch 1: X and submatch 2: Y
	mulPattern := `mul\((\d{1,3}),(\d{1,3})\)`

	// Regex to match mul(X,Y) instructions and conditional instructions do() and don't()
	// in case of mul(X,Y) submatch 1: X and submatch 2: Y
	condPattern := `do\(\)|don't\(\)|mul\((\d{1,3}),(\d{1,3})\)`

	// Compile the regex
	multRe := regexp.MustCompile(mulPattern)
	condRe := regexp.MustCompile(condPattern)

	totalMulResult := 0
	totalCondMulResult := 0
	mulEnabled := true // state should persist across lines of input

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		memLine := scanner.Text()
		matches := multRe.FindAllStringSubmatch(memLine, -1)

		for _, match := range matches {
			if len(match) == 3 {
				// Convert the matched strings to integers
				x, y := convertToIntegers(match[1], match[2])
				totalMulResult += x * y
			}
		}

		matches = condRe.FindAllStringSubmatch(memLine, -1)
		for _, match := range matches {
			fullMatch := match[0]
			if fullMatch == "do()" {
				mulEnabled = true
			} else if fullMatch == "don't()" {
				mulEnabled = false
			} else if len(match) == 3 && mulEnabled {
				// Convert the matched strings to integers
				x, y := convertToIntegers(match[1], match[2])
				totalCondMulResult += x * y
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	log.Printf("Total multiplication result: %d", totalMulResult)
	log.Printf("Total multiplication result with condition: %d", totalCondMulResult)
}

func convertToIntegers(x, y string) (int, int) {
	xInt, err1 := strconv.Atoi(x)
	yInt, err2 := strconv.Atoi(y)
	if err1 != nil {
		log.Fatalf("Failed to convert X: %s, error: %v", x, err1)
	}
	if err2 != nil {
		log.Fatalf("Failed to convert Y: %s, error: %v", y, err2)
	}
	return xInt, yInt
}
