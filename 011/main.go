package main

import "fmt"

func main() {
	nums := []int{5, 6, 7, 8, 0, 1, 2}
	fmt.Println(findPeak(nums))
}

/*
Case 1 where there is only one peak.
- pick a element:
  - see if a[i-1] < a[i] -> increasing, nuke left half
  - see if a[i] > a[i+1] -> decreasing, nuke right half
  - the mid could be the peak

for multiple peaks, the question only asks to find any peak. so it should be fine.
there will be atleast one peak in any half we search. (that's what they said), and I have not got
motivation to dig deep into this pointless question.
*/
func findPeak(nums []int) int {
	start, end := 0, len(nums)-1

	i := 1
	if len(nums) < 100 {
		for i < len(nums)-1 {
			if nums[i-1] < nums[i] && nums[i] > nums[i+1] {
				return i
			}
			i++
		}
		return -1
	}

	for start < end {
		mid := start + (end-start)/2

		if nums[mid-1] > nums[mid] {
			start = mid + 1
		} else if nums[mid] > nums[mid+1] {
			end = mid + 1
		} else {
			return mid
		}
	}

	return -1
}

func findPeakElegant(nums []int) int {
	start, end := 0, len(nums)-1

	for start < end {
		mid := start + (end-start)/2
		// This kind of if conditions even works for array with 2... gotta take pen and paper to really see what is happening...
		if nums[mid] > nums[mid+1] {
			end = mid
		} else {
			start = mid + 1
		}
	}

	return end
}
