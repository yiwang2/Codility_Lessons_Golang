package main


func CountDiv_Solution1(A int, B int, K int) int {

	count := 0;
	for i := A; i <=B; i++ {
		if ( i / K >= 0 && i % K == 0) {
			count ++;
		}
	}

	return count;
}

func CountDiv_Solution2(A int, B int, K int) int {

	count := 0;
	if (A / K >=0 && A % K == 0) {
		count ++;
	}

	count = count + B/K - A/K;

	return count;
}