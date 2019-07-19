package week1

import (
	"errors"
	"strings"
	"strconv"
)

type simpleExpression struct {
	operand1 float64
	operand2 float64
	operator string
}

func (exp *simpleExpression) parse (expStr string) (err error)  {
	expElements := strings.Split(expStr, " ")

	if len(expElements) != 3 {
		return errors.New(invalidExpresionMsg)
	}

	exp.operator = expElements[1]

	if exp.operand1, err = strconv.ParseFloat(expElements[0], 64); err != nil {
		return errors.New(invalidOperandMsg)
	}
	
	if exp.operand2, err = strconv.ParseFloat(expElements[2], 64); err != nil{
		return errors.New(invalidOperandMsg)
	}

	return
}

func (exp *simpleExpression) calculate() (float64, error) {
	switch exp.operator {
	case "+":
		return exp.operand1 + exp.operand2, nil
	case "-":
		return exp.operand1 - exp.operand2, nil
	case "*":
		return exp.operand1 * exp.operand2, nil
	case "/":
		if exp.operand2 == 0 {
			return 0, errors.New(dividedZyZeroMsg)
		}
		return exp.operand1 / exp.operand2, nil
	default: 
		return 0, errors.New(invalidOparatorMsg)
	}
}