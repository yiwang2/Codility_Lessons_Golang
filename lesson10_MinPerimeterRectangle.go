package main

import (
	"math"
	"fmt"
)

func MinPerimeterRectangle_Solution(N int) int {

	var perimeter int = 0;

	var start_value int = int(math.Ceil(math.Sqrt(float64(N))));

	for j:= start_value; j<=N; j ++ {
		for i:=start_value; i>=1; i -- {
			if (i*j < N) {
				break;
			}
			if (i*j == N) {
				return 2 * (i+j);
			}
		}
	}



	return perimeter;
}

func MinPerimeterRectangle_Solution2(N int) int {
	var result int = math.MaxInt32;
	for i:=1; i*i<=N; i ++ {
		if (N%i == 0) {
			if (result > 2*(N/i + i)) {
				result = 2*(N/i + i);
			}
		}
	}

	return result;
}

func main () {
	fmt.Println(MinPerimeterRectangle_Solution2(1234));
}