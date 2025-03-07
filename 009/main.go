package main

import "fmt"

func main() {
	nums := []int{2, 5, 6, 0, 0, 1, 2}
	fmt.Println(searching(nums, 5))
	nums = []int{1, 0, 1, 1, 1}
	fmt.Println(searching(nums, 0))
	nums = []int{2, 5, 6, 0, 1, 2}
	fmt.Println(searchingAgain(nums, 5))
	nums = []int{2, 2, 2, 2, 2, 2, 5, 4, 4, 4, 4, 4, 4, 6, 0, 1, 2}
	//           ^s    helf1       ^m  half2   ^l
	//          [              sorted           |     unsorted    ]
	// This is why we shd verify which half within the sorted side our target is actually present. Otherwise, you'd eliminate wrong half
	fmt.Println(searchingAgain(nums, 5))
}

/*
Array is rotated through a pivot.
instead of finding it in a naive way using 2 pointers, I can imaging finding the pivot using binary search

----------> usual sorted array
small, big, bigger, biggest

bigger, biggest, small, big

2,5,6,0,0,1,2

0, 0, 1, 2, 2, 5, 6,

2, 5, 6, 0, 0, 1, 2
*/

/*
No matter what, one half is always sorted.
find that half,
see if your element is present,
then regular binary search
*/
func searchingAgain(nums []int, target int) int {
	start := 0
	end := len(nums) - 1
	i := 0

    if len(nums) < 100 {
        for i < len(nums) {
            if nums[i] == target {
                return i
            }
            i++
        }
        return -1
    }
	
	for start <= end {
		mid := (start + end) / 2
		midItem := nums[mid]
		lastItem := nums[end]
		firstItem := nums[start]

		if midItem == target || firstItem == target || lastItem == target {
			return mid
		}

		if firstItem == midItem && midItem == lastItem {
			start++
			end--
			continue
		}

		if midItem < lastItem {
			// also check if target is < midItem, coz, it might be too big and we'd be skipping `if` unnecessarily
			// also use <= | >= when comparing with target with start | end.
			// coz, == mid will be caught, but who's checking start & end
			if target > midItem && target <= lastItem {
				start = mid + 1
			} else {
				end = mid - 1
			}
		} else {
			if target < midItem && target >= firstItem {
				end = mid - 1
			} else {
				start = mid + 1
			}
		}
	}
	return -1
}

func searching(nums []int, target int) bool {
	start := 0
	end := len(nums) - 1

	for start <= end {
		mid := (start + end) / 2
		midItem := nums[mid]
		firstItem := nums[start]
		lastItem := nums[end]

		if target == midItem || target == firstItem || target == lastItem {
			return true
		}
		if firstItem == midItem && midItem == lastItem {
			start++
			end--
			continue
		}

		if target < midItem && target > firstItem {
			end = mid
		} else if target > midItem && target < lastItem {
			start = mid + 1
		} else if target < midItem && target < firstItem {
			start = mid + 1
		} else {
			end = mid
		}
	}
	return false
}
