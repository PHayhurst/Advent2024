package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func isSafeCount(a, b int) bool {
	diff := absIntDiff(a, b)
	return diff >= 1 && diff <= 3
}

func absIntDiff(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}

func getDirection(a, b int) string {
	if a > b {
		return "descending"
	}
	return "ascending"
}

func main() {
	file, err := os.Open("input2.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	safeCount := 0
	for scanner.Scan() {
		currLine := scanner.Text()
		words := strings.Split(currLine, " ")

		if len(words) < 2 {
			continue
		}

		lastNum, _ := strconv.Atoi(words[0])
		currDirection := ""
		isSafe := true

		for i := 1; i < len(words); i++ {
			currNum, err := strconv.Atoi(words[i])
			if err != nil {
				fmt.Printf("Invalid number: %s (Error: %v)\n", words[i], err)
				isSafe = false
				break
			}

			if !isSafeCount(lastNum, currNum) {
				isSafe = false
				break
			}

			if i == 1 {
				currDirection = getDirection(lastNum, currNum)
			} else if getDirection(lastNum, currNum) != currDirection {
				isSafe = false
				break
			}

			lastNum = currNum
		}

		if isSafe {
			safeCount++
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading file: %v\n", err)
	}

	fmt.Printf("Total Safe Count: %d\n", safeCount)
}
