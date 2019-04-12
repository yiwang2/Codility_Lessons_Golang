package main

import (
	"math"
	"fmt"
)

//AbsDistinct

func AbsDistinct_Solution(A []int) int {

	var la int = len(A);
	var head int = 0;
	var tail int = la-1;
	var result int = 1;
	var currentMax int = int(math.Max(math.Abs(float64(A[head])),math.Abs(float64(A[tail]))));

	for head <= tail {
		var currentHead int = int(math.Abs(float64(A[head])));
		if (currentHead == currentMax) {
			head++;
			continue;
		}

		var currentTail int = int(math.Abs(float64(A[tail])));
		if (currentTail == currentMax) {
			tail--;
			continue;
		}

		// get the new current maximal value
		if (currentHead >= currentTail) {
			currentMax = currentHead;
			head++;
		} else {
			currentMax = currentTail;
			tail--;
		}
		// meet a new distinct absolute value
		result++;
	}

	return result
}

func main () {
	var test []int = []int{-5,-3,-1,-0,3,6};
	fmt.Println(AbsDistinct_Solution(test));
}