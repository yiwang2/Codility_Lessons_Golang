package main

import (
	"fmt"
	"math"
	"sort"
)

type IntArray struct {
	mArr       [][]int
	firstIndex int
}

func (arr *IntArray) Len() int {
	return len(arr.mArr)
}

func (arr *IntArray) Swap(i, j int) {
	arr.mArr[i], arr.mArr[j] = arr.mArr[j], arr.mArr[i]
}

func (arr *IntArray) Less(i, j int) bool {
	arr1 := arr.mArr[i]
	arr2 := arr.mArr[j]

	for index := arr.firstIndex; index < len(arr1); index++ {
		if arr1[index] < arr2[index] {
			return true
		} else if arr1[index] > arr2[index] {
			return false
		}
	}

	return i < j
}

func getMinimumIndexOfNail(startPos int, endPos int, nails [][]int, prevResultIndex int) int {

	var ln int = len(nails);
	var left int = 0;
	var right int = ln - 1;
	for left <= right {
		var mid int = (right-left)/2 + left;
		if nails[mid][1] >= startPos {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	if (left == ln || nails[left][1] > endPos) {
		return ln;
	}

	var resultIndex int = nails[left][0]

	for nextPossiblePos := left + 1; nextPossiblePos < ln; nextPossiblePos++ {
		if nails[nextPossiblePos][1] > endPos {
			break;
		}

		resultIndex = int(math.Min(float64(resultIndex), float64(nails[nextPossiblePos][0])));
		if prevResultIndex >= resultIndex {
			return prevResultIndex
		}
	}
    return resultIndex;
}

//NailingPlanks  87%  running time: 0.10 sec., time limit: 0.10 sec.
func NailingPlanks_Solution(A []int, B []int, C []int) int {

	var lc int = len(C);
	var la int = len(A);
	var nails [][]int  = make([][]int, lc);
	for  i := 0; i < lc; i ++{
		nails[i] = make([]int, 2);
		nails[i][0] = i;
		nails[i][1] = C[i];
	}

	mIntArray := &IntArray{nails, 1}
	sort.Sort(mIntArray)
	nails = mIntArray.mArr

	var globalMinimumIndex int = 0

	for plankIndex:=0; plankIndex< la; plankIndex ++ {
		localMinimumIndex := getMinimumIndexOfNail(A[plankIndex], B[plankIndex], nails, globalMinimumIndex)
		globalMinimumIndex = int(math.Max(float64(globalMinimumIndex), float64(localMinimumIndex)))
		if globalMinimumIndex == lc {
			return -1
		}
	}

	return globalMinimumIndex +1;
}

func main() {
	var A []int = []int{1, 4, 5, 8};
	var B []int = []int{4, 5, 9, 10};
	var C []int = []int{4, 6, 7, 10, 2};
	fmt.Println(NailingPlanks_Solution(A, B, C));
}
