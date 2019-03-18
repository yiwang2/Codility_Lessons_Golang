package main

import (
	"math"
	"fmt"
)

func MinMaxDivision_Solution(K int, M int, A []int) int {

	var max int = 0; //sum as max
	var min int = 0; // largest num as min

	for _, num :=range A {
		max = max + num;
		min = int(math.Max(float64(num), float64(min)));
	}
	var result int = max;
	for min <= max {
		var midNum int = (min+max)/2;
		if (divisionSolvable(midNum, K - 1, A)) { // group 0 - group k-1
			max = midNum - 1;
			result = midNum;
		} else {
			min = midNum + 1;
		}
	}


	return result;
}

//is this mid num divisible? into k+1 (0 - K-1) groups
func divisionSolvable (mindnum int, k int, A []int) bool{

	var sum int = 0;
	for _, num :=range A {
		sum = sum + num;
		if sum > mindnum {
			sum = num;
			k--;
		}

		if k < 0 {
			return false;
		}
	}

	return true;
}

func main () {
	var test[]int= []int{2,1,5,1,2,2,2};
	fmt.Println(MinMaxDivision_Solution(3, 5, test));
}