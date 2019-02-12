package main

import "fmt"

//go get github.com/golang-collections/collections
type (
	Stack struct {
		top *node
		length int
	}
	node struct {
		value interface{}
		prev *node
	}
)
// Create a new stack
func New() *Stack {
	return &Stack{nil,0}
}
// Return the number of items in the stack
func (this *Stack) Len() int {
	return this.length
}
// View the top item on the stack
func (this *Stack) Peek() interface{} {
	if this.length == 0 {
		return nil
	}
	return this.top.value
}
// Pop the top item of the stack and return it
func (this *Stack) Pop() interface{} {
	if this.length == 0 {
		return nil
	}

	n := this.top
	this.top = n.prev
	this.length--
	return n.value
}
// Push a value onto the top of the stack
func (this *Stack) Push(value interface{}) {
	n := &node{value,this.top}
	this.top = n
	this.length++
}


func Brackets_Solution(S string) int {

	var sLength int  = len(S);
	if (sLength % 2 != 0) {
		return 0;
	}

	var s *Stack = New();

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