package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	filename := "input.txt"

	rules, pages, err := loadRulesAndPages(filename)
	if err != nil {
		log.Fatalf("Error reading rules and pages from file '%s': %v", filename, err)
	}

	filteredPages := getValidPages(pages, rules)
	sum := calculateMiddlePageSum(filteredPages)
	log.Printf("Sum of middle pages: %d", sum)
}

// loadRulesAndPages reads rules and pages from the given file.
func loadRulesAndPages(filename string) (map[int][]int, [][]int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	rules, err := parseRules(bufio.NewScanner(file))
	if err != nil {
		return nil, nil, fmt.Errorf("failed to parse rules: %w", err)
	}

	pages, err := parsePages(bufio.NewScanner(file))
	if err != nil {
		return nil, nil, fmt.Errorf("failed to parse pages: %w", err)
	}

	return rules, pages, nil
}

// parseRules parses the rules section from the input file.
func parseRules(scanner *bufio.Scanner) (map[int][]int, error) {
	rules := make(map[int][]int)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}

		parts := strings.Split(line, "|")
		if len(parts) != 2 {
			return nil, fmt.Errorf("invalid rule line format: %s", line)
		}

		key, err1 := strconv.Atoi(parts[0])
		value, err2 := strconv.Atoi(parts[1])
		if err1 != nil || err2 != nil {
			return nil, fmt.Errorf("invalid rule key or value in line '%s': %v, %v", line, err1, err2)
		}

		rules[key] = append(rules[key], value)
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading rules: %w", err)
	}

	return rules, nil
}

// parsePages parses the pages section from the input file.
func parsePages(scanner *bufio.Scanner) ([][]int, error) {
	var pages [][]int
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}

		parts := strings.Split(line, ",")
		page := make([]int, len(parts))
		for i, part := range parts {
			value, err := strconv.Atoi(part)
			if err != nil {
				return nil, fmt.Errorf("invalid page number '%s': %w", part, err)
			}
			page[i] = value
		}
		pages = append(pages, page)
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading pages: %w", err)
	}

	return pages, nil
}

// getValidPages filters pages that satisfy the rules.
func getValidPages(pages [][]int, rules map[int][]int) [][]int {
	var filteredPages [][]int
	for _, page := range pages {
		if isPageValid(page, rules) {
			filteredPages = append(filteredPages, page)
		}
	}
	return filteredPages
}

// isPageValid checks if a page satisfies the given rules.
func isPageValid(page []int, rules map[int][]int) bool {
	for i := len(page) - 1; i > 0; i-- {
		if pageRules, ok := rules[page[i]]; ok {
			for j := i - 1; j >= 0; j-- {
				if containsValue(pageRules, page[j]) {
					return false
				}
			}
		}
	}
	return true
}

// containsValue checks if a slice contains a specific value.
func containsValue(slice []int, value int) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}
	return false
}

// calculateMiddlePageSum calculates the sum of the middle elements of each page.
func calculateMiddlePageSum(pages [][]int) int {
	sum := 0
	for _, page := range pages {
		if len(page) > 0 {
			sum += page[len(page)/2]
		}
	}
	return sum
}
