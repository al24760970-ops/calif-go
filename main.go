package main

import "fmt"

func main() {
	// Variables para los datos
	var alumno string
	var totalNotas int
	var suma float64

	fmt.Println("SISTEMA DE NOTAS V1")
	fmt.Print("Nombre del compa: ")
	fmt.Scan(&alumno)

	fmt.Print("Cuantas materias lleva?: ")
	fmt.Scan(&totalNotas)

	// El arreglo para guardar las notas
	notas := make([]float64, totalNotas)

	// Pedir notas con un for
	for i := 0; i < totalNotas; i++ {
		fmt.Printf("Dime la nota %d: ", i+1)
		fmt.Scan(&notas[i])
		suma = suma + notas[i] // Sumando
	}

	prom := suma / float64(totalNotas)

	fmt.Println("-----------------------")
	fmt.Println("Resultado de:", alumno)
	fmt.Printf("Salió con: %.1f\n", prom)

	// Ver si pasó o no (If/Else)
	if prom >= 7.0 {
		fmt.Println("Pasaste")
	} else if prom >= 6.0 && prom < 7.0 {
		fmt.Println("A extraordinario")
	} else {
		fmt.Println("F por ti, reprobado")
	}

	// El switch para el mensaje final
	switch {
	case prom >= 9.0:
		fmt.Println("felicidades pasaste")
	case prom >= 7.0:
		fmt.Println("Pasaste de panzazo")
	default:
		fmt.Println("A estudiar mas")
	}
}

