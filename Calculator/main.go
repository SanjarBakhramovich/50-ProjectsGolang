package main

// я хочу написать код калькулятора
import "fmt"



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
func main() {
	fmt.Println("Welcome to calculator")

	

	// создать функцию для добавления чисел
	// создать функцию для вычитания чисел

	// создать функцию для умножения чисел
	// создать функцию для деления чисел

	// теперь нужно эти функции реализовать внутри main функции
	// что бы получился калькулятор

	var a, b, result float64
	fmt.Print("Введите первое число:")
	fmt.Scan(&a)
	fmt.Print("Введите второе число:")
	fmt.Scan(&b)
	fmt.Print("Выберите операцию:\n1 - Сложение\n2 - Вычитание\n3 - Умножение\n4 - Деление\n")
	var op int
	fmt.Scan(&op)
	
	switch op {
	case 1: 
		result = add(a, b)
	case 2:
		result = subtract(a, b)
	case 3: 
		result = multiply(a, b)
	case 4: 
		result = divide(a,b)		
	default:
		(fmt.Println("Неверная операция, пожалуйста проверьте правильность операции"))	
	}
	fmt.Printf("Результат операции: %.2f\n", result)
}