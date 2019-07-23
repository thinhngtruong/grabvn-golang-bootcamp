package week1

import (
	"bufio"
	"fmt"
	"os"
)

// DoMath read single expression from stdin and calculate result
// only support simple expression with 2 operands and 1 operator
// operands, operator are separated by 1 space
// operators supported: + - * /
func DoMath() {
	simpleExpression := simpleExpression{}
	
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		expressionString := scanner.Text()
		if err := simpleExpression.parse(expressionString); err != nil {
			fmt.Println(err)
			continue
		}

		result, err := simpleExpression.calculate()
		if err != nil {
			fmt.Println(err)
			continue
		}

		fmt.Println(expressionString, "=", result)
	}
}
