package main

import "math"

//TapeEquilibrium
func TapeEquilibrium_Solution(A []int) int {
	var length int = len(A);
	var index int;
	var result int = 0;
	for index = 1; index < length; index ++ {
		var piece1 = A[0:index];
		var piece2 = A[index:];
		var piece1_sum int;
		var piece2_sum int;
		for _, num := range piece1 {
			piece1_sum = piece1_sum + num;
		}
		for _, num := range piece2 {
			piece2_sum = piece2_sum + num;
		}

		if (index == 1) {
			result = int(math.Abs(float64(piece1_sum - piece2_sum)));
		} else {
			result = int(math.Min(float64(result), math.Abs(float64(piece1_sum-piece2_sum))));
		}
	}

	return result;
}

func TapeEquilibrium_Solution2(A []int) int {
	var length int = len(A);
	var index int;
	var result int;
	for index = 1; index < length; index ++ {
		var value int;
		for i, num := range A {
			if (i <= index-1) {
				value = value + num;
			} else {
				value = value - num;
			}
		}

		if (index == 1) {
			result = int(math.Abs(float64(value)));
		} else {
			result = int(math.Min(float64(result), math.Abs(float64(value))));
		}
	}
	return result;
}

func TapeEquilibrium_Solution3 (A []int) int {

	var right int;
	var left int;
	var length int = len(A);
	for _, num := range A {
		right = right + num;
	}
	var index int;
	// 0, p-1   p, lenght-1
	var result int;
	for index = 1; index < length; index ++ {
		right = right - A[index-1];
		left = left + A[index-1];
		if (index == 1) {
			result = int(math.Abs(float64(right-left)));
		} else {
			result = int(math.Min(float64(result), math.Abs(float64(right-left))));
		}
	}

	return result;
}
