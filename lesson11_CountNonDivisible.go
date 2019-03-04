package main

import (
	"fmt"
)

func isdivisors(a int, b int) bool {
	return a%b == 0;
}

//performance 0 %
func CountNonDivisible_Solution(A []int) []int {

	var result []int;
	var l int = len(A);
	var tempMap []int = make([]int, l);
	for i := 0; i < l; i ++ {
		tempMap[i] = l - 1;
	}

	for i := 0; i < l; i ++ {
		for j := i + 1; j < l; j ++ {
			if isdivisors(A[i], A[j]) {
				tempMap[i] --;
			}

			if isdivisors(A[j], A[i]) {
				tempMap[j] --;
			}
		}
	}

	for _, val := range tempMap {
		result = append(result, val);
	}

	return result;
}

//performance 100%
func CountNonDivisible_Solution1(A []int) []int {

	var l int = len(A);
	var result []int = make([]int, l);
	var counts []int = make([]int, l*2+1)
	for _, num := range A {
		counts[num]++;
	}
	var divisorNums []int = make([]int, len(counts))
	for i := 1; i*i < len(divisorNums); i++ {
		for j := i * i; j < len(divisorNums); j += i {
			divisorNums[j] += counts[i];
			if (j != i*i) {
				divisorNums[j] += counts[j/i];
			}
		}
	}

	for i := 0; i < l; i++ {
		result[i] = l - divisorNums[A[i]];
	}

	return result;
}

//performance 50%
func CountNonDivisible_Solution2(A []int) []int {
	var l int = len(A);
	var result []int = make([]int, l);
	var counts []int = make([]int, l*2+1)
	for _, num := range A {
		counts[num]++;
	}

	for i := 0; i< l; i ++ {
		var divisors int = 0;
		for j:= 1; j*j <= A[i]; j ++ {
			if isdivisors(A[i], j) {
				divisors += counts[j]; // take duplication
				if A[i]/j != j {
					divisors += counts[A[i]/j]; // take duplication of A[i]/j
				}
			}

		}
		result[i] = l - divisors;
	}

	return result;
}

func main() {
	var test []int = []int{3, 1, 2, 3, 6};
	fmt.Println(CountNonDivisible_Solution2(test));
}
