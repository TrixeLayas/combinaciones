Sistema de Combinaciones de Productos en Go
Descripción

Este programa desarrollado en Go permite generar todas las posibles combinaciones de productos de una cafetería sin exceder un presupuesto ingresado por el usuario.

El sistema utiliza el algoritmo de Backtracking para explorar diferentes combinaciones de productos del catálogo, evitando repeticiones y optimizando la búsqueda mediante poda (pruning).

Características
Catálogo precargado de productos.
Búsqueda automática de combinaciones válidas.
Validación de presupuesto ingresado.
Agrupación de resultados por cantidad de productos.
Cálculo de:
Total gastado.
Cambio restante.
Mejor combinación posible.
Uso de recursividad y backtracking.
Tecnologías Utilizadas
Lenguaje: Go
Paradigma: Programación estructurada y recursiva
Algoritmo: Backtracking
Estructura del Proyecto
 proyecto/
 ├── main.go
 └── README.md
Explicación del Programa
Estructura Producto

Representa cada artículo del catálogo.

type Producto struct {
	Nombre      string
	Precio      float64
	Categoria   string
	Descripcion string
}
Catálogo de Productos

El arreglo catalogo contiene todos los productos disponibles de la cafetería.

Ejemplos:

Café Americano
Pizza personal
Combo del día
Ensalada César
Algoritmo Backtracking

La función principal del sistema es:

func encontrarCombinaciones(productos []Producto, presupuesto float64) [][]Producto

Esta función:

Explora todas las combinaciones posibles.
Evita repetir productos.
Descarta rutas que exceden el presupuesto.
Guarda únicamente combinaciones válidas.
Conceptos Implementados
1. Recursividad

La búsqueda de combinaciones se realiza mediante llamadas recursivas:

backtrack(i+1, sumaActual+producto.Precio)
2. Poda (Pruning)

Se evita seguir explorando cuando el presupuesto sería excedido:

if sumaActual+producto.Precio <= presupuesto
3. Backtracking

Después de explorar una combinación, el producto agregado se elimina para probar nuevas posibilidades:

combinacionActual = combinacionActual[:len(combinacionActual)-1]
Ejemplo de Ejecución
Ingresa tu presupuesto: $100

Salida:

Presupuesto: $100.00
Total de combinaciones: 12

[1] 1 producto(s) — Total: $20.00 — Cambio: $80.00
     • Agua mineral        $20.00

[2] 2 producto(s) — Total: $50.00 — Cambio: $50.00
     • Agua mineral        $20.00
     • Té verde            $30.00
Validaciones

El programa valida que:

El usuario ingrese un número válido.
El presupuesto sea mayor a 0.

Si ocurre un error:

Error: Por favor ingresa un presupuesto válido mayor a 0.
Cómo Ejecutar el Proyecto
1. Instalar Go

Descargar Go desde:

Go Programming Language

2. Ejecutar el programa

En la terminal:

go run main.go
