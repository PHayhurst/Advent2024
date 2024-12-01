package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func absIntDiff(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	wordCount := 0
	var leftSidenumbers []int
	var rightSidenumbers []int

	for scanner.Scan() {
		wordCount++

		currNum, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Printf("Invalid number: %s (Error: %v)\n", scanner.Text(), err)
			continue
		}
		if(wordCount%2 == 0){
			leftSidenumbers = append(leftSidenumbers, currNum);
		}else{
			rightSidenumbers = append(rightSidenumbers, currNum)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	sort.Ints(leftSidenumbers);
	sort.Ints(rightSidenumbers);
	
	runningTotalDifference := 0
	for i := 0; i < len(leftSidenumbers); i++ {
		
		runningTotalDifference += absIntDiff(leftSidenumbers[i], rightSidenumbers[i])
	}
	fmt.Printf("Total Difference: %d\n", runningTotalDifference)
}
