package main

import "fmt"

//MaxDoubleSliceSum
func MaxDoubleSliceSum_Solution(A []int) int {
	var l int = len(A);
	if (l <= 3) {
		return 0;
	}

	var range_1 []int = make([]int, l);
	var range_2 []int = make([]int, l);
	var max_result int = 0;
	for i := 1; i < l-1; i ++ {

		if (range_1[i-1] + A[i] < 0) {
			range_1[i] = 0;
		} else {
			range_1[i] = range_1[i-1] + A[i];
		}

	}

	for j:= l-2; j >=1; j-- {
		if (range_2[j+1] + A[j] < 0) {
			range_2[j] = 0;
		} else {
			range_2[j] = range_2[j+1] + A[j];
		}
	}

	for i := 1; i < l-1; i++ {

		if (max_result < range_1[i-1] + range_2[i+1]) {
			max_result = range_1[i-1] + range_2[i+1];
		}
	}

	return max_result;
}

func main () {
	var test []int = []int{3,2,6,-1,4,5,-1,2};
	//{3,2,6,-1,4,5,-1,2}
	//{0,2,8,7,11,16,15,15}
	//{16,16,14,8,9,5,0,0}
	fmt.Println(MaxDoubleSliceSum_Solution(test));
}
