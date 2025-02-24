package calculate_v3

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/holoti/CalculatorServiceGo/pkg/models"
)

var DecimalSeparator byte = '.' // TODO read this from .env

type Token struct {
	Num  float64
	Sign byte
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

func RPN(expr string) ([]Token, error) {
	output := make([]Token, 0)
	operators := models.NewStack[Token]()

	if err := Brackets(expr); err != nil {
		return nil, err
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
				return nil, err
			}
			output = append(output, Token{Num: num})
			continue
		}
		if expr[i] == '(' {
			operators.Push(Token{Sign: '('})
			continue
		}
		if expr[i] == ')' {
			for operators.Top().Sign != '(' {
				output = append(output, operators.Pop())
				// output.Push(operators.Pop())
			}
			operators.Pop()
			continue
		}
		if expr[i] == '-' && (i == 0 || expr[i-1] == '(') {
			output = append(output, Token{Num: 0})
			operators.Push(Token{Sign: '-'})
			continue
		}
		if IsOperation(expr[i]) {
			prior := Priority(expr[i])
			la := LeftAssociative(expr[i])
			for !operators.IsEmpty() {
				topSign := operators.Top().Sign
				currentPrior := Priority(topSign)
				if currentPrior < prior || currentPrior == prior && !la {
					break
				}
				output = append(output, operators.Pop())
			}
			operators.Push(Token{Sign: expr[i]})
			continue
		}
		return nil, fmt.Errorf("unsupported symbol: %v", expr[i])
	}
	for !operators.IsEmpty() {
		output = append(output, operators.Pop())
	}
	return output, nil
}
