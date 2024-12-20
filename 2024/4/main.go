package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	grid, err := readGridFromFile("input.txt")
	if err != nil {
		log.Fatalf("Error reading grid from file: %v", err)
	}

	// Part 1: Count all occurrences of "XMAS"
	xmasCount := countOccurrences(grid, "XMAS")
	fmt.Printf("Count of 'XMAS': %d\n", xmasCount)

	// Part 2: Count all occurrences of "X-MAS"
	xmasShapeCount := countXMASShapes(grid)
	fmt.Printf("Count of 'X-MAS': %d\n", xmasShapeCount)
}

// readGridFromFile reads a grid of runes from a file.
func readGridFromFile(filename string) ([][]rune, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	var grid [][]rune
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		grid = append(grid, []rune(scanner.Text()))
	}
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("failed to read file: %w", err)
	}
	return grid, nil
}

// countOccurrences counts all occurrences of a word in the grid, considering all directions.
func countOccurrences(grid [][]rune, word string) int {
	occurrenceCount := 0
	totalRows := len(grid)
	totalCols := len(grid[0])

	for row := 0; row < totalRows; row++ {
		for col := 0; col < totalCols; col++ {
			for rowStep := -1; rowStep <= 1; rowStep++ {
				for colStep := -1; colStep <= 1; colStep++ {
					if rowStep == 0 && colStep == 0 {
						continue
					}
					if checkWordInDirection(grid, word, row, col, rowStep, colStep) {
						occurrenceCount++
					}
				}
			}
		}
	}
	return occurrenceCount
}

// checkWordInDirection checks if the word exists starting from (startRow, startCol) in the specified direction.
func checkWordInDirection(grid [][]rune, word string, startRow, startCol, rowStep, colStep int) bool {
	totalRows := len(grid)
	totalCols := len(grid[0])

	for i, char := range word {
		currentRow := startRow + rowStep*i
		currentCol := startCol + colStep*i

		if isOutOfBounds(currentRow, currentCol, totalRows, totalCols) || grid[currentRow][currentCol] != char {
			return false
		}
	}
	return true
}

// countXMASShapes counts all X-MAS shapes in the grid.
func countXMASShapes(grid [][]rune) int {
	occurrenceCount := 0
	totalRows := len(grid)
	totalCols := len(grid[0])

	for row := 0; row < totalRows; row++ {
		for col := 0; col < totalCols; col++ {
			if checkXMASShape(grid, row, col) {
				occurrenceCount++
			}
		}
	}
	return occurrenceCount
}

// checkXMASShape verifies if there's an X-MAS shape centered at (row, col).
func checkXMASShape(grid [][]rune, row, col int) bool {
	totalRows := len(grid)
	totalCols := len(grid[0])

	if isOutOfBounds(row, col, totalRows, totalCols) || grid[row][col] != 'A' {
		return false
	}

	diagonals := [][2]int{
		{row - 1, col - 1}, {row - 1, col + 1},
		{row + 1, col - 1}, {row + 1, col + 1},
	}

	for _, coord := range diagonals {
		if isOutOfBounds(coord[0], coord[1], totalRows, totalCols) {
			return false
		}
	}

	return isValidDiagonalPair(grid, diagonals[0], diagonals[3]) &&
		isValidDiagonalPair(grid, diagonals[1], diagonals[2])
}

// isValidDiagonalPair validates that a diagonal contains either "M-S" or "S-M".
func isValidDiagonalPair(grid [][]rune, coord1, coord2 [2]int) bool {
	left := grid[coord1[0]][coord1[1]]
	right := grid[coord2[0]][coord2[1]]
	return (left == 'M' && right == 'S') || (left == 'S' && right == 'M')
}

// isOutOfBounds checks if a position is outside the grid boundaries.
func isOutOfBounds(row, col, totalRows, totalCols int) bool {
	return row < 0 || row >= totalRows || col < 0 || col >= totalCols
}
