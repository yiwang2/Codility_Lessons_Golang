package main

func Distinct_Solution(A []int) int {

	var count int = 0;
	distinc_map := make(map[int]bool, len(A));

	for _, num :=range A {
		var _, has = distinc_map[num];
		if (!has) {
			distinc_map[num] = true;
			count ++;
		}
	}

	return count;
}
