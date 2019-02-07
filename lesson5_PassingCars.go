package main

import "fmt"

func passingCar_Solution(A []int) int {

	var result int = 0;
	var pair_head int = 0;
	var multi_times int = 1;
	var head_assign bool = false;
	for _, num := range A {
		if (!head_assign && pair_head == num) {
			head_assign = true;
		} else {
			if (head_assign) {
				if (pair_head == num) {
					multi_times ++;
				} else {
					result = result + multi_times*1;
				}
			}
		}

		if (result > 1000000000) {
			result = -1;
			break;
		}
	}


	return result;

}

func main () {
	var test []int = []int {0, 1, 0, 1, 1};
	fmt.Println(passingCar_Solution(test));
}
