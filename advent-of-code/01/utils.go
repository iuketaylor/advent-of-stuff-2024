package utils

import (
	"os"
	"log"
	"bufio"
	"strings"
	"strconv"
)

func SeparateLists(filePath string) ([]int, []int) {
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
