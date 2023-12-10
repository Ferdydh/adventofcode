package main

import (
	_ "embed"
	"fmt"
	"github.com/Ferdydh/adventofcode/utils"
)

//go:embed input
var fileString string

// Possible optimization:
// - Tail recursion
// - Eliminate checkAllZero with only checking the first and last (the numbers are monotonic, so it's safe to assume this)
// - Change the recursiveLoop structure

// But I already spent too much time today, let's make tomorrow's prettier

func main() {
	data, err := utils.StringToTwoDArray(fileString)

	if err != nil {
		panic(err)
	}

	partOne(data)
	partTwo(data)
}

func partTwo(input [][]int) {
	res := 0
	for _, arr := range input {
		res += recursiveLoopReversed(arr)
	}
	fmt.Println(res)
}

func partOne(input [][]int) {
	res := 0
	for _, arr := range input {
		res += recursiveLoop(arr)
	}
	fmt.Println(res)
}

func recursiveLoop(input []int) int {
	difArray := numberArrayToDifferenceArray(input)
	isEnd := checkAllZero(difArray)

	if isEnd {
		return input[0]
	}

	//	Assume we have found the result
	delta := recursiveLoop(difArray)
	return delta + input[len(input)-1]
}

func recursiveLoopReversed(input []int) int {
	difArray := numberArrayToDifferenceArray(input)
	isEnd := checkAllZero(difArray)

	if isEnd {
		return input[0]
	}

	//	Assume we have found the result
	delta := recursiveLoopReversed(difArray)
	return input[0] - delta
}

// Naive first approach
//func recursiveLoop(input []int) []int {
//	difArray := numberArrayToDifferenceArray(input)
//	isEnd := checkAllZero(difArray)
//
//	if isEnd {
//		res := append(input, input[len(input)-1])
//		fmt.Println(res)
//		return res
//	}
//
//	//	Assume we have found the result
//	newDifArray := recursiveLoop(difArray)
//
//	numToAppend := newDifArray[len(newDifArray)-1] + input[len(input)-1]
//	input = append(input, numToAppend)
//	fmt.Println(input)
//	return input
//}

// Helper Methods
func numberArrayToDifferenceArray(input []int) []int {
	var result []int

	for i := 0; i < len(input)-1; i++ {
		result = append(result, input[i+1]-input[i])
	}

	return result
}

func checkAllZero(input []int) bool {
	for _, cur := range input {
		if cur != 0 {
			return false
		}
	}

	return true
}
