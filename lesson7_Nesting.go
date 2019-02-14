package main

import "./stack"

func Nesting_Solution2(S string) int {
	var sLength int  = len(S);
	var s []int= []int{};
	if (sLength == 0) {
		return 1;
	}

	for _, val :=range (S) {
		var val1 = string(val)
		if (val1 == "(") {
			s = append(s, 1);
		} else if (val1 == ")") {
			var ss int = len(s);
			if (ss == 0) {
				return 0;
			}

			if (s[ss-1] -1 == 0) {
				s= s[0:ss-1];
			} else {
				return 0;
			}
		}
	}

	if (len(s) != 0) {
		return 0;
	}

	return 1;
}

//87%   <= performance 75%  0.13 required 0.10
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
