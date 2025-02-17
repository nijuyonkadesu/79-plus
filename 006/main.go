package main

import "fmt"

/*
A - the repeated number
B - the missing number
*/
func main() {
	// nums := []int{1, 2, 3, 4, 5}
	nums := []int{1, 2, 3, 4, 5, 6, 7, 7, 8, 9, 11, 12, 13, 14}
	fmt.Println(findCorruption(nums))

	nums = []int{4, 3, 6, 2, 1, 1}
	fmt.Println(findCorruption(nums))
}

func findCorruption(nums []int) []int64 {
	// expected & actual sum of sequence
	se := sumOfSequence(nums, true)
	sa := sumOfSequence(nums, false)

	// expected & actual sum of sequence squared
	s2e := sumOfSquareOfSequence(nums, true)
	s2o := sumOfSquareOfSequence(nums, false)

	// (B - A)
	x := se - sa

	// B^2 - A^2  = (B + A) (B - A)
	y := s2e - s2o

	// B + A
	y = y / x

	b := (x + y) / 2
	a := (y - x) / 2

	return []int64{a, b}
}

func sumOfSequence(nums []int, compute bool) int64 {
	// n * (n+1) / 2
	if compute {
		// a := float32(1)
		// n := float32(len(nums))
		// d := float32(1)
		// sum := (n / 2) * (2*a + (n-1)*d)
		n := int64(len(nums))
		sum := int64(n * (n + 1) / 2)
		return sum
	} else {
		sum := int64(0)
		for _, val := range nums {
			sum += int64(val)
		}
		return sum
	}
}

func sumOfSquareOfSequence(nums []int, compute bool) int64 {
	// n (n+1) (2n+1) / 6
	if compute {
		n := int64(len(nums))
		sum := int64(n * (n + 1) * (2*n + 1) / 6)
		return sum
	} else {

		sum := int64(0)
		for _, val := range nums {
			sum += int64(val * val)
		}
		return sum
	}
}
