package main

import "fmt"

func contains(intSlice []int, searchInt int) bool {
	for _, value := range intSlice {
		if value == searchInt {
			return true
		}
	}
	return false
}

func PermMissingElem_solution (A []int) int {

	var result int = -1;
	var indexmap map[int]bool = make(map[int]bool);
	var length = len(A);
	var i int;
	for i=1; i<=length; i++ {
		indexmap[i] = false;
	}

	for _, num := range A {
		var _, ok = indexmap[num];
		if (ok) {
			indexmap[num] = true;
		}
	}

	for k, v := range indexmap {
		if (!v) {
			result = k;
			break;
		}
	}

	//which means, I found everything
	if (result == -1) {
		result = length + 1;
	}

	return result;
}

func main () {

	var test = []int{1,3,4,5,6};
	fmt.Println(PermMissingElem_solution(test))
}