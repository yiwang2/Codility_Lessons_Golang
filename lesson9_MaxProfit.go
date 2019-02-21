package main

import "fmt"

//MaxProfit
func MaxProfit_Solution(A []int) int {

	var mx_p int = 0;

	if (len(A) == 0) {
		return mx_p;
	}

	var currentMinValue int = A[0];
	var currentMaxValue int = A[0];

	for _, num := range A  {

		if ((num-currentMinValue) > mx_p) {
			mx_p = (num-currentMinValue);
		}

		if (currentMinValue > num) {
			currentMinValue = num;
		}

		if (num > currentMaxValue) {
			currentMaxValue = num;
		}
	}

	return mx_p;
}

func main () {
	var test []int = []int{23171, 21011, 21123, 21366, 21013, 21367};
	fmt.Println(MaxProfit_Solution(test));
}