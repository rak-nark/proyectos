package main

import (
	"fmt"
	"os"
)

func main() {
	for {
		mostrarMenuSwitch()
		opcion := leerOpcion()

		switch opcion {
		case 1:
			ejercicioSwitch1() // Días de la semana
		case 2:
			ejercicioSwitch2() // Calculadora simple
		case 3:
			ejercicioSwitch3() // Evaluación de notas
		case 0:
			fmt.Println("¡Hasta luego! 👋")
			os.Exit(0)
		default:
			fmt.Println("❌ Opción no válida. Intenta de nuevo.")
		}
	}
}

// --- Menú y funciones auxiliares ---
func mostrarMenuSwitch() {
	fmt.Println("\n=== EJERCICIOS DE SWITCH ===")
	fmt.Println("1. Días de la semana")
	fmt.Println("2. Calculadora simple")
	fmt.Println("3. Evaluación de notas")
	fmt.Println("0. Salir")
	fmt.Print("Elige una opción: ")
}

func leerOpcion() int {
	var opcion int
	fmt.Scanln(&opcion)
	return opcion
}

// --- Ejercicios con SWITCH ---

// Ejercicio 1: Días de la semana
func ejercicioSwitch1() {
	fmt.Println("\n[SWITCH1] Días de la semana")
	var dia int
	fmt.Print("Ingresa un número del 1 al 7: ")
	fmt.Scanln(&dia)

	switch dia {
	case 1:
		fmt.Println("Lunes")
	case 2:
		fmt.Println("Martes")
	case 3:
		fmt.Println("Miércoles")
	case 4:
		fmt.Println("Jueves")
	case 5:
		fmt.Println("Viernes")
	case 6:
		fmt.Println("Sábado")
	case 7:
		fmt.Println("Domingo")
	default:
		fmt.Println("Número inválido. Debe ser entre 1 y 7.")
	}
}

// Ejercicio 2: Calculadora simple
func ejercicioSwitch2() {
	fmt.Println("\n[SWITCH2] Calculadora simple")
	var (
		num1, num2 float64
		operador   string
	)

	fmt.Print("Ingresa el primer número: ")
	fmt.Scanln(&num1)
	fmt.Print("Ingresa el segundo número: ")
	fmt.Scanln(&num2)
	fmt.Print("Ingresa la operación (+, -, *, /): ")
	fmt.Scanln(&operador)

	switch operador {
	case "+":
		fmt.Printf("Resultado: %.2f\n", num1+num2)
	case "-":
		fmt.Printf("Resultado: %.2f\n", num1-num2)
	case "*":
		fmt.Printf("Resultado: %.2f\n", num1*num2)
	case "/":
		if num2 == 0 {
			fmt.Println("Error: división por cero")
		} else {
			fmt.Printf("Resultado: %.2f\n", num1/num2)
		}
	default:
		fmt.Println("Operador no válido")
	}
}

// Ejercicio 3: Evaluación de notas (rangos con switch)
func ejercicioSwitch3() {
	fmt.Println("\n[SWITCH3] Evaluación de notas")
	var nota float64
	fmt.Print("Ingresa tu nota (0-10): ")
	fmt.Scanln(&nota)

	switch {
	case nota >= 9 && nota <= 10:
		fmt.Println("Sobresaliente (A)")
	case nota >= 7 && nota < 9:
		fmt.Println("Notable (B)")
	case nota >= 5 && nota < 7:
		fmt.Println("Aprobado (C)")
	case nota >= 0 && nota < 5:
		fmt.Println("Reprobado (D)")
	default:
		fmt.Println("Nota inválida")
	}
}