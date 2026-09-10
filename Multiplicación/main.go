package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Print("Ingresa un número entero positivo: ")
	fmt.Scan(&n)

	fmt.Println("***Tabla de multiplicar***")
	
	for i := 1; i <= 10; i++ {
		fmt.Printf("%d x %d = %d\n", n, i, n*i)
	}
	fmt.Println("**********************")

	if n%2 == 0 {
		fmt.Println("El número es Par.")
	} else {
		fmt.Println("El número es Impar.")
	}

	fmt.Println("\n")
	switch {
	case n >= 1 && n <= 5:
		fmt.Println("Número pequeño")
	case n >= 6 && n <= 10:
		fmt.Println("Número mediano")
	case n > 10:
		fmt.Println("Número grande")
	}
}
