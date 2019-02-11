package main
import (
	"sort"
	"fmt"
);

type circul_range struct {
	left int;
	right int;
}

type circul_range_list []circul_range;

func (members circul_range_list) Len() int           { return len(members) }
func (members circul_range_list) Swap(i, j int)      { members[i], members[j] = members[j], members[i] }
func (members circul_range_list) Less(i, j int) bool {
	if members[i].left < members[j].left {
		return true
	}
	if members[i].left > members[j].left {
		return false
	}
	return members[i].right < members[j].right;
}

func NumberOfDiscIntersections_Solution(A []int) int {

	var result int = 0;
	var max_allow int = 10000000;
	var m int = len(A);
	var range_map circul_range_list = make(circul_range_list, m);

	for index, num := range A{
		range_map[index] = circul_range{index - num, index + num};
	}

	sort.Sort(range_map);

	for index, values :=range range_map {
		//var left int = values.left;
		var right int = values.right;
		var another_index int;
		for another_index = index +1; another_index < len(range_map); another_index++{
			if (right <range_map[another_index].left) {
				break;
			}
			result ++;
			if (result > max_allow) {
				result = -1;
				break;
			}

		}

		if (result == -1) {
			break;
		}
	}

	return result;
}

func main () {

	var test []int = []int {1,5,2,1,4,0};

	fmt.Println(NumberOfDiscIntersections_Solution(test));
}