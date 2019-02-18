package main


func Dominator_Solution(A []int) int {

	var targetNum int = len(A)/2 +1 ;// at least
	var statNum map[int]int = make(map[int]int);
	var findLeader bool = false;
	var leader int = 0;
	for _, num := range A {
		if _, has := statNum[num]; has {
			statNum[num] ++;
		} else {
			statNum[num] = 1;
		}

		if (statNum[num] >= targetNum) {
			findLeader = true;
			leader = num;
			break;
		}
	}

	if (!findLeader) {
		return -1;
	}

	for index, num := range A {
		if (leader == num) {
			return index;
		}
	}

	return -1;
}