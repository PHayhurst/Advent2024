package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	wordCount := 0
	frequencyMap := make(map[int]int)
	var leftSidenumbers []int

	for scanner.Scan() {
		wordCount++

		currNum, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Printf("Invalid number: %s (Error: %v)\n", scanner.Text(), err)
			continue
		}

		if wordCount%2 == 0 {
			leftSidenumbers = append(leftSidenumbers, currNum)
		} else {
			frequencyMap[currNum] += 1
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	runningTotalSimilarityScore := 0
	for i := 0; i < len(leftSidenumbers); i++ {
		runningTotalSimilarityScore += leftSidenumbers[i] * frequencyMap[leftSidenumbers[i]]
	}

	fmt.Printf("Total Similarity Score: %d\n", runningTotalSimilarityScore)
}
