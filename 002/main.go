package main

import (
	"fmt"
	"sort"
)

func main() {
	// nums := []int{-1, 0, 1, 2, -1, -4}
	nums := []int{-1, 0, 1, 2, -1, -4, 3, 1, -4, -2, 1, 6, -7}
	temp := triplets(nums)
	fmt.Println("idx", "triplet", temp)
}

func triplets(nums []int) [][]int {
	sort.Ints(nums)
	var results [][]int

	// all we do is avoid processing duplicate values, such that we do not produce duplicate items
	for start := 0; start < len(nums)-1; start++ {
		// skip re-processing target [1]
		if start > 0 && nums[start] == nums[start-1] {
			continue
		}
		target := nums[start]
		left := start + 1
		right := len(nums) - 1

		for left < right {
			// goal is to check 0 = a + b + (-target)
			// kek, checking for 0 slows down by ~3 ms kek
			sum := nums[left] + nums[right] + target

			if sum < 0 {
				left++
			} else if sum > 0 {
				right--
			} else {
				results = append(results, []int{target, nums[left], nums[right]})
				left++
				right--

				// skip re-processing values pointed by left & right [2], [3]
				for left < right && nums[left] == nums[left-1] {
					left++
				}
				for left < right && nums[right] == nums[right+1] {
					right--
				}
			}

		}
	}
	return results
}
