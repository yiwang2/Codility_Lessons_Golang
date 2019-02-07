package main

import "fmt"

//permutation

func permutation_Solution(A []int) int {
	// write your code in Go 1.4
	var length int = len(A);
	var countmap map[int]int = make(map[int]int);

	for _, num := range A {
		if (num > length) {
			return 0;
		}
		var _, has = countmap[num];
		if (has) {
			return 0;
		} else {
			countmap[num] = 1;
		}
	}

	return 1;
}

func main () {

	var test = []int{1,2,4,5,2};
	fmt.Println(permutation_Solution(test))
}