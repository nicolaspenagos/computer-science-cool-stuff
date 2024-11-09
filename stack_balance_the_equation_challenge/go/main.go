package main

import "fmt"

type Stack struct {
	items []rune
}

func (s *Stack) push(char rune) {
	s.items = append(s.items, char)
}

func (s *Stack) pop() (rune, error) {
	if !s.isEmpty() {
		lastItem := s.items[len(s.items)-1]
		s.items = s.items[:len(s.items)-1]
		return lastItem, nil
	}
	return 0, fmt.Errorf("Stack underflow")
}

func (s *Stack) isEmpty() bool {
	return len(s.items) == 0
}

func (s *Stack) peek() (rune, error) {
	if !s.isEmpty() {
		return s.items[len(s.items)-1], nil
	}
	return 0, fmt.Errorf("Stack underflow")

}

func isBalance(str string) bool {
	stack := &Stack{}

	for _, char := range str {
		if char == '{' || char == '[' || char == '(' {
			stack.push(char)
		} else if char == '}' || char == ']' || char == ')' {
			peek, err := stack.peek()
			if err != nil {
				return false
			}

			if char == ')' && peek != '(' {
				return false
			}

			if char == ']' && peek != '[' {
				return false
			}

			if char == '}' && peek != '{' {
				return false
			}

			stack.pop()
		}
	}

	return stack.isEmpty()
}

func main() {
	str := "{(a+b)*[a-(c+d)]}+{((a*a)+b)}"
	fmt.Printf("%s is balanced? %v\n", str, isBalance(str))

	str = "{(a+b)*[a-(c+d)]}]+{((a*a)+b)}"
	fmt.Printf("%s is balanced? %v\n", str, isBalance(str))
}
