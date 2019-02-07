package main
//MissingInteger
func MissingInteger_Solution(A []int) int {

	var result int = 0;
	var countmap map[int]int = make(map[int]int);

	for _, num := range A {
		if (num <= 0) {
			continue;
		}
		_, has := countmap[num];
		if (!has) {
			countmap[num] = 1;
		} else {
			countmap[num] ++ ;
		}
	}
	var i int;
	for i =1; i <= len(countmap); i ++ {
		_, has := countmap[i];
		if (!has) {
			result = i;
			break;
		}
	}

	if (result == 0) {
		result = len(countmap) +1;
	}

	return result;
}