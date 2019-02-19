package main

import (
	"fmt"
	"math"
)

//MaxSliceSum 92%
func MaxSliceSum_Solution(A []int) int {
	var result int = 0;
	var countMap map[int]int = make(map[int]int);
	var sum_max int = math.MinInt32;
	var max_index int = 0;

	for index, num := range A {
		if(index == 0) {
			countMap[index] = num;
		} else {
			countMap[index] = countMap[index-1]+num;
		}

		if (countMap[index] > sum_max) {
			sum_max = countMap[index];
			max_index = index;
		}
	}

	result = A[max_index];
	var temp int = result;
	for i:=max_index-1; i >=0; i-- {
		temp = temp + A[i];
		if (result < temp) {
			result = temp;
		}
	}

	if (result < sum_max) {
		result = sum_max;
	}

	if (result < 0) {
		for _, num := range A {
			if (result < num) {
				result = num;
			}
		}
	}

	return result;
}


func MaxSliceSum_Solution2(A []int) int {

	var maxEnd int = A[0];
	var maxSoFar int = A[0];

	for index, num :=range A {
		if (index == 0) {
			continue;
		}
		maxEnd =int(math.Max(float64(num), float64(num+maxEnd)));
		maxSoFar = int(math.Max(float64(maxEnd), float64(maxSoFar)));
	}

	return maxSoFar;
}

func MaxSliceSum_Solution3(A []int) int {
	var sum_max_range int = A[0];
	var sum_max int = A[0];
	var temp int = A[0];

	for i:=1 ; i < len(A); i ++ {
		temp = temp + A[i];
		sum_max_range = int(math.Max(float64(temp),float64(A[i])));
		temp = sum_max_range;
		if (sum_max < sum_max_range) {
			sum_max = sum_max_range;
		}
	}

	return sum_max;
}

func main () {
	var test []int = []int{3,2,-6,4,0}// {3,2,-6,4,0};{3,-2,3}
	fmt.Println(MaxSliceSum_Solution3(test));
}