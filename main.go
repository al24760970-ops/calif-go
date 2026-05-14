package main

import "fmt"

func main() {
	// Usando constantes para que el profe no me regañe
	const NOTA_MINIMA = 6.0
	const ASISTENCIA_OK = 75

	// El array de 10 alumnos (bueno, un slice)
	// Mezclo var con := para cumplir los requisitos
	var listaEstudiantes = []struct {
		nombre  string
		notas   []float64
		faltas  int
	}{
		{"Ana García", []float64{9.5, 8.2, 10, 9.1, 9.7}, 95},
		{"Luis Pérez", []float64{6.1, 7.0, 5.8, 6.2, 6.5}, 80},
		{"María López", []float64{8.0, 7.5, 9.2, 6.8, 8.5}, 60},
		{"Carlos Ruiz", []float64{4.0, 5.2, 3.5, 4.1, 5.0}, 92},
		{"Elena Sanz", []float64{10, 9.8, 10, 10, 9.9}, 100},
		{"Jordi Alba", []float64{7.2, 6.8, 7.4, 8.1, 6.3}, 85},
		{"Lucía Gil", []float64{3.2, 2.8, 4.1, 5.2, 3.8}, 45},
		{"David Romero", []float64{4.1, 4.3, 4.6, 3.9, 4.2}, 82},
		{"Sofía Vega", []float64{8.5, 9.0, 8.7, 9.1, 8.8}, 94},
		{"Raúl Cano", []float64{5.8, 6.2, 5.5, 6.0, 5.9}, 72},
	}

	var totalClase float64
	aprobadosContador := 0
	mejorNota := 0.0
	peorNota := 11.0
	var elMejor, elPeor string

	fmt.Println("=== SISTEMA DE CALIFICACIONES ===")
	fmt.Println("")

	// 1. Usando range para recorrer la lista
	for _, alumno := range listaEstudiantes {
		suma := 0.0

		// 2. Usando for tradicional para las notas
		for i := 0; i < len(alumno.notas); i++ {
			suma += alumno.notas[i]
		}

		promedio := suma / 5.0
		totalClase += promedio

		// Switch para ver qué tan bien le fue
		var mensaje string
		switch {
		case promedio >= 9.0:
			mensaje = "Excelente"
		case promedio >= NOTA_MINIMA:
			mensaje = "Aprobado"
		default:
			mensaje = "Reprobado"
		}

		// If/Else anidados para la lógica de asistencia
		if promedio >= NOTA_MINIMA {
			aprobadosContador++
			if alumno.faltas < ASISTENCIA_OK {
				fmt.Printf("%s: %s con baja asistencia\n", alumno.nombre, mensaje)
			} else {
				if promedio > 9.8 {
					fmt.Printf("%s: %s (sigue asi )\n", alumno.nombre, mensaje)
				} else {
					fmt.Printf("%s: %s\n", alumno.nombre, mensaje)
				}
			}
		} else {
			if alumno.faltas < 50 {
				fmt.Printf("%s: %s (Casi no vino)\n", alumno.nombre, mensaje)
			} else {
				fmt.Printf("%s: %s\n", alumno.nombre, mensaje)
			}
		}

		// Checar quién es el mejor y el peor
		if promedio > mejorNota {
			mejorNota = promedio
			elMejor = alumno.nombre
		}
		if promedio < peorNota {
			peorNota = promedio
			elPeor = alumno.nombre
		}
	}

	// 3. For como si fuera un While para imprimir el final
	iter := 0
	for iter < 1 {
		fmt.Println("\n=== ESTADÍSTICAS ===")

		promedioFinal := totalClase / 10.0
		porcentaje := (float64(aprobadosContador) / 10.0) * 100

		fmt.Printf("Promedio general: %.1f\n", promedioFinal)
		fmt.Printf("Mejor calificación: %.1f (%s)\n", mejorNota, elMejor)
		fmt.Printf("Peor calificación: %.1f (%s)\n", peorNota, elPeor)
		fmt.Printf("Aprobados: %.0f%%\n", porcentaje)

		iter++
	}
}