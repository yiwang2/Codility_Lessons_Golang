package main

import "fmt";
import  "./stack"

func Fish_Solution2(A []int, B []int) int {

	var m int = len(A);
	var count int = m;
	var comingFish *stack.Stack = stack.New();
	for i := 0; i < m; i ++ {
		if (B[i] == 0) {
			if (comingFish.Len() == 0) {
				continue;
			} else {
				for comingFish.Len() != 0 {
					val :=comingFish.Peek();
					if (val.(int) > A[i]) {
						count --;
						break;
					} else if (val.(int) < A[i]) {
						count --;
						comingFish.Pop();
					}
				}
			}
		} else if (B[i] == 1) {
			comingFish.Push(A[i]);
		}
	}
	return count;
}


func main () {
	var A[]int = []int {5,1,1,1,1,1,1};
	var B[]int = []int {1,0,0,0,0,0,0};

	fmt.Println(Fish_Solution2(A, B));
}
