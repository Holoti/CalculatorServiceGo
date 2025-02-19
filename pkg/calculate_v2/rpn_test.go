package calculate_v2

import (
	"fmt"
	"log"
	"testing"
)

func TestRPN(t *testing.T) {
	fmt.Println(byte('+'), byte('-'), byte('*'), byte('/'))
	expr := "(26/(-(81))-46)*62*30-(3)*(85)-0/17"
	// expr := "(2+21)*2"
	// expr := "0.1+0.2"
	output, err := RPN(expr)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(output.items)
}
