package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

func main() {

	var aF, bF, res float64
	var a, b, op string
	var err error
	// ввод первого числа -------
	fmt.Print("Введите первое число: ")
	_, err = fmt.Scan(&a)
	if err != nil {
		log.Fatalf("ошибка ввода первого числа: %v", err)
	}

	aF, err = strconv.ParseFloat(a, 64)
	if err != nil {
		log.Fatalf("ошибка ввода первого числа: %v", err)
	}
	// ввод первого числа -------

	// ввод второго числа -------
	fmt.Print("Введите второе число: ")
	_, err = fmt.Scan(&b)
	if err != nil {
		log.Fatalf("Вы ввели не число а %s", a)
	}

	bF, err = strconv.ParseFloat(b, 64)
	if err != nil {
		log.Fatalf("ошибка ввода первого числа: %v", err)
	}
	// ввод второго числа -------

	fmt.Print("Введите арифметическую операцию (+, -, *, /, %, ^): ")
	_, err = fmt.Scanln(&op)
	if err != nil {
		log.Fatalf("Вы ввели не число а %s", a)
	}

	switch op {
	case "+":
		res = aF + bF
	case "-":
		res = (aF - bF)
	case "*":
		res = aF * bF
	case "/":
		res = aF / bF
	case "%":
		res = float64(int32(aF) % int32(bF))
	case "^":
		res = math.Pow(float64(aF), float64(bF))
	default:
		fmt.Println("Операция выбрана неверно")
		os.Exit(1)
	}

	fmt.Printf("Результат выполнения операции: %f\n", res)
}
