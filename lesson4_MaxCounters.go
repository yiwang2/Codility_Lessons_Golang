package main

func maxcounters_Solution1(N int, A []int) []int {

	var result []int = make([]int, N);
	var maxcount int = 0;
	for _, X :=range A {
		if (X >= 1 && X <= N) {
			result[X-1]++;
			if (result[X-1] > maxcount) {
				maxcount = result[X-1];
			}
		}else if (X == (N+1)) {
			for index, _ :=range result {
				result[index] = maxcount;
			}
		}
	}

	return result;
}

func maxcounters_Solution2(N int, A []int) []int {
	var result []int = make([]int, N);
	var maxcount int = 0;
	var lastincrease int = 0;
	for _, X :=range A {
		if (X >= 1 && X <= N) {
			if (result[X-1] < lastincrease) {
				result[X-1] = lastincrease;
			}
			result[X-1]++;
			if (result[X-1] > maxcount) {
				maxcount = result[X-1];
			}
		}else if (X == (N+1)) {
			lastincrease = maxcount;
		}
	}

	for index, _ :=range result {
		if (result[index] < lastincrease) {
			result[index] = lastincrease;
		}
	}

	return result;
}