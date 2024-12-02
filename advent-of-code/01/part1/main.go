package main

import (
	utils "advent-of-code/01"
	"fmt"
	"math"
	"sort"
)


func main() {
	leftIds, rightIds := utils.SeparateLists("../input.txt")
	
	sort.Ints(leftIds)
	sort.Ints(rightIds)
	
	totalDistance := 0

	for i := 0; i < len(leftIds); i++ {
		distance := int(math.Abs(float64(leftIds[i] - rightIds[i])))
		totalDistance += distance
	}

	fmt.Println("Total Distance:", totalDistance)
}
