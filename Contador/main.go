package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Print("Ingresa un número entero positivo: ")
	fmt.Scan(&n)


	temp := n
	digitos := 0
	suma := 0

	for temp > 0 {

		ultimoDigito := temp % 10 
		
		suma += ultimoDigito
		
		// Le quitamos el último dígito al número
		temp = temp / 10 
		digitos++
	}

	fmt.Printf("El número %d tiene %d dígitos.\n", n, digitos)
	fmt.Printf("La suma de sus dígitos es: %d\n", suma)

	switch digitos {
	case 1:
		fmt.Println("Número de una cifra")
	case 2:
		fmt.Println("Número de dos cifras")
	case 3:
		fmt.Println("Número de tres cifras")
	default:
		fmt.Println("Número de varias cifras")
	}
}