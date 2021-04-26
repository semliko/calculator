package main

import (
	calculator "calculator/calculations"
	"fmt"
)

func main() {
	var operator string
	var num1, num2 float64
	fmt.Println("Enter first number")
	fmt.Scan(&num1)
	fmt.Println("Enter math function from these |+,-,*,/,%| ")
	fmt.Scan(&operator)
	fmt.Println("Enter second number")
	fmt.Scan(&num2)
	calculation := calculator.Calc{FirstNumber: num1, SecondNumber: num2, Operator: operator}
	calculator.PrintCalculation(calculation)

}
