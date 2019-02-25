package main

import "fmt"

//CountFactors

func CountFactors_Solution(N int) int {
	var count int = 0;
	for i:=1; i<= N; i++ {

		if (N%i == 0) {
			count ++;
		}
	}

	return count;
}

func CountFactors_Solution2(N int) int {
	var count int = 0;
	var i int = 1;
	for i*i <= N {

		if (N%i == 0) {
			if (i*i == N) {
				count ++;
			} else {
				count = count + 2;
			}
		}

		i++;
	}

	return count;
}

func main () {
	fmt.Println(CountFactors_Solution2(24));

}
