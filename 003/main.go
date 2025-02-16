package main

import "fmt"

func main() {
	// nums := []int{-1, -2}
	nums := []int{-8, 0, 0, 0, 0, 5, 4, -1, 7, 8}
	fmt.Println(largest(nums))
}

func largest(nums []int) int {
	maxSum := nums[0]
	subsetSum := 0

	endIdx := 0
	offset := 1

	for idx, num := range nums {
		subsetSum += num

		if subsetSum > maxSum {
			maxSum = subsetSum
			endIdx = idx
			offset--
		} else if subsetSum < 0 {
			subsetSum = 0 // resetting the counter. which is equivalent to reseizing the subarray
			offset = 1
		}
	}
	fmt.Println("subset: ", nums[endIdx+offset:endIdx+1])
	return maxSum
}
