package main

// я хочу написать код калькулятора
import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)



func add(a, b float64) float64{
	return a + b
}

func subtract(a, b float64) float64 {
  return a - b
}

func multiply(a, b float64) float64 {
  return a * b
}

func divide(a, b float64) float64 {
  return a / b
}

func power (a, b float64) float64 {
	return math.Pow(a, b)
}

func modulus(a, b float64) float64 {
	return math.Mod(a, b)
}

func calculate(expression string) (float64, error) {
	var result float64
	var err error
	var op string
	var a, b float64

	switch {
	case strings.Contains(expression, "**"):
		op = "**"
	case strings.Contains(expression, "%"):
		op = "%"
	case strings.Contains(expression, "+"):
		op = "+"
	case strings.Contains(expression, "-"):
		op = "-"
	case strings.Contains(expression, "*"):
		op = "*"
	case strings.Contains(expression, "/"):
		op = "/"
		default:
			return -1, fmt.Errorf("Неверная операция, пожалуйста проверьте правильность операции")
	
			
	}

	parts := strings.Split(expression, op)
	if len(parts) != 2 {
		return 0, fmt.Errorf("Неверная операция, пожалуйста проверьте правильность операции")
	}

	a, err = strconv.ParseFloat(strings.TrimSpace(parts[0]), 64)
	if err != nil {
		return 0, fmt.Errorf("Неверное первое число")
	}

	b, err = strconv.ParseFloat(strings.TrimSpace(parts[1]), 64)
	if err != nil {
		return 0, fmt.Errorf("Неверное второе число")
	}

	switch op {
	case "+":
		result = add(a, b)
	case "-":
		result = subtract(a, b)
	case "*": 
		result = multiply(a, b)
	case "/":
		result = divide(a,b)
	case "**":
		result = power(a, b)
	case "%":
		result = modulus(a, b)			
	}
	return result, nil
}
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Welcome to calculator")
	fmt.Println("Можете ввести то что вы хотите посчитать")

	for {
		fmt.Print("> ")
		scanner.Scan()

		expression := scanner.Text()

		if strings.TrimSpace(expression) == "end" {
			fmt.Println("Завершение работы калькулятора.")
			break
		}

		result, err := calculate(expression)
		if err != nil {
			fmt.Println("Ошибка", err)
		} else {
			fmt.Printf("Результат: %.2f\n", result)
		}
	}
}
