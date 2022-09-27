package main

import (
	"fmt"
	"strconv"
)

func main() {
	opMap := map[string]opFuncType{
		"+": add,
		"-": sub,
		"*": mul,
		"/": div,
	}

	expressions := [][]string{
		{"1", "+", "2"},
		{"2", "-", "3"},
		{"8", "*", "2"},
		{"15", "/", "2"},
		{"15", "%", "2"},
		{"15"},
		{"Two", "-", "One"},
	}

	for _, expression := range expressions {
		if len(expression) < 3 {
			fmt.Println("invalid expression", expression)
			continue
		}

		p1, err := strconv.Atoi(expression[0])
		if err != nil {
			fmt.Println(err)
			continue
		}

		op := expression[1]
		opFunc, ok := opMap[op]
		if !ok {
			fmt.Println("unsupported operator", op)
			continue
		}

		p2, err := strconv.Atoi(expression[2])
		if err != nil {
			fmt.Println(err)
			continue
		}

		result := opFunc(p1, p2)
		fmt.Println(result)
	}
}

type opFuncType func(i, j int) int

func add(i, j int) int { return i + j }
func sub(i, j int) int { return i - j }
func mul(i, j int) int { return i * j }
func div(i, j int) int { return i / j }
