package utils

type Stack []string

func Push(stack Stack, v string) Stack {
	return append(stack, v)
}

func Pop(stack Stack) (string, Stack) {
	l := len(stack)
	if l == 0 {
		return "", stack
	}
	return stack[l-1], stack[:l-1]
}
