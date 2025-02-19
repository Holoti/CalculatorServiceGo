package calculate_v2

import (
	"errors"
	"fmt"
	"strconv"
)

var DecimalSeparator byte = '.' // TODO read this from .env

type StackElement struct {
	num  float64
	sign byte
}

type Stack struct {
	items []StackElement
}

func NewStack() *Stack {
	return &Stack{}
}

func (s *Stack) Push(data StackElement) {
	s.items = append(s.items, data)
}

func (s *Stack) IsEmpty() bool {
	return len(s.items) == 0
}

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

func Priority(operation byte) int {
	switch operation {
	case '(':
		return 1
	case '+', '-':
		return 2
	case '*', '/':
		return 3
	case '^':
		return 4
	default:
		return 0
	}
}

func Brackets(s string) error {
	bal := 0
	err := errors.New("invalid brackets")
	for _, i := range s {
		if i == '(' {
			bal++
			continue
		}
		if i == ')' {
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

func IsOperation(sign byte) bool {
	return sign == '+' || sign == '-' || sign == '*' || sign == '/' || sign == '^'
}

func LeftAssociative(sign byte) bool {
	return sign != '^'
}

func RPN(expr string) (*Stack, error) {
	output := NewStack()
	operators := NewStack()

	if err := Brackets(expr); err != nil {
		return NewStack(), err
	}

	for i := 0; i < len(expr); i++ {
		if expr[i] == ' ' || expr[i] == '\t' {
			continue
		}
		if '0' <= expr[i] && expr[i] <= '9' {
			startIndex := i
			for i < len(expr) && ('0' <= expr[i] && expr[i] <= '9' || expr[i] == DecimalSeparator) {
				i++
			}
			endIndex := i
			i--
			num, err := strconv.ParseFloat(expr[startIndex:endIndex], 64)
			if err != nil {
				return NewStack(), err
			}
			output.Push(StackElement{num: num})
			continue
		}
		if expr[i] == '(' {
			operators.Push(StackElement{sign: '('})
			continue
		}
		if expr[i] == ')' {
			for operators.Top().sign != '(' {
				output.Push(operators.Pop())
			}
			operators.Pop()
			continue
		}
		if expr[i] == '-' && (i == 0 || expr[i-1] == '(') {
			output.Push(StackElement{num: 0})
			operators.Push(StackElement{sign: '-'})
			continue
		}
		if IsOperation(expr[i]) {
			prior := Priority(expr[i])
			la := LeftAssociative(expr[i])
			for !operators.IsEmpty() {
				topSign := operators.Top().sign
				currentPrior := Priority(topSign)
				if currentPrior < prior || currentPrior == prior && !la {
					break
				}
				output.Push(operators.Pop())
			}
			operators.Push(StackElement{sign: expr[i]})
			continue
		}
		return NewStack(), fmt.Errorf("unsupported symbol: %v", expr[i])
	}
	for !operators.IsEmpty() {
		output.Push(operators.Pop())
	}
	return output, nil
}
