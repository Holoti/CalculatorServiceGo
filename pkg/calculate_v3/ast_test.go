package calculate_v3

import (
	"fmt"
	"log"
	"testing"
)

func TestAST(t *testing.T) {
	expr := "2*3+4*5"
	rpn, err := RPN(expr)
	if err != nil {
		log.Fatal(err)
	}
	result := AST(rpn)
	fmt.Println(result)
}
