package main

import "fmt"
import  "../stack"

func Brackets_Solution(S string) int {

	var sLength int  = len(S);
	if (sLength % 2 != 0) {
		return 0;
	}
	var s *stack.Stack = stack.New();
	var runeS []rune = []rune(S);

	for i :=0; i < sLength; i ++ {
		if string(runeS[i]) == "{" {
			s.Push(1);
		} else if string(runeS[i]) == "[" {
			s.Push(2);
		} else if string(runeS[i]) == "(" {
			s.Push(3);
		} else if string(runeS[i]) == "}" {
			val :=s.Pop()
			if (val == nil || val.(int) -1 != 0) {
				return 0;
			}
		} else if string(runeS[i]) == "]" {
			val :=s.Pop()
			if (val == nil ||val.(int) -2 != 0) {
				return 0;
			}
		} else if string(runeS[i]) == ")" {
			val :=s.Pop()
			if (val == nil || val.(int) -3 != 0) {
				return 0;
			}
		}

	}

	if (s.Len() > 0) {
		return 0;
	}


	return 1;
}

func main () {

	var test string = "{{{{{{{"; //{[()()]}
	fmt.Println(Brackets_Solution(test));
}