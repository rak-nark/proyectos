package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	for {
		mostrarMenuFor()
		opcion := leerOpcion()

		switch opcion {
		case 1:
			ejercicioFor1() // Contador 1-5 (como while)
		case 2:
			ejercicioFor2() // Sumar hasta ingresar 0 (como do-while)
		case 3:
			ejercicioFor3() // Buscar número aleatorio
		case 0:
			fmt.Println("¡Hasta luego! 👋")
			os.Exit(0)
		default:
			fmt.Println("❌ Opción no válida. Intenta de nuevo.")
		}
	}
}

// --- Menú y funciones auxiliares ---
func mostrarMenuFor() {
	fmt.Println("\n=== EJERCICIOS DE FOR ===")
	fmt.Println("1. Contador del 1 al 5 (while-style)")
	fmt.Println("2. Sumar números hasta ingresar 0 (do-while-style)")
	fmt.Println("3. Buscar número aleatorio 7 (break)")
	fmt.Println("0. Salir")
	fmt.Print("Elige una opción: ")
}

func leerOpcion() int {
	var opcion int
	fmt.Scanln(&opcion)
	return opcion
}

// --- Ejercicios con FOR ---

// Ejercicio 1: Contador del 1 al 5 (simula "while")
func ejercicioFor1() {
	fmt.Println("\n[FOR1] Contador del 1 al 5 (while-style)")
	i := 1
	for i <= 5 { // Equivalente a while(i <= 5)
		fmt.Println(i)
		i++
	}
}

// Ejercicio 2: Sumar números hasta que se ingrese 0 (simula "do-while")
func ejercicioFor2() {
	fmt.Println("\n[FOR2] Sumar números (do-while-style)")
	suma := 0
	var numero int

	for {
		fmt.Print("Ingresa un número (0 para terminar): ")
		fmt.Scanln(&numero)
		if numero == 0 {
			break
		}
		suma += numero
	}
	fmt.Printf("La suma total es: %d\n", suma)
}

// Ejercicio 3: Buscar el número 7 en valores aleatorios (uso de "break")
func ejercicioFor3() {
	fmt.Println("\n[FOR3] Buscando el número 7...")
	rand.Seed(time.Now().UnixNano()) // Inicializa el generador de números aleatorios

	for {
		num := rand.Intn(10) // Genera un número entre 0 y 9
		fmt.Println("Número generado:", num)
		if num == 7 {
			fmt.Println("¡Encontrado el 7! 🎉")
			break
		}
		time.Sleep(500 * time.Millisecond) // Pequeña pausa para dramatismo
	}
}