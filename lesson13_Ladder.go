package main

import (
	"fmt"
)

func ladder_Solution(A []int, B []int) []int {
	var l int = len(A);
	var result []int = make([]int, l);
	var max int = 0;
	for _, num :=range A {
		if (max < num) {
			max = num;
		}
	}

	var fab []int = make([]int, max+2);

	fab[0] = 1;
	fab[1] = 1;

	for i:=2; i<=max+1; i ++ {
		fab[i] = (fab[i-1]+fab[i-2])%(1<<30);
	}


	for i:=0; i<l; i ++ {
		result[i] = fab[A[i]]%int(1<<uint (B[i]));
	}

	return result;
}

//if l == 50000
//it is why 1<<30 for max limitation
/*func test50000 () {
	var max int = 50000;
	var fab []int = make([]int, max+1);
	fab[0] = 1;
	fab[1] = 1;
	for i:=2; i<=max; i ++ {
		fab[i] = (fab[i-1]+fab[i-2]);
	}

	fmt.Println(fab)
	fmt.Println(1<<30)
}*/

func main () {
	var a []int = []int{4,4,5,5,1};
	var b []int = []int{3,2,4,3,1};

	fmt.Println(ladder_Solution(a,b))

	//test50000();
}