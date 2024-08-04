package main

import "fmt"

func main() {

	fmt.Println("Two Number Sum task")

	// Example array
	array := []int{1, 2, 3, 4, 5}

	result := TwoNumberSum(array, 6)

	fmt.Println(result)
}

func TwoNumberSum(array []int, target int) []int {

	nums := make(map[int]bool)

	for i := 0; i < len(array); i++ {
		if nums[target-array[i]] {
			return []int{target - array[i], array[i]}
		}
		nums[array[i]] = true
	}
	return make([]int, 0)
}
