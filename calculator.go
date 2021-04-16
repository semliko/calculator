package calculator

import (
	"fmt"
)

type Calc struct {
	FirstNumber  int
	SecondNumber int
	Operator     string
}

func (c Calc) AddNumbers() int {
	return c.FirstNumber + c.SecondNumber
}

func (c Calc) SubtractNumbers() int {
	return c.FirstNumber - c.SecondNumber
}

func (c Calc) MultiplyNumbers() int {
	return c.FirstNumber * c.SecondNumber
}

func (c Calc) DivideNumbers() int {
	return c.FirstNumber / c.SecondNumber
}

func (c Calc) Percents() int {
	return c.FirstNumber % c.SecondNumber
}

func (c Calc) Calculate() int {

	var result int
	var op string

	op = c.Operator

	switch op {
	case "+":
		result = c.AddNumbers()
	case "-":
		result = c.SubtractNumbers()
	case "*":
		result = c.MultiplyNumbers()
	case "/":
		result = c.DivideNumbers()
	case "%":
		result = c.Percents()
	default:
		fmt.Println("Given operator is invalid")
	}
	return result
}
