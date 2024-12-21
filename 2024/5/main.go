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
	rules, pages, err := extractRulesandPages("input.txt")
	if err != nil {
		log.Fatalf("Error reading rules and pages from file: %v", err)
	}
	filteredPages := FilterRightOrderPages(pages, rules)

	log.Printf("Sum of middle pages: %d", sumMiddlePages(filteredPages))

}

func extractRulesandPages(filename string) (map[int][]int, [][]int, error) {
	rules := make(map[int][]int)
	pages := make([][]int, 0)

	file, err := os.Open(filename)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// I want to read until I find a blank line
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}

		parts := strings.Split(line, "|")
		if len(parts) != 2 {
			log.Fatalf("invalid line: %s", line)
		}
		key, err1 := strconv.Atoi(parts[0])
		value, err2 := strconv.Atoi(parts[1])
		if err1 != nil || err2 != nil {
			return nil, nil, fmt.Errorf("failed to convert rule or page number: %w", err)
		}
		if cVal, ok := rules[key]; ok {
			rules[key] = append(cVal, value)
		} else {
			rules[key] = []int{value}
		}
	}
	if err := scanner.Err(); err != nil {
		return nil, nil, fmt.Errorf("failed to read file: %w", err)
	}

	// Now I want to read the pages
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ",")
		page := make([]int, len(parts))
		for i, part := range parts {
			page[i], err = strconv.Atoi(part)
			if err != nil {
				return nil, nil, fmt.Errorf("failed to convert page number: %w", err)
			}
		}
		pages = append(pages, page)
	}
	if err := scanner.Err(); err != nil {
		return nil, nil, fmt.Errorf("failed to read file: %w", err)
	}
	return rules, pages, nil
}

func FilterRightOrderPages(pages [][]int, rules map[int][]int) [][]int {
	filteredPages := make([][]int, 0)
	for _, page := range pages {
		if IsRightOrder(page, rules) {
			filteredPages = append(filteredPages, page)
		}
	}
	return filteredPages
}

func IsRightOrder(page []int, rules map[int][]int) bool {
	for i := len(page) - 1; i > 0; i-- {
		if pageRules, ok := rules[page[i]]; ok {
			for j := i - 1; j >= 0; j-- {
				if contains(pageRules, page[j]) {
					return false
				}
			}
		}
	}
	return true
}

func contains(slice []int, value int) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}
	return false
}

func sumMiddlePages(pages [][]int) int {
	sum := 0
	for _, page := range pages {
		// middle item
		sum += page[len(page)/2]
	}
	return sum
}
