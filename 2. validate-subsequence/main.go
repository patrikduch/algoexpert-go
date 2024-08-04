package main

import "fmt"

func main() {
	fmt.Println("2. Validate Subsequence task")

	firstArray := []int{1, 2, 3}
	secondArray := []int{1, 2, 3}

	result := IsValidSubsequence(firstArray, secondArray)
	fmt.Println(result)
}

func IsValidSubsequence(array []int, subsequence []int) bool {
	arrIndex := 0
	seqIndex := 0

	for arrIndex < len(array) && seqIndex < len(subsequence) {

		if array[arrIndex] == subsequence[seqIndex] {
			seqIndex++
		}

		arrIndex++
	}

	return seqIndex == len(subsequence)

}
