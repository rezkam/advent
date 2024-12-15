package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	safeCount := 0
	safeCountDampened := 0

	for scanner.Scan() {
		report := scanner.Text()
		levels := strings.Fields(report)

		// Convert levels to integers
		intLevels, err := convertToIntegers(levels)
		if err != nil {
			log.Println("Skipping invalid input due to conversion error:", report)
			continue
		}

		// Check if the report is safe directly
		if isSafe(intLevels) {
			safeCount++
			safeCountDampened++ // Safe without dampener is also safe with dampener
		} else if isSafeWithDampener(intLevels) {
			safeCountDampened++
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	log.Printf("Safe count: %d", safeCount)
	log.Printf("Safe count with dampener: %d", safeCountDampened)
}

// Convert string levels to integers
func convertToIntegers(levels []string) ([]int, error) {
	intLevels := make([]int, len(levels))
	for i, level := range levels {
		intValue, err := strconv.Atoi(level)
		if err != nil {
			return nil, err
		}
		intLevels[i] = intValue
	}
	return intLevels, nil
}

// Check if levels are safe without a dampener
func isSafe(levels []int) bool {
	return isIncreasing(levels) || isDecreasing(levels)
}

// Check if levels are safe with a dampener
func isSafeWithDampener(levels []int) bool {
	for i := range levels {
		// Create a copy of levels excluding the current level
		modified := append([]int{}, levels[:i]...)
		modified = append(modified, levels[i+1:]...)
		if isSafe(modified) {
			return true
		}
	}
	return false
}

// Check if levels are strictly increasing with valid differences
func isIncreasing(levels []int) bool {
	for i := 0; i < len(levels)-1; i++ {
		diff := levels[i+1] - levels[i]
		if diff < 1 || diff > 3 {
			return false
		}
	}
	return true
}

// Check if levels are strictly decreasing with valid differences
func isDecreasing(levels []int) bool {
	for i := 0; i < len(levels)-1; i++ {
		diff := levels[i] - levels[i+1]
		if diff < 1 || diff > 3 {
			return false
		}
	}
	return true
}
