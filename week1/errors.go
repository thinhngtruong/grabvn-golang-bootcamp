package week1

import "errors"

const (
	invalidExpresionMessage = "invalid expression, only support simple expression with 2 operands and 1 operator"
	invalidOperandMessage   = "invalid operand, operands must be number"
	dividedZyZeroMessage    = "divided by zero"
	invalidOperatorMessage  = "invalid oparator, only support + - * /"
)

var (
	errInvalidExpresion = errors.New(invalidExpresionMessage)
	errInvalidOperand   = errors.New(invalidOperandMessage)
	errDividedZyZero    = errors.New(dividedZyZeroMessage)
	errInvalidOperator  = errors.New(invalidOperatorMessage)
)
