// Marco Antonio Vizcarra Valle  #25760062
package main

import (
	"fmt"
)

// Producto representa un artículo del catálogo de la cafetería.
type Producto struct {
	Nombre      string
	Precio      float64
	Categoria   string
	Descripcion string
}

// catalogo contiene todos los productos disponibles.
// No modifiques esta variable.
var catalogo = []Producto{
	{Nombre: "Agua mineral", Precio: 20.00, Categoria: "Bebida", Descripcion: "Botella 600ml sin gas"},
	{Nombre: "Té verde", Precio: 30.00, Categoria: "Bebida", Descripcion: "Infusión caliente"},
	{Nombre: "Café Americano", Precio: 35.00, Categoria: "Bebida", Descripcion: "Café negro doble shot"},
	{Nombre: "Jugo de naranja", Precio: 45.00, Categoria: "Bebida", Descripcion: "Natural exprimido 355ml"},
	{Nombre: "Pastel de chocolate", Precio: 55.00, Categoria: "Postre", Descripcion: "Rebanada individual 120g"},
	{Nombre: "Burrito de res", Precio: 75.00, Categoria: "Comida", Descripcion: "Tortilla, carne, frijoles"},
	{Nombre: "Sandwich de pollo", Precio: 85.00, Categoria: "Comida", Descripcion: "Pan integral, pollo, verduras"},
	{Nombre: "Ensalada César", Precio: 95.00, Categoria: "Comida", Descripcion: "Lechuga romana, crutones"},
	{Nombre: "Pizza personal", Precio: 110.00, Categoria: "Comida", Descripcion: "4 rebanadas, queso y jitomate"},
	{Nombre: "Combo del día", Precio: 130.00, Categoria: "Combo", Descripcion: "Plato fuerte + bebida + postre"},
}

// encontrarCombinaciones recibe el catálogo y un presupuesto y
// devuelve todas las combinaciones de productos distintos que
// se pueden comprar sin exceder ese monto.
// Cada producto puede aparecer solo una vez por combinación.
func encontrarCombinaciones(productos []Producto, presupuesto float64) [][]Producto {
	var resultados [][]Producto
	var combinacionActual []Producto

	// Función recursiva para buscar combinaciones (Backtracking)
	var backtrack func(inicio int, sumaActual float64)
	backtrack = func(inicio int, sumaActual float64) {
		// Comentario 1: Guardamos una copia de la combinación válida actual para no perderla por referencia
		// cuando combinacionActual se modifique en las siguientes iteraciones.
		if len(combinacionActual) > 0 {
			copia := make([]Producto, len(combinacionActual))
			copy(copia, combinacionActual)
			resultados = append(resultados, copia)
		}

		// Comentario 2: El ciclo inicia desde el índice 'inicio' para evitar permutaciones duplicadas
		// y asegurar que no se repitan productos ya evaluados en la combinación actual.
		for i := inicio; i < len(productos); i++ {
			producto := productos[i]

			// Comentario 3: Poda (Pruning) - Si al sumar el precio del producto no excedemos el presupuesto,
			// lo agregamos temporalmente y exploramos esa rama llamando a backtrack recursivamente.
			if sumaActual+producto.Precio <= presupuesto {
				combinacionActual = append(combinacionActual, producto)
				backtrack(i+1, sumaActual+producto.Precio)
				// Backtracking: eliminamos el último producto añadido para intentar con el siguiente en el ciclo.
				combinacionActual = combinacionActual[:len(combinacionActual)-1]
			}
		}
	}

	// Iniciamos la búsqueda desde el índice 0 y una suma acumulada de 0.0
	backtrack(0, 0.0)
	return resultados
}

// imprimirResultados muestra en consola el resumen de combinaciones.
func imprimirResultados(combis [][]Producto, presupuesto float64) {
	fmt.Printf("\nPresupuesto: $%.2f\n", presupuesto)
	fmt.Printf("Total de combinaciones: %d\n\n", len(combis))

	if len(combis) == 0 {
		fmt.Println("No se encontró ninguna combinación para este presupuesto.")
		return
	}

	// Agrupación por cantidad de productos
	fmt.Println("Por cantidad de productos:")
	agrupacion := make(map[int]int)
	var maxProductos int
	for _, combi := range combis {
		cant := len(combi)
		agrupacion[cant]++
		if cant > maxProductos {
			maxProductos = cant
		}
	}

	for i := 1; i <= maxProductos; i++ {
		if count, exists := agrupacion[i]; exists {
			if count == 1 {
				fmt.Printf("  %d producto(s): %d combinación\n", i, count)
			} else {
				fmt.Printf("  %d producto(s): %d combinaciones\n", i, count)
			}
		}
	}
	fmt.Println()

	var mejorCombinacion []Producto
	var mayorGasto float64 = -1.0

	// Detalle de cada combinación
	for i, combi := range combis {
		var total float64
		for _, p := range combi {
			total += p.Precio
		}
		cambio := presupuesto - total

		// Actualizamos la combinación que representa el mayor gasto
		if total > mayorGasto {
			mayorGasto = total
			mejorCombinacion = combi
		}

		fmt.Printf("[%d] %d producto(s) — Total: $%.2f — Cambio: $%.2f\n", i+1, len(combi), total, cambio)
		for _, p := range combi {
			fmt.Printf("     • %-20s $%.2f\n", p.Nombre, p.Precio)
		}
		fmt.Println()
	}

	// La combinación de mayor valor
	if len(mejorCombinacion) > 0 {
		fmt.Println("Mejor combinación (mayor gasto):")
		for _, p := range mejorCombinacion {
			fmt.Printf("     • %-20s $%.2f\n", p.Nombre, p.Precio)
		}
		fmt.Printf("     Total: $%.2f  Cambio: $%.2f\n", mayorGasto, presupuesto-mayorGasto)
	}
}

func main() {
	var presupuesto float64
	fmt.Print("Ingresa tu presupuesto: $")

	// Validar la entrada del usuario
	_, err := fmt.Scan(&presupuesto)
	if err != nil || presupuesto <= 0 {
		fmt.Println("\nError: Por favor ingresa un presupuesto válido mayor a 0.")
		return
	}

	// Llamar a encontrarCombinaciones
	combis := encontrarCombinaciones(catalogo, presupuesto)

	// Llamar a imprimirResultados
	imprimirResultados(combis, presupuesto)
}
