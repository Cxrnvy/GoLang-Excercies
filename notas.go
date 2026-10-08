package main

import (
	"fmt"
)

func main() {
	var notas [6][4]float64
	var sumaClase float64

	fmt.Println("INGRESO DE NOTAS")

	for i := 0; i < 6; i++ {
		fmt.Printf("\nEstudiante %d\n", i+1)
		for j := 0; j < 4; j++ {
			fmt.Printf("Ingrese la nota de la materia %d: ", j+1)
			fmt.Scanln(&notas[i][j])
		}
	}

	fmt.Println("\nRESULTADOS")

	for i := 0; i < 6; i++ {
		sliceNotas := notas[i][:]

		var sumaEstudiante float64 = 0
		maxNota := sliceNotas[0]
		minNota := sliceNotas[0]

		for j := 0; j < len(sliceNotas); j++ {
			notaActual := sliceNotas[j]
			sumaEstudiante = sumaEstudiante + notaActual

			if notaActual > maxNota {
				maxNota = notaActual
			}

			if notaActual < minNota {
				minNota = notaActual
			}
		}

		promedioEstudiante := sumaEstudiante / 4.0
		sumaClase = sumaClase + sumaEstudiante

		fmt.Printf("Estudiante %d -> Promedio: %.2f | Nota Mayor: %.2f | Nota Menor: %.2f\n",
			i+1, promedioEstudiante, maxNota, minNota)
	}

	promedioGeneral := sumaClase / 24.0
	fmt.Printf("\nPromedio general de la clase: %.2f\n", promedioGeneral)
}
