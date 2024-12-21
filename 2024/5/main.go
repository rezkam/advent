package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// Rule represents a dependency rule between two pages
type Rule struct {
	before int
	after  int
}

// parseRule parses a single line into a Rule
func parseRule(line string) (Rule, error) {
	parts := strings.Split(line, "|")
	if len(parts) < 2 {
		return Rule{}, fmt.Errorf("invalid rule format")
	}
	before, err := strconv.Atoi(parts[0])
	if err != nil {
		return Rule{}, fmt.Errorf("failed to convert to number: %w", err)
	}
	after, err := strconv.Atoi(parts[1])
	if err != nil {
		return Rule{}, fmt.Errorf("failed to convert to number: %w", err)
	}
	return Rule{before, after}, nil
}

// parseUpdate parses a single line into a slice of integers
func parseUpdate(line string) ([]int, error) {
	var update []int
	for _, num := range strings.Split(line, ",") {
		n, err := strconv.Atoi(num)
		if err != nil {
			return nil, fmt.Errorf("failed to convert to number: %w", err)
		}
		update = append(update, n)
	}
	return update, nil
}

// parseInput reads the input file and returns the rules and updates
func parseInput(filename string) ([]Rule, [][]int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	var rules []Rule
	var updates [][]int

	scanner := bufio.NewScanner(file)
	parsingRules := true
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			parsingRules = false
			continue
		}

		if parsingRules {
			rule, err := parseRule(line)
			if err != nil {
				return nil, nil, fmt.Errorf("failed to parse rule: %w", err)
			}
			rules = append(rules, rule)
		} else {
			update, err := parseUpdate(line)
			if err != nil {
				return nil, nil, fmt.Errorf("failed to parse update: %w", err)
			}
			updates = append(updates, update)
		}
	}

	return rules, updates, nil
}

// getTopologicalOrder returns the correct order for a sequence using topological sorting
// It also returns a boolean indicating if the sequence is valid
// This is using Kahn's algorithm for topological sorting
func getTopologicalOrder(update []int, rules []Rule) ([]int, bool) {
	// graph map to keep track of outgoing edges for each node
	graph := make(map[int][]int)
	// inDegree map to keep track of incoming edges for each node
	inDegree := make(map[int]int)

	// Initialize graph for all numbers in the update
	for _, num := range update {
		graph[num] = []int{}
		inDegree[num] = 0
	}

	// Add edges from applicable rules
	for _, rule := range rules {
		// only check rules that both before and after are in the update
		// other rules are not relevant to this update
		if contains(update, rule.before) && contains(update, rule.after) {
			graph[rule.before] = append(graph[rule.before], rule.after)
			inDegree[rule.after]++
		}
	}

	// Kahn's algorithm
	// Initialize queue with nodes that have no incoming edges (inDegree = 0)
	// Add nodes to the result in topological order(removing edges as we go) and update inDegree for neighbors
	// Continue until queue is empty and all nodes are processed
	// If the original sequence follows the topological order, the sequence is valid
	//
	// Find initial nodes with no dependencies
	var queue []int
	for num := range graph {
		if inDegree[num] == 0 {
			queue = append(queue, num)
		}
	}

	// Perform topological sort
	var result []int
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		result = append(result, node)

		for _, neighbor := range graph[node] {
			inDegree[neighbor]--
			if inDegree[neighbor] == 0 {
				queue = append(queue, neighbor)
			}
		}
	}

	// Check if the result contains all nodes from the update
	// If not, there was a cycle
	noCycle := len(result) == len(update)

	// Check if the original sequence matches the topological order
	// If it does, the sequence is valid
	isValid := true
	if noCycle {
		orderMap := make(map[int]int)
		for i, num := range result {
			orderMap[num] = i
		}

		// Check if original sequence follows the topological order
		for i := 1; i < len(update); i++ {
			if orderMap[update[i]] < orderMap[update[i-1]] {
				isValid = false
				break
			}
		}
	}

	return result, isValid
}

func contains(slice []int, val int) bool {
	for _, item := range slice {
		if item == val {
			return true
		}
	}
	return false
}

func getMiddleNumber(nums []int) int {
	return nums[len(nums)/2]
}

func main() {

	rules, updates, err := parseInput("input.txt")
	if err != nil {
		log.Fatalf("input parse failed err: %v", err)
	}

	// Part 1: Using topological sort to validate sequences
	var part1Sum int
	for _, update := range updates {
		_, isValid := getTopologicalOrder(update, rules)
		if isValid {
			part1Sum += getMiddleNumber(update)
		}
	}
	fmt.Printf("Part 1: Sum of middle numbers for correct updates: %d\n", part1Sum)

	// Part 2: Reordering incorrect sequences using topological sort
	var part2Sum int
	for _, update := range updates {
		if correctOrder, valid := getTopologicalOrder(update, rules); !valid {
			part2Sum += getMiddleNumber(correctOrder)
		}
	}
	fmt.Printf("Part 2: Sum of middle numbers after ordering incorrect updates: %d\n", part2Sum)
}
