package main

import "fmt"

func main() {
	nums := []int{5, 3, 2, 1, 4}
	fmt.Println(findInversions(nums))

	nums = []int{5, 4, 3, 2, 1}
	fmt.Println(findInversions(nums))
}

/*
Count all the elements that satisfies:

	arr[i] > arr[j] & i < j
*/

func findInversions(nums []int) int {
	swapCounter := 0
	chopTillSingularity(nums, &swapCounter)
	return swapCounter
}

func chopTillSingularity(nums []int, swapCounter *int) []int {

	if len(nums) <= 1 {
		return nums
	}

	pivot := len(nums) / 2
	leftPart := chopTillSingularity(nums[:pivot], swapCounter)
	rightPart := chopTillSingularity(nums[pivot:], swapCounter)

	return conditionalCounter(leftPart, rightPart, swapCounter)
}

func conditionalCounter(leftPart, rightPart []int, swapCounter *int) []int {
	i, j := 0, 0
	buff := make([]int, 0, len(leftPart)+len(rightPart))
	invertedPairs := 0
	for i < len(leftPart) && j < len(rightPart) {
		if leftPart[i] > rightPart[j] {
			buff = append(buff, rightPart[j])
			// all pairs for each j, note: we always have sorted arrays in this func
			invertedPairs += len(leftPart) - i
			j++
		} else {
			buff = append(buff, leftPart[i])
			i++
		}
	}

	*swapCounter += invertedPairs
	buff = append(buff, leftPart[i:]...)
	buff = append(buff, rightPart[j:]...)
	return buff
}
