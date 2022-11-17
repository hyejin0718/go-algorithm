package main

import "fmt"

func main() {
	test := []string{
		"([)]{}", "()()[\"\"]{}[]{}",
		"[([{}])]''{([''])}", "'\"\"'",
		"'(\")\"'", "''\"\"''", "'\"'\"",
		"}}{{", "{}())", "[['()'", "{[()]}",
	}

	for i, t := range test {
		fmt.Println("#", i, ": ", t, checkBracketsAreValid(t))
	}
}

func checkBracketsAreValid(str string) bool {
	var stack Stack
	stack.init(len(str))

	sqm := 0 // 작은 따옴표 개수
	dqm := 0 // 큰 따옴표 개수

	for _, s := range str {
		switch s {
		case '(':
			stack.push(s)
		case '{':
			stack.push(s)
		case '[':
			stack.push(s)
		case ')':
			value := stack.pop()
			if value == nil || *value != '(' {
				return false
			}
		case '}':
			value := stack.pop()
			if value == nil || *value != '{' {
				return false
			}
		case ']':
			value := stack.pop()
			if value == nil || *value != '[' {
				return false
			}
		case '\'':
			sqm = sqm + 1
			if sqm%2 == 1 {
				stack.push(s)
			} else {
				value := stack.pop()
				if value == nil || *value != '\'' {
					return false
				}
			}
		case '"':
			dqm = dqm + 1
			if dqm%2 == 1 {
				stack.push(s)
			} else {
				value := stack.pop()
				if value == nil || *value != '"' {
					return false
				}
			}
		}
	}
	if stack.isEmpty() {
		return true
	}
	return false
}

/* Stack */

type Stack struct {
	size, top int
	array     []interface{}
}

func (s *Stack) init(size int) {
	s.size = size
	s.top = 0
	s.array = []interface{}{}
}

func (s *Stack) push(value interface{}) {
	if s.isOverflow() {
		return
	}
	s.array = append(s.array, value)
	s.top = s.top + 1
}

func (s *Stack) pop() *interface{} {
	if s.isEmpty() {
		return nil
	}
	s.top = s.top - 1
	result := s.array[s.top]
	s.array = s.array[:s.top]
	return &result
}

func (s *Stack) isEmpty() bool {
	if s.top <= 0 {
		return true
	}
	return false
}

func (s *Stack) isOverflow() bool {
	if len(s.array) >= s.size {
		return true
	}
	return false
}

func (s *Stack) showStack() {
	fmt.Println("stack: ", s)
}
