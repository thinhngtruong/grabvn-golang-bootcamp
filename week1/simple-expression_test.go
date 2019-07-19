package week1

import (
	"testing"
)

func TestParse(t *testing.T) {
	simpleExp := simpleExpression{}

	checkNoError := func(t *testing.T, err error) {
		t.Helper()
		if err != nil {
			t.Errorf("expected no error but got one")
		}
	}
	
	checkExpectedErrorMsg := func(t *testing.T, expectedErr, err string) {
		t.Helper()
		if err != expectedErr {
			t.Errorf("expected %q but got %q", expectedErr, err)
		}
	}

	t.Run("1 + 1 should be a valid expression", func(t *testing.T) {
		result := simpleExp.parse("1 + 1")
		checkNoError(t, result)
	})
	t.Run("1 - 1 should be a valid expression", func(t *testing.T) {
		result := simpleExp.parse("1 - 1")
		checkNoError(t, result)
	})
	t.Run("1 * 1 should be a valid expression", func(t *testing.T) {
		result := simpleExp.parse("1 * 1")
		checkNoError(t, result)
	})
	t.Run("1 / 1 should be a valid expression", func(t *testing.T) {
		result := simpleExp.parse("1 / 1")
		checkNoError(t, result)
	})
	
	t.Run("1 % 1 should be a valid expression, even % operator is not supported", func(t *testing.T) {
		result := simpleExp.parse("1 % 1")
		checkNoError(t, result)
	})
	
	t.Run("1/ 1 should be an invalid expression", func(t *testing.T) {
		result := simpleExp.parse("1/ 1")
		checkExpectedErrorMsg(t, invalidExpresionMsg, result.Error())
	})
	t.Run("1+1 should be an invalid expression", func(t *testing.T) {
		result := simpleExp.parse("1+1")
		checkExpectedErrorMsg(t, invalidExpresionMsg, result.Error())
	})
	t.Run(" 1 + 1 should be an invalid expression", func(t *testing.T) {
		result := simpleExp.parse(" 1+1")
		checkExpectedErrorMsg(t, invalidExpresionMsg, result.Error())
	})
	t.Run("1 -1 should be an invalid expression", func(t *testing.T) {
		result := simpleExp.parse("1 -1")
		checkExpectedErrorMsg(t, invalidExpresionMsg, result.Error())
	})
	
	t.Run("a - 3 should be marked as invalid operand", func(t *testing.T) {
		result := simpleExp.parse("a - 3")
		checkExpectedErrorMsg(t, invalidOperandMsg, result.Error())
	})
	t.Run("5 * 3 should be marked as invalid operand", func(t *testing.T) {
		result := simpleExp.parse("a - 3")
		checkExpectedErrorMsg(t, invalidOperandMsg, result.Error())
	})
}

func TestCalculate(t *testing.T) {
	checkNoError := func(t *testing.T, err error) {
		t.Helper()
		if err != nil {
			t.Errorf("expected no error but got one")
		}
	}
	
	checkExpectedErrorMsg := func(t *testing.T, expectedErr, err string) {
		t.Helper()
		if err != expectedErr {
			t.Errorf("expected %q but got %q", expectedErr, err)
		}
	}
	
	checkExpectedResult := func(t *testing.T, expected, result float64) {
		t.Helper()
		if result != expected {
			t.Errorf("expected %f but got %f", expected, result)
		}
	}

	t.Run("1 + 1 should be equal 2", func(t *testing.T) {
		simpleExp := simpleExpression{1, 1, "+"}
		result, _ := simpleExp.calculate()
		checkExpectedResult(t, 2, result)
	})
	t.Run("2 - 1 should be equal 1", func(t *testing.T) {
		simpleExp := simpleExpression{2, 1, "-"}
		result, _ := simpleExp.calculate()
		checkExpectedResult(t, 1, result)
	})
	t.Run("1 * 9 should be equal 9", func(t *testing.T) {
		simpleExp := simpleExpression{1, 9, "*"}
		result, _ := simpleExp.calculate()
		checkExpectedResult(t, 9, result)
	})
	t.Run("9 / 3 should be equal 3", func(t *testing.T) {
		simpleExp := simpleExpression{9, 3, "/"}
		result, _ := simpleExp.calculate()
		checkExpectedResult(t, 3, result)
	})

	t.Run("1 + 1 should be no error", func(t *testing.T) {
		simpleExp := simpleExpression{1, 1, "+"}
		_, err := simpleExp.calculate()
		checkNoError(t, err)
	})
	t.Run("2 - 1 should be no error", func(t *testing.T) {
		simpleExp := simpleExpression{2, 1, "-"}
		_, err := simpleExp.calculate()
		checkNoError(t, err)
	})
	t.Run("1 * 9 should be no error", func(t *testing.T) {
		simpleExp := simpleExpression{1, 9, "*"}
		_, err := simpleExp.calculate()
		checkNoError(t, err)
	})
	t.Run("9 / 3 should be no error", func(t *testing.T) {
		simpleExp := simpleExpression{9, 3, "/"}
		_, err := simpleExp.calculate()
		checkNoError(t, err)
	})
	
	t.Run("9 / 0 should be marked as divided by zero", func(t *testing.T) {
		simpleExp := simpleExpression{9, 0, "/"}
		_, err := simpleExp.calculate()
		checkExpectedErrorMsg(t, dividedZyZeroMsg, err.Error())
	})
	t.Run("9 % 2 should be marked as invalid operator", func(t *testing.T) {
		simpleExp := simpleExpression{9, 3, "%"}
		_, err := simpleExp.calculate()
		checkExpectedErrorMsg(t, invalidOparatorMsg, err.Error())
	})
}