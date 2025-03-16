package main

import (
	"fmt"
	"sort"
)

func main() {
	var T int
	fmt.Scan(&T)
 
	for t := 0; t < T; t++ {
		var N, C int
		fmt.Scan(&N, &C)
 
		stalls := make([]int, N)
		for i := 0; i < N; i++ {
			fmt.Scan(&stalls[i])
		}
 
		fmt.Println(aggressiveCows(stalls, C))
	}
}

/*

max possible difference b/w all the stalls provided is [1, max-min]
our answer lies somewhere between that space.

for a given distance b/w each slot, we'll try to place all the cows by checking
the gap b/w each stall.

if we could place all, that's a possible solution.
now, aim for even ~lower~ solutions.
*maximum solution


*/

func areTheCowsCalmNow(stalls []int, space int, cows int) bool {
	lastOccupiedStall := stalls[0]
	cowsPlacedSoFar := 1

	for i := 0; i < len(stalls); i++ {
		if stalls[i]-lastOccupiedStall >= space {
			lastOccupiedStall = stalls[i]
			cowsPlacedSoFar++
			if cowsPlacedSoFar == cows {
				return true
			}
		}
	}

	return false
}

func aggressiveCows(stalls []int, cows int) int {
	sort.Ints(stalls)
	smol := 1
	big := stalls[len(stalls)-1] - stalls[0]

	preferredSpace := 0

	for smol <= big {
		guess := (smol + big) / 2
		if areTheCowsCalmNow(stalls, guess, cows) {
			smol = guess + 1
			preferredSpace = guess
		} else {
			big = guess - 1
		}
	}

	return preferredSpace
}
