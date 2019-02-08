package main

import "fmt"

func MinAvgTwoSlice_Solution(A []int) int {

	var min_value float64 = 10000;
	var min_index int = 0;
	var a_l int = len(A);
	for index, _ :=range A {
		// 0 - al-1
		if (index < a_l -1) {
			var average float64 = (float64)(A[index] + A[index+1])/2.0;
			if (average < min_value) {
				min_value = average;
				min_index = index;
			}
		}

		if (index < a_l -2) {
			var average float64 = (float64)(A[index] + A[index+1]+A[index + 2])/3.0;
			if (average < min_value) {
				min_value = average;
				min_index = index;
			}
		}

	}

	return min_index;
}

func main () {
	var test []int = []int{4, 2, 2, 5, 1 ,5 ,8};
	fmt.Println(MinAvgTwoSlice_Solution(test));
}