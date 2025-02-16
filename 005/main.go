package main

import "fmt"

/*
instead of finding the largest summing subarray, it's about finding the subarray that gives 6.
instead of sum, it's xor.

O(n^2) is possible...
O(n) ??
*/

func main() {
	xorme := []int{4, 2, 2, 6, 4}
	// fmt.Println(xorAllElements(xorme))
	// example()
	fmt.Println(xorAll(xorme, 6))
}

func xorAll(arr []int, target int) int {
	count := 0

	for start := 0; start < len(arr); start++ {
		xors := 0
		offset := 1

		for _, value := range arr[start:] {
			xors ^= value
			offset++
			if xors == target {
				count++
				fmt.Println("subset", arr[start:start+offset], xors)
				offset = 1
			}
		}
	}

	return count
}

func example() {
	a := 5 // 0101 in binary
	b := 3 // 0011 in binary

	fmt.Printf("a: %d (%04b)\n", a, a)
	fmt.Printf("b: %d (%04b)\n", b, b)
	fmt.Printf("a ^ b = %d (%04b)\n", a^b, a^b)
	fmt.Printf("(a ^ b) ^ b = %d (%04b)\n", (a^b)^b, (a^b)^b)
	fmt.Printf("a ^ (a ^ b) = %d (%04b)\n", a^(a^b), a^(a^b))
	fmt.Printf("a ^ 0 = %d (%04b)\n", a^0, a^0)
}

func xorAllElements(arr []int) int {
	xorResult := 0
	for _, num := range arr {
		xorResult ^= num
	}
	return xorResult
}
