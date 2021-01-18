package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	var a int32

	fmt.Print("Введите 3-х значное число ")
	_, err := fmt.Scan(&a)
	if err != nil {
		log.Print(err)
		os.Exit(128)
	}

	if a > 999 || a < 100 {
		log.Println("Просили вас вести 3-х значное число, а вы ввели", a)
		os.Exit(128)
	}

	fmt.Printf("%d сот., %d дес., %d ед\n", int8(a/100), int8(a/10%10), int8(a%10))
}
