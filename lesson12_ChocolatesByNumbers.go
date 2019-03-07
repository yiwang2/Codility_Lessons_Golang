package main

import "fmt"

func ChocolatesByNumbers_Solution(N int, M int) int {

	var counts map[int]int = make(map[int]int);
	var result int = 0;
	for i := 0 ;; i ++{
		var index int;
		if (M*i <= N-1) {
			index = M*i
		} else {
			index = (M*i)%N;
		}

		if (counts[index ] == 1) {
			break;
		} else {
			counts[index] = 1
			result++
		}
	}
	return result;
}

func ChocolatesByNumbers_Solution2(N int, M int) int {
	return N/gcd(N, M)
}

func gcd (a int, b int) int {
	if(a % b == 0)  {
		return b;
	}

	return gcd(b,a%b);
}

func main () {
	fmt.Println(ChocolatesByNumbers_Solution2(10,4));
}
