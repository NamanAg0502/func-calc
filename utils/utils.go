package utils

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func ReadLine() (string, error) {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	return input[:len(input)-1], nil
}

func IsOperator(ch rune) bool {
	return ch == '+' || ch == '-' || ch == 'x' || ch == '/'
}

func ParseInputToStack(input string) Stack {
	var stack Stack
	var num string

	for _, ch := range input {
		if IsOperator(ch) {
			if num != "" {
				stack = Push(stack, strings.TrimSpace(num))
				num = ""
			}
			stack = Push(stack, string(ch))
		} else {
			num += string(ch)
		}
	}
	if num != "" {
		stack = Push(stack, strings.TrimSpace(num))
	}

	return stack
}

func EvaluteStackToResult(input []string) (string, error) {
	if len(input) == 0 {
		return "", fmt.Errorf("no input provided")
	}

	res := input[0]
	input = input[1:]

	for len(input) > 0 {
		operator := input[0]
		input = input[1:]

		if len(input) == 0 {
			return "", fmt.Errorf("incomplete expression")
		}

		nextOperand := input[0]
		input = input[1:]

		var err error
		res, err = Calculate(res, nextOperand, operator)
		if err != nil {
			return "", err
		}
	}

	return res, nil
}
