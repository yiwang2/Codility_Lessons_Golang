package main

import "fmt"

func Flags_Solution(A []int) int {

	var l int = len(A);
	if (l <3) {
		return 0;
	}
	var peaks []int;

	for i:=1 ; i<=l-2; i ++ {
		if ((A[i]> A[i+1]) && (A[i]> A[i-1])) {
			peaks = append(peaks, i);
		}
	}
	var peakCounts int = len(peaks);
	if (peakCounts == 0 || peakCounts == 1) {
		return peakCounts;
	}

	var lastCount int = 1;
	for i:=2; i<=peakCounts && (i-1)*(i-1) <= l; i++ {

		var flagsCount int = 1;
		var current int = peaks[0];
		for j:=1; j<peakCounts; j++ {
			if (flagsCount >= i) {
				break;
			}

			if (peaks[j] - current >= i) {
				flagsCount ++;
				current = peaks[j];
			}
		}

		if (flagsCount < i) {
			lastCount = i-1;
			break;
		} else {
			lastCount = i;
		}
	}

	return lastCount;
}

func main () {
	var test []int = []int{0, 0, 0, 0, 0, 1, 0, 1, 0, 1};// // 1,5,3,4,3,4,1,2,3,4,6,2
	fmt.Println(Flags_Solution(test));
}
