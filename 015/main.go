package main

import "fmt"

func main() {
	a1 := []int{2, 4, 6, 8, 9, 10}
	a2 := []int{1, 3, 5, 15, 20}
	fmt.Println(median(a1, a2))
}

/*
We know what median. it's always total-length / 2
so, we'll guess the median and see if it's actually closer to the actual median.

Guess, I'll realize truly why this approach doesn't work tomorrow...
*/

func median(arr1, arr2 []int) float64 {
	// Don't do -2
	totalLength := len(arr1) + len(arr2)
	medianIdx := totalLength / 2

	small := arr1
	big := arr2
	if len(arr1) > len(arr2) {
		big, small = small, big
	}

	start := 0
	end := len(small) - 1
	for start <= end {
		guess := (start + end) / 2
		res := combinedIndexFor(guess, small, big)
		combined := res[0]
		match1 := res[1]
		match2 := res[2]
		if combined == float64(medianIdx) {
			return (match1 + match2) / 2
		} else if combined > float64(medianIdx) {
			end = guess - 1
		} else {
			start = guess + 1
		}
	}
	return -1
}

func binarySearchEqualOrSmaller(arr []int, target int) int {
	start := 0
	end := len(arr) - 1

	match := 0

	for start <= end {
		mid := (start + end) / 2
		if arr[mid] <= target {
			match = mid
			start = mid + 1
		} else if arr[mid] > target {
			end = mid - 1
		}
	}
	return match
}

func combinedIndexFor(idx int, small, big []int) []float64 {
	element := small[idx]
	idx1 := float64(idx)
	idx2 := float64(binarySearchEqualOrSmaller(big, element))

	return []float64{idx1 + idx2 + 1, idx1, idx2}
}
