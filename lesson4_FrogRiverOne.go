package main

func FrogRiverOne_solution (X int, A []int) int {

	var jumpmap map[int]bool = make(map[int]bool);
	var i int;
	for i=1; i<=X; i++ {
		jumpmap[i] =false;
	}

	var count int = 0;
	for index, num := range A{
		var _, has = jumpmap[num];
		if (!has) {
			continue;
		}

		if (jumpmap[num]) {
			continue;
		}
		jumpmap[num] = true;
		count ++;
		if (count == X) {
			return index;
		}
	}

	return -1;
}