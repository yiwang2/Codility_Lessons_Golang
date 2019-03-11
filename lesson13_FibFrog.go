package main

import "fmt"

func FibFrog_Solution(A []int) int {

	//add start and end point
	A = append(A, 1);
	A = append([]int{1}, A...);

	var l int = len(A);
	var reachable []int = make([]int, l);
	var fab []int = make([]int, 26); // 26 num less than 100,000

	fab[0] = 0;
	fab[1] = 1;

	for i:=2; i<=26; i ++ {

		if ((fab[i-1]+fab[i-2]) > l-1) {
			break;
		}
		fab[i] = fab[i-1]+fab[i-2];
	}

	for _, num :=range fab {
		if A[num] == 1 {
			reachable[num] = 1;
		}
	}

	for i:=0; i < l; i ++{
		if A[i] == 0 || reachable[i] > 0 {
			continue;
		}

		var min_idx int = -1
		var min_value int = 100000

		for _, num :=range fab {
			previous_idx := i - num;
			if previous_idx < 0 {
				break;
			}

			if reachable[previous_idx] > 0 && min_value > reachable[previous_idx] {
				min_value = reachable[previous_idx]
				min_idx = previous_idx
			}
		}

		if min_idx != -1 {
			reachable[i] = min_value +1
		}

	}

	if (reachable[len(A)-1] == 0) {
		return -1;
	}

	return reachable[len(A)-1];
}

func main () {
	var test []int = []int{0,0,0}; //0,0,0,1,1,0,1,0,0,0,0
	fmt.Println(FibFrog_Solution(test));
}