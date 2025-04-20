package main

import (
	"fmt"
	"os"
)

func main() {
	for {
		mostrarMenuIf()
		opcion := leerOpcion()

		switch opcion {
		case 1:
			ejercicioIf1() // Número positivo (con input)
		case 2:
			ejercicioIf2() // Comparar números (con input)
		case 3:
			ejercicioIf3() // Validar contraseña (ya tenía input)
		case 0:
			fmt.Println("¡Hasta luego! 👋")
			os.Exit(0)
		default:
			fmt.Println("❌ Opción no válida. Intenta de nuevo.")
		}
	}
}

// --- Menú y funciones auxiliares ---
func mostrarMenuIf() {
	fmt.Println("\n=== EJERCICIOS DE IF ===")
	fmt.Println("1. Verificar número positivo")
	fmt.Println("2. Comparar dos números")
	fmt.Println("3. Validar contraseña")
	fmt.Println("0. Salir")
	fmt.Print("Elige una opción: ")
}

func leerOpcion() int {
	var opcion int
	fmt.Scanln(&opcion)
	return opcion
}

// --- Ejercicios con input del usuario ---

// Ejercicio 1: Verificar si un número es positivo (con Scanln)
func ejercicioIf1() {
	fmt.Println("\n[IF1] Verificar número positivo")
	var numero int
	fmt.Print("Ingresa un número: ")
	fmt.Scanln(&numero)

	if numero > 0 {
		fmt.Println("El número es positivo.")
	} else if numero < 0 {
		fmt.Println("El número es negativo.")
	} else {
		fmt.Println("El número es cero.")
	}
}

// Ejercicio 2: Comparar dos números (con Scanln)
func ejercicioIf2() {
	fmt.Println("\n[IF2] Comparar dos números")
	var a, b int
	fmt.Print("Ingresa el primer número (a): ")
	fmt.Scanln(&a)
	fmt.Print("Ingresa el segundo número (b): ")
	fmt.Scanln(&b)

	if a > b {
		fmt.Printf("%d es mayor que %d\n", a, b)
	} else if a < b {
		fmt.Printf("%d es menor que %d\n", a, b)
	} else {
		fmt.Println("Ambos números son iguales.")
	}
}

// Ejercicio 3: Validar contraseña (ya interactivo)
func ejercicioIf3() {
	fmt.Println("\n[IF3] Validar contraseña")
	passwordCorrecta := "golang123" // Puedes cambiarla
	var input string

	fmt.Print("Ingresa la contraseña: ")
	fmt.Scanln(&input)

	if input == passwordCorrecta {
		fmt.Println("✅ Acceso concedido")
	} else {
		fmt.Println("❌ Acceso denegado")
	}
}