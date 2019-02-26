package main

import "fmt"

//Peaks
func Peaks_Solution(A []int) int {

	var l int = len(A);
	if (l < 3) {
		return 0;
	}

	var peaks []int;
	var result int = 0;
	for i:=1 ; i<=l-2; i ++ {
		if ((A[i]> A[i+1]) && (A[i]> A[i-1])) {
			peaks = append(peaks, i);
		}
	}
	if (len(peaks) == 0) {
		return 0;
	}

	for size:=1; size<=l; size++ {
		if (l%size != 0) {
			continue;
		}

		var groups int = l/size;
		var find int = 0;
		var ok bool = true;
		for _, key := range peaks {
			if (key/size > find) {
				ok = false;
				break;
			}
			if (key/size == find) {
				find ++;
			}
		}

		if (find != groups) {
			ok = false;
		}
		if (ok) {
			result = groups;
			break;
		}
	}
	return result;
}

func main () {
	var test []int = []int {0,1,0,0,1,0,0,1,0}; // 1,2,3,4,3,4,1,2,3,4,6,2
	fmt.Println(Peaks_Solution(test));
}