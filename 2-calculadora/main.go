package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type calc struct {
	
}

func (calc) operate(values string, operator string) int {
	cleanInput := strings.Split(values, operator)
	value1 := parse(cleanInput[0])
	value2 := parse(cleanInput[1])
	// Using Case
	switch operator {
	case "+":
		operation := value1 + value2
		fmt.Println("Operation Result:", operation)
		return operation
	case "-":
		operation := value1 - value2
		fmt.Println("Operation Result:", operation)
		return operation
	case "*":
		operation := value1 * value2
		fmt.Println("Operation Result:", operation)
		return operation
	case "/":
		operation := value1 / value2
		fmt.Println("Operation Result:", operation)
		return operation
	default:
		fmt.Println(operator, "Is not supported")
		return 0
	}
}

func parse(input string) int {
	parseValue, _ := strconv.Atoi(input) // _ wildcard para recibir la variables pero no utilizarla
	return parseValue
}

func readUserInput() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	inputUser := scanner.Text()
	fmt.Printf("input: %v of type: %T \n", inputUser, inputUser)
	return inputUser
}

func main()  {
	// Operation eg: 2+2
	userInput := readUserInput()
	// Operator eg: +
	operator := readUserInput()

	c := calc{}.operate(userInput, operator)
	println(c)
}