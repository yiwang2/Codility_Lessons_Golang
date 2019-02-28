package main

import "fmt"

func isPrime(value int) bool {
	if value <= 3 {
		return value >= 2
	}
	if value%2 == 0 || value%3 == 0 {
		return false
	}
	for i := 5; i*i <= value; i += 6 {
		if value%i == 0 || value%(i+2) == 0 {
			return false
		}
	}
	return true
}

//CountSemiprimes
func CountSemiprimes_Solution(N int, P []int, Q []int) []int {
	var result []int;
	var primes []int;
	var semiprimes map[int]int = make(map[int]int);
	for i := 2; i <= N; i ++ {
		if (isPrime(i)) {
			primes = append(primes, i);
		}
	}

	if (len(primes) == 0) {
		return []int{0};
	}

	for i := 0; i < len(primes); i ++ {
		for j := 0; j < len(primes) && primes[i] * primes[j]<=N; j ++ {
			semiprimes[primes[i]*primes[j]] = 1;
		}
	}

	var semiprimesCheck []int = make([]int, N+1);
	var semiprimesAgg []int = make([]int, N+1);
	for i:=1; i <=N ; i ++ {
		if _, has := semiprimes[i]; has {
			semiprimesCheck [i] = 1;
		} else {
			semiprimesCheck [i] = 0;
		}
	}

	for i:=1; i <=N ; i ++ {
		semiprimesAgg[i] = semiprimesCheck [i];
		semiprimesAgg[i] = semiprimesAgg[i] + semiprimesAgg[i-1];
	}

	var l int = len(P);
	for i := 0; i < l; i++ {
		var startNum int = P[i];
		var endNum int = Q[i];
		var count int = semiprimesAgg[endNum] - semiprimesAgg[startNum] + semiprimesCheck[startNum];
		result = append(result, count);
	}

	return result;
}

/*
func CountSemiprimes_Solution2(N int, P []int, Q []int) []int {
}
*/



func main() {
	var testP []int = []int{1, 4, 16};
	var testQ []int = []int{26, 10, 20};
	var N int = 26;
	fmt.Println(CountSemiprimes_Solution(N, testP, testQ));
}
