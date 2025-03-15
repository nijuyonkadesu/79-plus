package main

import (
	"fmt"
	"math"
)

func main() {
	nums := []int{3, 6, 7, 11}
	fmt.Println(yourSpeedOnEachSprint(nums, 8))
	nums = []int{30, 11, 23, 4, 20}
	fmt.Println(yourSpeedOnEachSprint(nums, 6))
}

/*
k = speed = tickers / hr
**instead of computing remainder, and adding it to the quotient, we here perform Ceil operation**
okay, now: 

total time spent on stories should be less than or equal the given target time.
so, we need to compute how much it takes for different speed while still not crossing target time. 
so, we'll find the max possible speed, and reduce it one by one to find a sweet spot. 

total time spent = maxSpeed / (preferred speed)
(write the units on paper to understand this)

now, we'll check total time spent against the target time. 
*/
func yourSpeedOnEachSprint(stories []int, time int) int {
	biggestStory := 0
	maxSpeed := 0

	for _, story := range stories {
		biggestStory = max(biggestStory, story)
	}

	maxSpeed = biggestStory / 1 // x tickets / 1hr
	preferredSpeed := maxSpeed

	// reducing one by one will be too slow, this is where binary search helps?
	// for totalHours(stories, preferredSpeed) <= time {
	// 	preferredSpeed--
	// }

	// OHHHHHH!!! WE'RE DOING BINARY SEARCH TO FIND A CANDIDATE !!
	// WOAHHHH... THAT'S THE TRUE PURPOSE OF BINARY SEARCH HERE... 
	preferredSpeed = 1 
	for preferredSpeed < maxSpeed {
		candidate := (preferredSpeed + maxSpeed) / 2
		if totalHours(stories, candidate) <= time {
			maxSpeed = candidate
		} else {
			preferredSpeed = candidate + 1
		}
	}

	return preferredSpeed
}

func totalHours(stories []int, speed int) int {
	totalHours := 0.0

	for _, story := range stories {
		val := math.Ceil(float64(story) / float64(speed))
		totalHours += val
	}
	return int(totalHours)
}
/*
8 hrs is target time to complete all tickets in all story

		 n: total tickets
    len(n): total stories

       avg: k tickets / story
avg / time: k tickets / story, hour
avg/time x len(n): k tickets / hour

*/
func yourSpeedOnEachSprintCompleicatedDoesntWork(nums []int, time int) int {
	size := len(nums)
	fmt.Println("---")
	avg := 0.0
	biggest := 0
	total := 0

	for _, num := range nums {
		val := float64(num) / float64(time)
		biggest = max(biggest, num)
		total += num
		avg += val
	}
	fmt.Println("avg", math.Ceil(avg))
	ceilAvg := int(math.Ceil(avg))

	diff := biggest - ceilAvg
	fmt.Println("diff", diff)
	ratio := float64(ceilAvg) / float64(biggest)
	fmt.Println("r:", ratio)
	adjustment := int(math.Ceil((ratio * float64(biggest)))) * size
	fmt.Println("---")

	return adjustment + ceilAvg
}
