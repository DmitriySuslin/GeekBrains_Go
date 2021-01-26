package main

import (
	"fmt"
	"log"
	"math/big"
	"strconv"
)

func main() {
	var N string
	var n, f int64
	var err error

	fmt.Println("Поиск простых от 1 до N")
	fmt.Print("Ведите ")
	_, err = fmt.Scan(&N)
	if err != nil {
		log.Fatalf("Ошибка ввода N числа %v", err)
	}
	n, err = strconv.ParseInt(N, 10, 64)
	if err != nil {
		log.Fatalf("Ошибка ввода N числа %v", err)
	}

	for f = 1; f <= n; f++ {
		if big.NewInt(f).ProbablyPrime(0) {
			fmt.Print(" ", f)
		}

	}
	fmt.Println()
}
