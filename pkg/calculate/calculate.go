package calculate

import (
	"fmt"
	"strconv"
)

type StackElement struct {
	num  float64
	sign byte
}

type Stack struct {
	items []StackElement
}

func (s *Stack) Push(data StackElement) {
	s.items = append(s.items, data)
}

func (s *Stack) IsEmpty() bool {
	return len(s.items) == 0
}

// func (s *Stack) Pop() (StackElement, error) {
// 	if s.IsEmpty() {
// 		return StackElement{}, fmt.Errorf("stack empty")
// 	}
// 	s.items = s.items[:len(s.items)-1]
// 	return s.items[len(s.items)-1], nil
// }

func (s *Stack) Pop() StackElement {
	if s.IsEmpty() {
		return StackElement{}
	}
	result := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return result
}

func (s *Stack) Top() StackElement {
	if s.IsEmpty() {
		return StackElement{}
	}
	return s.items[len(s.items)-1]
}

func Dict(operation byte, a float64, b float64) (float64, error) {
	switch operation {
	case '+':
		return a + b, nil
	case '-':
		return a - b, nil
	case '*':
		return a * b, nil
	case '/':
		if b == 0 {
			return 0, fmt.Errorf("division by zero")
		}
		return a / b, nil
	default:
		return 0, fmt.Errorf("invalid operation")
	}
}

func Priority(operation byte) (int, error) {
	switch operation {
	case '(':
		return 1, nil
	case '+', '-', ')':
		return 2, nil
	case '*', '/':
		return 3, nil
	default:
		return 0, fmt.Errorf("bad operation sign")
	}
}

func Brackets(s string) error {
	bal := 0
	err := fmt.Errorf("invalid brackets")
	for _, i := range s {
		if i == '(' {
			bal++
		} else if i == ')' {
			if bal <= 0 {
				return err
			}
			bal--
		}
	}
	if bal != 0 {
		return err
	}
	return nil
}

func apply(stack *Stack, operation byte) error {
	if len(stack.items) < 2 {
		return fmt.Errorf("not enough operands")
	}
	b := stack.Pop().num
	a := stack.Pop().num
	c, err := Dict(operation, a, b)
	if err != nil {
		return err
	}
	stack.Push(StackElement{num: c})
	return nil
}

func Calculate(expression string) (float64, error) {
	var operands Stack
	var operations Stack
	s := ""

	for _, i := range expression {
		if i != ' ' {
			s += string(i)
		}
	}

	err := Brackets(s)
	if err != nil {
		return 0, err
	}

	for i := 0; i < len(s); i++ {
		if '0' <= s[i] && s[i] <= '9' {
			num := ""
			for '0' <= s[i] && s[i] <= '9' {
				num += string(s[i])
				i++
				if i == len(s) {
					break
				}
			}
			i--
			number, err := strconv.ParseFloat(num, 64)
			if err != nil {
				return 0, err
			}
			operands.Push(StackElement{num: number})
		} else if s[i] == '(' {
			operations.Push(StackElement{sign: '('})
		} else if s[i] == ')' {
			for operations.Top().sign != '(' {
				err := apply(&operands, operations.Pop().sign)
				if err != nil {
					return 0, err
				}
			}
			operations.Pop()
		} else if s[i] == '-' && (i == 0 || s[i-1] == '(') {
			operands.Push(StackElement{num: -1})
			operations.Push(StackElement{sign: '*'})
		} else {
			prior, err := Priority(s[i])
			if err != nil {
				return 0, err
			}
			for !operations.IsEmpty() {
				current_prior, err := Priority(operations.Top().sign)
				if err != nil {
					return 0, err
				}
				if current_prior < prior {
					break
				}
				err1 := apply(&operands, operations.Pop().sign)
				if err1 != nil {
					return 0, err1
				}
			}
			operations.Push(StackElement{sign: s[i]})
		}
	}
	for !operations.IsEmpty() {
		err := apply(&operands, operations.Pop().sign)
		if err != nil {
			return 0, err
		}
	}
	if operands.IsEmpty() || len(operands.items) > 1 {
		return 0, fmt.Errorf("invalid operands number")
	}
	return operands.Top().num, nil
}
