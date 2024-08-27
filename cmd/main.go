package main

import (
	"fmt"

	"github.com/namanag0502/func-calc/utils"
)

func calcExpression() {
	var result string = ""
	for {
		text, err := utils.ReadLine()
		if err != nil {
			fmt.Println("Readline error:", err.Error())
			return
		}

		if text == "exit" {
			break
		}

		result += text

		input := utils.ParseInputToStack(result)
		if len(input) == 0 {
			fmt.Println("Empty input, please provide an expression.")
			continue
		}

		res, err := utils.EvaluteStackToResult(input)
		if err != nil {
			fmt.Println("Error:", err.Error())
			continue
		}

		fmt.Print(res)
		result = res
	}
}

func main() {
	print("Input string to eval:")
	calcExpression()
}
