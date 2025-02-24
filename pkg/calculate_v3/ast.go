package calculate_v3

import "github.com/holoti/CalculatorServiceGo/pkg/models"

type Node struct {
	Token
	first, second *Node
}

func AST(rpn []Token) models.Stack[Node] {
	result := models.NewStack[Node]()
	for _, item := range rpn {
		if item.Sign == 0 {
			result.Push(Node{Token: item})
		} else {
			var a, b Node
			if item.Sign != '^' {
				b = result.Pop()
			}
			a = result.Pop()
			result.Push(Node{Token: item, first: &a, second: &b})
		}
	}
	return *result
}
