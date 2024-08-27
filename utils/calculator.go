package utils

import (
	"fmt"
	"strconv"
)

// Calculate is a pure function that performs the operation on two operands
func Calculate(a, b, op string) (string, error) {
	operandA, err := strconv.ParseFloat(a, 64)
	if err != nil {
		return "", fmt.Errorf("cannot convert '%s' to float", a)
	}
	operandB, err := strconv.ParseFloat(b, 64)
	if err != nil {
		return "", fmt.Errorf("cannot convert '%s' to float", b)
	}

	switch op {
	case "+":
		return strconv.FormatFloat(operandA+operandB, 'f', -1, 64), nil
	case "-":
		return strconv.FormatFloat(operandA-operandB, 'f', -1, 64), nil
	case "x":
		return strconv.FormatFloat(operandA*operandB, 'f', -1, 64), nil
	case "/":
		if operandB == 0 {
			return "", fmt.Errorf("division by zero")
		}
		return strconv.FormatFloat(operandA/operandB, 'f', -1, 64), nil
	default:
		return "", fmt.Errorf("operator '%s' not supported", op)
	}
}
