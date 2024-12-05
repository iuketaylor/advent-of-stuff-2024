package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	leftIds, rightIds := separateLists("./input.txt")

	sort.Ints(leftIds)
	sort.Ints(rightIds)

	totalDistance := 0

	for i := 0; i < len(leftIds); i++ {
		distance := int(math.Abs(float64(leftIds[i] - rightIds[i])))
		totalDistance += distance
	}

	rightCounts := make(map[int]int)
	for _, num := range rightIds {
		rightCounts[num]++
	}

	similarityScore := 0

	for _, leftNum := range leftIds {
		count := rightCounts[leftNum]
		similarityScore += leftNum * count
	}

	fmt.Println("Total Distance:", totalDistance)
	fmt.Println("Similarity Score:", similarityScore)
}

func separateLists(filePath string) ([]int, []int) {
	input, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Error reading input.txt", err)
	}
	defer input.Close()

	var leftIds []int
	var rightIds []int

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		parts := strings.Fields(scanner.Text())

		left, err := strconv.Atoi(parts[0])
		if err != nil {
			log.Fatal(err)
		}
		leftIds = append(leftIds, left)

		right, err := strconv.Atoi(parts[1])
		if err != nil {
			log.Fatal(err)
		}
		rightIds = append(rightIds, right)
	}

	return leftIds, rightIds
}
