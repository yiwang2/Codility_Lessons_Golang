package main


import "sort"

//Triangle
func Triangle_Solution(A []int) int {

	sort.Ints(A);
	 var m int = len(A);
	 for index, _ := range A {
	 	if (index <= m-3) {
	 		if ((A[index] + A[index+1] > A[index + 2]) &&
					(A[index+2] + A[index+1]> A[index]) &&
					(A[index] + A[index+2] > A[index + 1])) {
						return 1;
			}
		}

	 }

	return 0;
}