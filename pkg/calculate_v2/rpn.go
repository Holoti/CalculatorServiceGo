package calculate_v2

import (
	"errors"
	"log"
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
	result := s.Top()
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
	case '+', '-', ')': // TODO am i sure about ) priority?
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

func RPN(expr string) { // TODO this is concurrent; how to return result?
	output := NewStack()
	operators := NewStack()

	if err := Brackets(expr); err != nil {
		// TODO handle error
	}

	for i := range expr {
		if expr[i] == ' ' || expr[i] == '\t' {
			continue
		}
		if '0' <= expr[i] && expr[i] <= '9' {
			startIndex := i
			for '0' <= expr[i] && expr[i] <= '9' || expr[i] == DecimalSeparator {
				i++
			}
			endIndex := i
			i--
			num, err := strconv.ParseFloat(expr[startIndex:endIndex], 64)
			if err != nil {
				log.Fatal(err) // TODO handle error
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
				// TODO do the operations somehow
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
			topSign := operators.Top().sign
			for topSign != '(' && (Priority(topSign) > prior || (Priority(topSign) == prior && la)) {
				// TODO do the operations somehow
			}
			operators.Push(StackElement{sign: expr[i]})
			continue
		}
		// TODO handle error: wrong symbol
	}
	for !operators.IsEmpty() {
		// TODO do the operations somehow
	}
	if len(output.items) != 1 {
		// TODO handle error
	}
	// TODO return the output.Top().num somehow
}
