package main

import (
	"reflect"
	"fmt"
	"sort"
)

func isPrimeLesson12(value int) bool {
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

func hasCommonDivisor(a int, b int) bool {

	if (a == b) {
		return true;
	}

	if (a != b && isPrimeLesson12(a) && isPrimeLesson12(b)) {
		return false;
	}

	var ad []int;
	var bd []int;
	for i := 1; i*i <= a; i ++ {
		if (a%i == 0 && isPrimeLesson12(i)) {
			ad = append(ad, i);
			if (isPrimeLesson12(a/i) && a/i != i) {
				ad = append(ad, a/i);
			}
		}
	}

	if (isPrimeLesson12(a)) {
		ad = append(ad, a);
	}

	for i := 1; i*i <= b; i ++ {
		if (b%i == 0 && isPrimeLesson12(i)) {
			bd = append(bd, i);
			if (isPrimeLesson12(b/i) && b/i != i) {
				bd = append(ad, b/i);
			}
		}
	}

	if (isPrimeLesson12(b)) {
		bd = append(bd, b);
	}

	sort.Ints(ad);
	sort.Ints(bd);
	return reflect.DeepEqual(ad, bd);
}

func gcdlesson12(a int, b int) int {
	if (a%b == 0) {
		return b;
	}

	return gcdlesson12(b, a%b);
}

func hasCommonDivisor2(a int, b int) bool {

	if (a == b) {
		return true;
	}

	var commonDivisor int = gcdlesson12(a, b);
	var aa int = a;
	var bb int = b;
	var gcda int;
	var gcdb int;
	for aa != 1 {
		gcda = gcdlesson12(aa, commonDivisor);
		if gcda == 1 {
			break;
		}

		aa = aa/gcda;
	}

	if aa != 1 {
		return false;
	}

	for bb != 1 {
		gcdb = gcdlesson12(bb, commonDivisor);
		if gcdb == 1 {
			break;
		}

		bb = bb/gcdb;
	}

	return bb == 1;
}

func CommonPrimeDivisors_Solution(A []int, B []int) int {
	var l int = len(A);
	var count int = 0;
	for i := 0; i < l; i ++ {
		if (hasCommonDivisor2(A[i], B[i])) {
			count ++;
		}
	}

	return count;
}

func main() {
	var a []int = []int{3, 9};
	var b []int = []int{9, 81};
	fmt.Println(CommonPrimeDivisors_Solution(a, b))

}
