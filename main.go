package main

import (
	"fmt"
  )

// Программа для вычисления факториала
func main() {
	fmt.Println("========================FACTORIAL=========================")
	fmt.Println("Введите число для вычисления")
	num := 0
	fmt.Scanf("%d", &num) // cканим число для вычисления
	start := 1
	result := 1
	for start <= num {
		if num == 0 {
			goto Result
		}
		result = start * (result)
		start++
	}
Result:
	fmt.Printf("Factorial: %v\n", result)
}
