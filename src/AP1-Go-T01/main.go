package main

import (
	"fmt"
)

type Expression struct {
	Left     float64
	Operator rune
	Right    float64
}

func getUserInput(expr *Expression) {
	fmt.Println("Input left operand: ")

	var err error
	for {
		_, err = fmt.Scanln(&expr.Left)
		if err == nil {
			break
		} else {
			fmt.Println("invalid input")
		}
	}

	fmt.Println("Input operation: ")
	var op string
	for {
		_, err = fmt.Scanln(&op)
		if err == nil && len(op) == 1 {
			switch op[0] {
			case '+', '-', '*', '/':
				expr.Operator = rune(op[0])
			default:
				fmt.Println("invalid operator. Try again")
				continue // продолжить цикл при неверном операторе
			}
			break // выйти из цикла при успешном вводе
		} else {
			fmt.Println("Please enter a single character (+, -, *, /).")
		}
	}

	fmt.Println("Input right operand: ")
	for {
		_, err = fmt.Scanln(&expr.Right)
		if err != nil {
			fmt.Println("Invalid input")
			continue
		}
		if expr.Operator == '/' && expr.Right == 0 {
			fmt.Println("division by zero")
			continue
		}
		break
	}
}

func calcExp(expr *Expression) float64 {
	var result float64

	switch expr.Operator {
	case '+':
		result = expr.Left + expr.Right
	case '-':
		result = expr.Left - expr.Right
	case '*':
		result = expr.Left * expr.Right
	case '/':
		result = expr.Left / expr.Right
	}
	return result
}

func main() {
	expr := Expression{}
	getUserInput(&expr)
	result := calcExp(&expr)
	fmt.Printf("%.3f\n", result)
}
