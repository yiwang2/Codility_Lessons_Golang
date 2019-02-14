package main

import (
	"./stack"
	"fmt"
)

func StoneWall_Solution2(H []int) int {

	var result int = 0;
	var stack []int = []int{};
	for i := 0; i < len(H); i++ {
		for len(stack) > 0 && stack[len(stack)-1] > H[i] {
			stack = stack[0:len(stack)-1];
		}
		if (len(stack) == 0 || stack[len(stack)-1] < H[i]) {
			stack = append(stack, H[i]);
			result ++;
		}

	}

	return result;
}

func StoneWall_Solution(H []int) int {
	var l *stack.Stack = stack.New();
	result := 0

	for _, value := range (H) {
		for l.Len() > 0 && l.Peek().(int) > value {
			l.Pop()
		}

		if l.Len() == 0 || l.Peek().(int) < value {
			l.Push(value)
			result += 1
		}
	}

	return result
}

func main() {
	var test []int = []int{8, 8, 5, 7, 9, 8, 7, 4, 8};
	fmt.Println(StoneWall_Solution2(test));
}
