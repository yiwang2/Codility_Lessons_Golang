package main

import "fmt"

//CountDistinctSlices

//bad performance
func CountDistinctSlices_Solution(M int, A []int) int {
	var result int = 0;
	var la = len(A);
	var count []int = make([]int, M+1);

	var start int = 0;
	var end int = 0;

	for i:=start; i < la; i ++{
		count = make([]int, M+1);
		end = i;
		for j:= end; j < la; j ++{
			if (count[A[j]] == 0) {
				result ++;
				count[A[j]] ++;
			} else {
				start ++;
				break;
			}
		}

		if (result > 1000000000) {
			result = 1000000000;
			break;
		}
	}

	return result;
}

/*def solution(M, A):
    the_sum = 0
    front = back = 0
    seen = [False] * (M+1)
    while (front < len(A) and back < len(A)):
        while (front < len(A) and seen[A[front]] != True):
            the_sum += (front-back+1)
            seen[A[front]] = True
            front += 1
        else:
            while front < len(A) and back < len(A) and A[back] != A[front]:
                seen[A[back]] = False
                back += 1

            seen[A[back]] = False
            back += 1


    return min(the_sum, 1000000000)
*/


func CountDistinctSlices_Solution1(M int, A []int) int {
	var result int = 0;
	var la = len(A);
	var count []int = make([]int, M+1);
	var start int = 0;
	var end int = 0;

	for start < la && end < la{
		if count[A[end]] == 0{
			result = result + (end-start+1);
			if (result > 1000000000) {
				return 1000000000;
			}
			count[A[end]] ++;
			end ++;
		}else {
			count[A[start]] = 0;
			start++;
		}
	}

	return result;
}


func main () {
	var test []int =[]int{3,4,5,5,2};
	fmt.Println(CountDistinctSlices_Solution1(6, test));
}