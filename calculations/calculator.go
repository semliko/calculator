package calculator

import "fmt"

type Calc struct {
	FirstNumber  float64
	SecondNumber float64
	Operator     string
}

type CalculationResult interface {
	Calculate() float64
	AddNumbers() float64
	SubtractNumbers() float64
	MultiplyNumbers() float64
	DivideNumbers() float64
	Percents() float64
	Integer() (int, int)
	ReturnFields() (float64, float64, string)
}

func (c Calc) Integer() (int, int) {
	return int(c.FirstNumber), int(c.SecondNumber)

}
func (c Calc) AddNumbers() float64 {
	return c.FirstNumber + c.SecondNumber
}

func (c Calc) SubtractNumbers() float64 {
	return c.FirstNumber - c.SecondNumber
}

func (c Calc) MultiplyNumbers() float64 {
	return c.FirstNumber * c.SecondNumber
}

func (c Calc) DivideNumbers() float64 {
	return c.FirstNumber / c.SecondNumber
}

func (c Calc) Percents() float64 {
	var result, firstNumber, secondNumber int
	firstNumber, secondNumber = c.Integer()
	result = firstNumber % secondNumber
	return float64(result)
}

func (c Calc) Calculate() float64 {

	var result float64
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

func (c Calc) ReturnFields() (float64, float64, string) {
	return c.FirstNumber, c.SecondNumber, c.Operator
}

func PrintCalculation(c CalculationResult) {
	var result, num1, num2 float64
	var operator string
	num1, num2, operator = c.ReturnFields()
	fmt.Println(c.Calculate())
	fmt.Println("====================See the result below===========================")
	fmt.Println(num1, " ", operator, " ", num2, " ", "= ", result)
}
