package main

import (
	"fmt"
	"math"
)

func main() {
	pages := []int{25, 46, 28, 49, 24}
	fmt.Println(keepPagesAllocationToMinimum(pages, 8))
	fmt.Println(fairLoad(pages, 4))
	pages = []int{12, 34, 67, 90}
	fmt.Println(keepPagesAllocationToMinimum(pages, 2))
	pages = []int{1, 17, 14, 9, 15, 9, 14}
	fmt.Println(keepPagesAllocationToMinimum(pages, 7))
}

/*
- Guess a upperbound.
- Keep allocating books to the student unless you meet that upperbound for that student
- If limit is reached, pick the next student and do the same
*/
func checkStudentCapacity(books []int, students int, upperbound int) bool {
	currentLoad := 0
	stuffedStudents := 1

	for _, pages := range books {
		currentLoad += pages
		if currentLoad > upperbound {
			// = pages, coz we're moving to next student
			stuffedStudents++
			currentLoad = pages
			if stuffedStudents > students {
				return false
			}
		}
	}
	return true
}

func keepPagesAllocationToMinimum(books []int, students int) int {

	if students > len(books) {
		return -1
	}

	maxFeasibleLoadOnStudent := 0
	maxLoad := 0

	for _, page := range books {
		maxFeasibleLoadOnStudent = max(maxFeasibleLoadOnStudent, page)
		maxLoad += page
	}

	preferredLoad := 0

	for maxFeasibleLoadOnStudent <= maxLoad {
		guess := (maxFeasibleLoadOnStudent + maxLoad) / 2
		if checkStudentCapacity(books, students, guess) {
			preferredLoad = guess
			maxLoad = guess - 1
		} else {
			maxFeasibleLoadOnStudent = guess + 1
		}
	}

	return preferredLoad
}

/*
BIGGEST FLAW: They ask of all books allocated to students, which student has the maximum books,

	While, keeping the individual allocation to each students the minimum.

	They don't ask for the minimum subset from all the pages.
*/
func fairLoad(pages []int, students int) int {
	maxBooksPerStudent := len(pages) - students + 1
	if maxBooksPerStudent <= 0 {
		return -1
	}

	leastFeasibleLoad := math.MaxInt
	currentLoad := 0

	//              3, 0 to 2
	for i := 0; i < maxBooksPerStudent; i++ {
		currentLoad += pages[i]
	}
	leastFeasibleLoad = currentLoad
	fmt.Println("window", maxBooksPerStudent, ":", leastFeasibleLoad)

	//       3
	for i := maxBooksPerStudent; i < len(pages); i++ {
		//                   3-3 = 0, moving window
		currentLoad -= pages[i-maxBooksPerStudent]
		currentLoad += pages[i]
		leastFeasibleLoad = min(leastFeasibleLoad, currentLoad)
	}
	return leastFeasibleLoad
}
