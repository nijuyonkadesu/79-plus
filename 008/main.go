package main

import "fmt"

func main() {
	nums := []int{0, 2}
	fmt.Println(largestProduct(nums))
	nums = []int{2, 3, -2, 4}
	fmt.Println(largestProduct(nums))
	nums = []int{3, -1, 4}
	fmt.Println(largestProduct(nums))

}

func largestProduct(nums []int) int {
	minProduct := nums[0]
	maxProduct := nums[0]
	result := nums[0]

	for i := 1; i < len(nums); i++ {
		num := nums[i]
		if num < 0 {
			minProduct, maxProduct = maxProduct, minProduct
		}
		// 0 is handled implicitly because of the min / max
		minProduct = min(num, minProduct*num)
		maxProduct = max(num, maxProduct*num)
		if maxProduct > result {
			result = maxProduct
		}
	}

	return result
}

func goodTryButSorry(nums []int) int {
	subarrayProduct := 1
	maxProduct := nums[0]

	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			subarrayProduct = 1
			maxProduct = max(maxProduct, 0)
			continue
		}

		subarrayProduct *= nums[i]
		maxProduct = max(maxProduct, subarrayProduct)

		if subarrayProduct == 0 {
			subarrayProduct = 1
		}
	}
	return maxProduct
}

//  0 -> kills the product
// -1 -> flips the product which is fine
