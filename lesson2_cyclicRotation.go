package main

import "fmt"

func CyclicRotation_Solution(A []int, K int) []int {

	var length int = len(A);
	var result []int = make ([]int, length);

	for index, num :=range A {
		var nextIndex int = (index + K)%length;
		result[nextIndex] = num;
	}

	return result;
}

func main () {

	var test []int = []int{3, 8, 9, 7, 6};
	var k int = 3;

	fmt.Println(CyclicRotation_Solution(test, k));

}