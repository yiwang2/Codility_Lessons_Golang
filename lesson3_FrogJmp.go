package main

func Frog_Solution(X int, Y int, D int) int {

	var times int = 0;
	times = (Y-X)/D;
	if ((Y-X)%D) > 0 {
		times += 1;
	}

	return times;
}