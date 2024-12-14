package main

import (
	"bufio"
	"log"
	"os"
	"sort"
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
	var left, right []int

	for scanner.Scan() {
		line := scanner.Text()
		words := strings.Fields(line)
		if len(words) != 2 {
			log.Fatalf("Invalid input format: %s", line)
		}

		l, err := strconv.Atoi(words[0])
		if err != nil {
			log.Fatalf("Invalid number in input: %s", words[0])
		}
		r, err := strconv.Atoi(words[1])
		if err != nil {
			log.Fatalf("Invalid number in input: %s", words[1])
		}

		left = append(left, l)
		right = append(right, r)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// Sort the slices
	sort.Ints(left)
	sort.Ints(right)

	// Compute the total distance
	distance := 0
	for i := range left {
		distance += abs(left[i] - right[i])
	}

	log.Println("Distance:", distance)

	// Compute the similarity score
	rightMap := make(map[int]int)
	for _, r := range right {
		rightMap[r]++
	}

	similarityScore := 0
	for _, l := range left {
		if count, exists := rightMap[l]; exists {
			similarityScore += l * count
		}
	}

	log.Println("Similarity Score:", similarityScore)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
