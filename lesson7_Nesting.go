package main

import "./stack"

func Nesting_Solution(S string) int {
	var sLength int  = len(S);
	var s *stack.Stack = stack.New();
	if (sLength == 0) {
		return 1;
	}

	for _, val :=range (S) {
		var val1 = string(val)
		if (val1 == "(") {
			s.Push(val1)
		} else if (val1 == ")") {
			if (s.Len() == 0) {
				return 0;
			}

			var value string = s.Pop().(string);
			if (value != "(") {
				return 0;
			}
		}
	}

	if (s.Len() != 0) {
		return 0;
	}

	return 1;
}
