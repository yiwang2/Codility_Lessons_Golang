package main

import "sort";

//MaxProductOfThree
func MaxProductOfThree_Solution(A []int) int {

	sort.Ints(A);
	var m = len(A);
	var a int = A[0] * A[1] * A[m-1];
	var b int = A[m-1] * A[m-2]* A [m-3];

	if a >= b {
		return a;
	} else {
		return b;
	}
}

func main () {
	var test []int = []int {-3, 1, 2, -2, 5, 6};
	MaxProductOfThree_Solution(test);
}