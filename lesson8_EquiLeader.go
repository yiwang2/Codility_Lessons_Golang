package main

import "fmt"

//EquiLeader

func EquiLeader_Solution(A []int) int {

	result := 0;
	A_lenght := len(A);
	if (A_lenght == 0 || A_lenght == 1) {
		return result;
	}
	var target int = A_lenght/2+1;
	var leaders map[int]int = make(map[int]int);

	for _, num := range A {
		if _, has := leaders[num]; has {
			leaders[num] ++;
		} else {
			leaders[num] = 1;
		}
	}
	leader := -1;
	num_of_leader := 0;
	for key, val :=range leaders {
		if (val >= target) {
			leader = key;
			num_of_leader = val;
			break;
		}
	}

	count := 0;
	for index, num := range A {
		if (num == leader) {
			count ++;
		}
		var a int = num_of_leader - count;
		//left 0 - index

		if ((count >= ((index +1)/2 +1)) && (a>=(A_lenght-index -1)/2 +1)) {
			result ++;
		}
	}

	return result;
}

func main () {
	var test []int = []int {4, 4, 2, 5, 3, 4, 4, 4};
	fmt.Println(EquiLeader_Solution(test));
}