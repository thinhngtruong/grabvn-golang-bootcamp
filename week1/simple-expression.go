package week1

import (
	"strings"
	"strconv"
)

type simpleExpression struct {
	operand1 float64
	operand2 float64
	operator string
}

func (expression *simpleExpression) parse (expressionString string) (err error)  {
	elements := strings.Fields(expressionString)

	if len(elements) != 3 {
		return errInvalidExpresion
	}

	expression.operator = elements[1]

	if expression.operand1, err = strconv.ParseFloat(elements[0], 64); err != nil {
		return errInvalidOperand
	}
	
	if expression.operand2, err = strconv.ParseFloat(elements[2], 64); err != nil{
		return errInvalidOperand
	}

	return
}

func (expression *simpleExpression) calculate() (float64, error) {
	switch expression.operator {
	case "+":
		return expression.operand1 + expression.operand2, nil
	case "-":
		return expression.operand1 - expression.operand2, nil
	case "*":
		return expression.operand1 * expression.operand2, nil
	case "/":
		if expression.operand2 == 0 {
			return 0, errDividedZyZero
		}
		return expression.operand1 / expression.operand2, nil
	default: 
		return 0, errInvalidOperator
	}
}