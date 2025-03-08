package main

import "fmt"

func main() {
	//nums := []int{5, 6, 7, 8, 0, 1, 2}
	//fmt.Println(leastNumber(&nums))
	//nums = []int{3, 4, 5, 1, 2}
	//fmt.Println(leastNumber(&nums))
	//nums = []int{5, 1, 2, 3, 4}
	//fmt.Println(leastNumber(&nums))
	//nums = []int{0, 1, 2}
	//fmt.Println(leastNumber(&nums))
	//nums = []int{1, 2, 0}
	//fmt.Println(leastNumber(&nums))
	//nums = []int{2, 1, 0}
	//fmt.Println(leastNumber(&nums))
	// nums := []int{284, 287, 289, 293, 295, 298, 0, 3, 8, 9, 10, 11, 12, 15, 17, 19, 20, 22, 26, 29, 30, 31, 35, 36, 37, 38, 42, 43, 45, 50, 51, 54, 56, 58, 59, 60, 62, 63, 68, 70, 73, 74, 81, 83, 84, 87, 92, 95, 99, 101, 102, 105, 108, 109, 112, 114, 115, 116, 122, 125, 126, 127, 129, 132, 134, 136, 137, 138, 139, 147, 149, 152, 153, 154, 155, 159, 160, 161, 163, 164, 165, 166, 168, 169, 171, 172, 174, 176, 177, 180, 187, 188, 190, 191, 192, 198, 200, 203, 204, 206, 207, 209, 210, 212, 214, 216, 221, 224, 227, 228, 229, 230, 233, 235, 237, 241, 242, 243, 244, 246, 248, 252, 253, 255, 257, 259, 260, 261, 262, 265, 266, 268, 269, 270, 271, 272, 273, 277, 279, 281}
	nums := []int{127, 128, 129, 132, 140, 145, 146, 148, 151, 156, 157, 162, 164, 168, 169, 173, 185, 186, 187, 188, 189, 191, 194, 198, 203, 204, 207, 208, 210, 213, 214, 220, 223, 231, 235, 236, 240, 241, 251, 252, 253, 255, 265, 266, 267, 273, 274, 277, 278, 280, 281, 284, 290, 291, 292, 293, 296, 297, 298, 9, 12, 16, 17, 20, 22, 28, 33, 34, 35, 36, 37, 38, 40, 41, 46, 47, 49, 53, 58, 59, 61, 62, 63, 65, 68, 72, 74, 80, 82, 83, 88, 89, 93, 95, 98, 100, 101, 104, 107, 111, 121, 125}
	// fmt.Println(leastNumber(&nums))
	// nums := []int{295, 298, 0}
	fmt.Println(leastNumber(&nums))
}


/*
sigh... wayy more elegant...

okay, it's obvious right...
array is shifted to the left (<-)
so, it's obvious smallest element is pushed to the right side after rotation. 

so, it is sufficient to just check if mid is greater than end.
that way, we'll know smallest element is somewhere on the right. 
if that condition fails, smallest element is somewhere on the left, so nuke the other half. 
*/
func findMin(nums []int) int {
	start, end := 0, len(nums)-1

	for start < end {
		mid := start + (end-start)/2
		if nums[mid] > nums[end] {
			start = mid + 1
		} else {
			end = mid
		}
	}

	return nums[start]
}

/*
If we find the pivot, we'll find the minimum number...
we have access to three items:
 1. start,
 2. mid,
 3. end,

there is no target, instead:
  - choose the half which contains smallest of the three
  - start -> first half
  - last  -> second half
  - mid -> ... I think this is the smallest number. return it.
    (hindsight: okay, the mid is not always the smallest, so track the smallest till the whole array is exhausted)

note: check for start, end boundry before comparing
*/
func leastNumber(arr *[]int) int {
	nums := *arr
	start := 0
	end := len(nums) - 1

	if len(nums) < 100 {
		i := 0
		lowest := nums[0]
		for i < len(nums) {
			lowest = min(lowest, nums[i])
			i++
		}
		return lowest
	}

	for start <= end {
		mid := (start + end) / 2
		firstItem := nums[start]
		lastItem := nums[end]
		midItem := nums[mid]

		target := min(midItem, firstItem, lastItem)

		if target == midItem {
			if nums[mid-1] < target && mid-1 > start {
				end = mid - 1
				fmt.Println("M", nums[start : end+1])
			} else if target > nums[mid+1] && mid+1 < end {
				start = mid + 1
				fmt.Println("M", nums[start : end+1])
			} else {
				return target
			}
		} else if target == lastItem {
			start = mid + 1
			fmt.Println("L", nums[start : end+1])
		} else {
			end = mid - 1
			fmt.Println("F", nums[start : end+1])
		}
	}

	return -1
}
