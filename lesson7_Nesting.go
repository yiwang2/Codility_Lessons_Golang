package main

import "./stack"

func Nesting_Solution(S string) int {
	var runeS []rune = []rune(S);
	var sLength int  = len(S);
	var s *stack.Stack = stack.New();
	if (sLength == 0) {
		return 1;
	}

	for i :=0; i < sLength; i ++ {
		var val = string(runeS[i]);
		if (val == "(") {
			s.Push(1)
		} else if (val == ")") {
			if (s.Len() == 0) {
				return 0;
			}

			var value int = s.Pop().(int);
			if (value -1 != 0) {
				return 0;
			}
		}
	}

	if (s.Len() != 0) {
		return 0;
	}

	return 1;
}
