package main

// go fucking read this: https://www.nayuki.io/page/next-lexicographical-permutation-algorithm
// ref: https://math.stackexchange.com/questions/976594/finding-next-permutation-of-a-number

import (
	"fmt"
	"sort"
)

func nextPermutation(num []int) []int {
	// Find the largest decreasing prefix
	pivot := -1 
	for i := len(num) - 1; i > 0; i-- {
		if num[i-1] >= num[i] {
			continue
		} else {
			pivot = i - 1
			break
		}
	}
	// at the greatest permutation
	if pivot == -1 {
		sort.Ints(num)
		return num
	}

	// current state: xxxx7[largestnum]
	//                    | ^^^^^^^^ -> note: reversing this gives the smallest possible number.
	//                    |-> increasing it ever slightly - by swapping pivot with the smallest digit in the suffix
	// 						  that is greater than the current pivot. Now, reverse the largestnum suffix.
	//                    (The number you get this way - the original number) will be the smallest amongst all other possible arrangements.
	successor := len(num) - 1
	for successor > pivot {
		if num[successor] > num[pivot] {
			break
		}
		successor--
	}
	num[pivot], num[successor] = num[successor], num[pivot]

	// reverse suffix alone
	left := pivot + 1
	right := len(num) - 1
	for left < right {
		num[left], num[right] = num[right], num[left]
		left++
		right--
	}
	return num
}

func main() {
	numbers := []int{2, 1, 0}
	fmt.Println(nextPermutation(numbers))
}

/**
 1, 2, 3, 4
[x][x][x][x]

 4, 3, 2, 1

14253876
67835241
^. the pivot (the no smaller than the next one)

all permutations of 3 bit number (truth table)
000
001
010
011
100
101
110
111

1234
1243
1423
4123

4132
4312

4321

**/

/**
digits, upper, lower & ignore spaces / punctuations
0, 1, 2, ... 9

give the next greater permutation of { x, y , z }, which is > original
if nothing is found, just return the sorted array.

Example:
1, 3, 2 -> 2, 1, 3
swap the next greatest element of a[0] with a[0]
but, only once. otherwise, you'll not get the next greatest permutation.

3, 2, 1 -> 1, 2, 3 (simple sort)
**/

/**
// 2025-02-11 - I should have iterated backwards... simple mistake
func swapOnce(arr []int) []int {
	lastIdx := 0
	for idx, value := range arr {
		if value > arr[lastIdx] {
			arr[idx] = arr[idx] ^ arr[lastIdx]
			arr[lastIdx] = arr[idx] ^ arr[lastIdx]
			arr[idx] = arr[idx] ^ arr[lastIdx]
			break
		}
		lastIdx = idx
	}

	if lastIdx == len(arr) - 1 {
		sort.Ints(arr)
	}
	return arr
}
**/
