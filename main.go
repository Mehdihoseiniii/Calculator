package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type OperandError rune

func (oe OperandError) Error() string {
	return fmt.Sprintf("operand %s is not valid", string(oe))
}

func main() {
	fmt.Println("Welcome to program")
	input := getinput()
	// convert to float
	n1, n2, op, err := pars(input)
	if err != nil {
		fmt.Print(err)

	}
	// calculator
	result, err := calculate(n1, n2, op)
	if err != nil {
		fmt.Print(err)
	}
	//Print
	fmt.Println(result)
}

func getinput() string {
	var input string
	fmt.Print("inter a number : ")
	fmt.Scanln(&input)
	return input
}

func pars(input string) (number_1, number_2 float64, operand rune, err error) {
	for _, operand := range "+-*/%" {
		subs := strings.Split(input, string(operand))
		if len(subs) != 1 {
			number_1, err = strconv.ParseFloat(subs[0], 64)
			if err != nil {
				return 0, 0, 0, err
			}
			number_2, err = strconv.ParseFloat(subs[1], 64)
			if err != nil {
				return 0, 0, 0, err
			}
			return number_1, number_2, operand, nil
		}
	}
	return 0, 0, 0, OperandError(operand)
}
func calculate(number_1, number_2 float64, operand rune) (float64, error) {
	switch operand {
	case '+':
		return number_1 + number_2, nil
	case '-':
		return number_1 - number_2, nil
	case '*':
		return number_1 * number_2, nil
	case '/':
		return number_1 / number_2, nil
	case '%':
		return math.Mod(number_1, number_2), nil
	case '^':
		return math.Pow(number_1, number_2), nil

	}
	return 0, OperandError(operand)
}
