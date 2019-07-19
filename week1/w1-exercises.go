package week1

import (
	"bufio"
	"os"
	"fmt"
)

// DoMath read single expression from stdin and calculate result
// only support simple expression with 2 operands and 1 operator
// operands, operator are separated by 1 space
// operators supported: + - * /
func DoMath() {
	simpleExp := simpleExpression{}
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		expStr := scanner.Text()
		if err := simpleExp.parse(expStr); err != nil {
			fmt.Println(err)
			continue
		}

		showResult(expStr, &simpleExp)
	}	
}

// showResult show the result after calculating of an expression,
// for demonstrating use of expression interface
// that we can pass any exressions that satify expression interface
func showResult(expStr string, parsedExp expression) {
	result, err := parsedExp.calculate()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(expStr, "=", result)
	}
}


