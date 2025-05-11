package main

import (
	"fmt"
)

type Expression struct {
	Left     int
	Operator rune
	Right    int
}

func getUserInput(expr *Expression) {
	fmt.Println("Input left operand: ")
	fmt.Scanln(expr.Left)
}

func main() {
	expr := Expression{}
	getUserInput(&expr)

}
