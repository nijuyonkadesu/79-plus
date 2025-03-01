package sortloc

func Sort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	mid := len(arr) / 2
	left := Sort(arr[:mid])
	right := Sort(arr[mid:])

	return merge(left, right)
}

func merge(left, right []int) []int {
	buff := make([]int, 0, len(left)+len(right))
	l, r := 0, 0

	for l < len(left) && r < len(right) {
		if left[l] < right[r] {
			buff = append(buff, left[l])
			l++
		} else {
			buff = append(buff, right[r])
			r++
		}
	}
	buff = append(buff, left[l:]...)
	buff = append(buff, right[r:]...)

	return buff
}

/*

Merge sort is a nightmare for cpu cache fyi
*/
