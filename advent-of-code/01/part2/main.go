package main

import (
	"fmt"
	utils "advent-of-code/01"
)

func main() {
	leftIds, rightIds := utils.SeparateLists("../input.txt")
	
    rightCounts := make(map[int]int)
    for _, num := range rightIds {
        rightCounts[num]++
    }

    similarityScore := 0
    
    for _, leftNum := range leftIds {
        count := rightCounts[leftNum] 
        similarityScore += leftNum * count
    }

    fmt.Println("Similarity Score:", similarityScore)
}
