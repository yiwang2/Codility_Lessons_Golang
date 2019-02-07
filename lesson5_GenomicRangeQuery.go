package main

import "sync"
import (
	"strings"
	"fmt"
)

func GenomicRangeQuery_Solution1(S string, P []int, Q []int) []int {

	var waitG sync.WaitGroup;
	var lock sync.Mutex;
	var nucleotides_type []string = []string{"A", "C", "G", "T"};
	var nucleotides_impact_factor []int = []int{1, 2, 3, 4};
	var m = len(P);
	var result []int = make([]int, m);
	//var temp_result []string = make ([]string, m);
	var index int;
	for index = 0; index < m; index ++ {
		var p int = P[index];
		var q int = Q[index];
		//temp_result[index] = S[p:q+1];
		waitG.Add(1)
		go func(a int, b string) {
			defer waitG.Done();
			lock.Lock();
			if (strings.Contains(b, nucleotides_type[0])) {
				result[a] = nucleotides_impact_factor[0];
			} else if (strings.Contains(b, nucleotides_type[1])) {
				result[a] = nucleotides_impact_factor[1];
			} else if (strings.Contains(b, nucleotides_type[2])) {
				result[a] = nucleotides_impact_factor[2];
			} else if (strings.Contains(b, nucleotides_type[3])) {
				result[a] = nucleotides_impact_factor[3];
			}
			lock.Unlock();
		}(index, S[p:q+1]);
	}

	waitG.Wait();
	return result;
}

func GenomicRangeQuery_Solution2(S string, P []int, Q []int) []int {
	var nucleotides_type []string = []string{"A", "C", "G", "T"};
	var nucleotides_impact_factor []int = []int{1, 2, 3, 4};
	var m = len(P);
	var result []int = make([]int, m);

	var index int;
	for index = 0; index < m; index ++ {
		var p int = P[index];
		var q int = Q[index];
		var b string = S[p:q+1];
		if (strings.Contains(b, nucleotides_type[0])) {
			result[index] = nucleotides_impact_factor[0];
		} else if (strings.Contains(b, nucleotides_type[1])) {
			result[index] = nucleotides_impact_factor[1];
		} else if (strings.Contains(b, nucleotides_type[2])) {
			result[index] = nucleotides_impact_factor[2];
		} else if (strings.Contains(b, nucleotides_type[3])) {
			result[index] = nucleotides_impact_factor[3];
		}
	}

	return result;
}


func GenomicRangeQuery_Solution3(S string, P []int, Q []int) []int {
	var nucleotides_type []string = []string{"A", "C", "G", "T"};
	var m = len(P);
	var result []int = make([]int, m);
	//var xx int = len(S)+1;
	genoms   := make([][]int, 3);
	for  i := 0; i < 3; i ++{
		genoms[i] = make([]int, len(S)+1);
	}
	for  i := 0; i < len(S); i ++{
		var a, b, c = 0, 0 ,0;
		if string(S[i]) == nucleotides_type[0] {
			a = 1;
		} else if string(S[i]) == nucleotides_type[1] {
			b = 1;
		} else if string(S[i]) == nucleotides_type[2] {
			c = 1;
		}
		genoms[0][i+1] = genoms[0][i] + a;
		genoms[1][i+1] = genoms[1][i] + b;
		genoms[2][i+1] = genoms[2][i] + c;
	}

	for index := 0; index < m; index ++ {
		var fromIndex int = P[index];
		var toIndex int = Q[index]+1;
		if (genoms[0][toIndex] - genoms[0][fromIndex] > 0) {
			result[index] = 1;
		} else if (genoms[1][toIndex] - genoms[1][fromIndex] > 0) {
			result[index] = 2;
		} else if (genoms[2][toIndex] - genoms[2][fromIndex] > 0) {
			result[index] = 3;
		} else {
			result[index] = 4;
		}
	}

	return result;
}

func main () { // ('CAGCCTA', [2, 5, 0], [4, 5, 6])

	fmt.Println(GenomicRangeQuery_Solution3("GT", []int{0,0,1},[]int{0,1,1}));

}





