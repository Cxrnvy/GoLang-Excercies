package main

import (
	"fmt"
)

func main() {
	votos := map[string]int{
		"deportes":    0,
		"videojuegos": 0,
		"cine":        0,
		"musica":      0,
	}

	fmt.Println("ENCUESTA DE ACTIVIDADES")
	fmt.Println("Opciones válidas: deportes, videojuegos, cine, musica")

	votosIngresados := 0
	for votosIngresados < 5 {
		var voto string
		fmt.Printf("Ingrese el voto %d: ", votosIngresados+1)
		fmt.Scanln(&voto)

		if _, existe := votos[voto]; existe {
			votos[voto] = votos[voto] + 1
			votosIngresados++
		} else {
			fmt.Println("Opción incorrecta. Escriba deportes, videojuegos, cine o musica.")
		}
	}

	fmt.Println("\nRESULTADOS")

	for actividad, cantidad := range votos {
		fmt.Printf("- %s: %d votos\n", actividad, cantidad)
	}

	var maxVotos int = -1
	var actividadGanadora string = ""

	for actividad, cantidad := range votos {
		if cantidad > maxVotos {
			maxVotos = cantidad
			actividadGanadora = actividad
		}
	}

	fmt.Printf("\nLa actividad ganadora es: %s con %d votos\n", actividadGanadora, maxVotos)
}
