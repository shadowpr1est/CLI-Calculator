package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("=== Terminal Calculator ===")
	fmt.Println("Supported operations: +, -, *, /, sqrt, ^")

	for {
		fmt.Print("\nEnter operation (+, -, *, /, sqrt, ^): ")
		opInput, _ := reader.ReadString('\n')
		op := strings.TrimSpace(opInput)

		var num1, num2 float64
		var err error

		if op == "sqrt" {
			fmt.Print("Enter number: ")
			numStr, _ := reader.ReadString('\n')
			numStr = strings.TrimSpace(numStr)
			num1, err = strconv.ParseFloat(numStr, 64)
			if err != nil {
				fmt.Println("Invalid number.")
				continue
			}
			if num1 < 0 {
				fmt.Println("Cannot take square root of a negative number.")
				continue
			}
			result := math.Sqrt(num1)
			fmt.Printf("Result: %.4f\n", result)
		} else {
			fmt.Print("Enter first number: ")
			num1Str, _ := reader.ReadString('\n')
			num1Str = strings.TrimSpace(num1Str)
			num1, err = strconv.ParseFloat(num1Str, 64)
			if err != nil {
				fmt.Println("Invalid first number.")
				continue
			}

			fmt.Print("Enter second number: ")
			num2Str, _ := reader.ReadString('\n')
			num2Str = strings.TrimSpace(num2Str)
			num2, err = strconv.ParseFloat(num2Str, 64)
			if err != nil {
				fmt.Println("Invalid second number.")
				continue
			}

			switch op {
			case "+":
				fmt.Printf("Result: %.4f\n", num1+num2)
			case "-":
				fmt.Printf("Result: %.4f\n", num1-num2)
			case "*":
				fmt.Printf("Result: %.4f\n", num1*num2)
			case "/":
				if num2 == 0 {
					fmt.Println("Error: Division by zero.")
				} else {
					fmt.Printf("Result: %.4f\n", num1/num2)
				}
			case "^":
				fmt.Printf("Result: %.4f\n", math.Pow(num1, num2))
			default:
				fmt.Println("Unsupported operation.")
			}
		}

		fmt.Print("\nDo you want to perform another operation? (y/n): ")
		again, _ := reader.ReadString('\n')
		if strings.ToLower(strings.TrimSpace(again)) != "y" {
			fmt.Println("Goodbye!")
			break
		}
	}
}
