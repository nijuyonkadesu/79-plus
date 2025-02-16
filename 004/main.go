package main

import "fmt"

func main() {
	nums := []int{3, 2, 3}
	fmt.Println(overachievedElements(nums))
	fmt.Println(booreVoting(nums))
}

/*
There can't be more than 2 elements present in an list of 'n' items.
This algo exploits that fact....

umm, this could be used in multiple places...
like the most frequently occuring element etc...
*/
func booreVoting(arr []int) []int {
	candidate1, candidate2, count1, count2 := 0, 1, 0 , 0
	bar := len(arr)/3

	for _, val := range arr {
		if candidate1 == val {
			count1++ 
		} else if candidate2 == val { 
			count2++
		} else if count1 == 0 {
			candidate1, count1 = val, 1
		} else if count2 == 0 {
			candidate2, count2 = val, 1 
		} else {
			count1--
			count2--
		}
	}
	// idea is, the frequently occuring numbers, will somehow slip through the if-else ladder.
	// because of their count (occurances). it is what it is. 

	count1 = 0
	count2 = 0
	for _, val := range arr {
		if candidate1 == val {
			count1++
		}
		if candidate2 == val {
			count2++
		}
	}

	res := make([]int, 0, 2)

	if count1 > bar {
		res = append(res, candidate1)
	} 
	if count2 > bar {
		res = append(res, candidate2)
	} 

	return res
}

func overachievedElements(arr []int) []int {
	m := make(map[int]int)
	res := []int{}
	bar := len(arr) / 3

	for _, val := range arr {
		m[val] += 1
	}

	for key, value := range m {
		if value > bar {
			res = append(res, key)
		}
	}

	return res
}
