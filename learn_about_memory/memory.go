package main

import (
	"fmt"
)

// Stack Se usa para almacenar variables locales y seguir la ejecución de tareas
// Heap Se usa para almacenar variables globales y objetos que se crean dinámicamente

func maain() {
	a:= 10 // variable a de tipo int, se almacena en la pila (STACK)
	b:= "A" // variable b de tipo string, se almacena en la pila (STACK)
	c:= 3.14 // variable c de tipo float64, se almacena en la pila (STACK)
	d:= true // variable d de tipo bool, se almacena en la pila (STACK)
	
	fmt.Println("a:", a) // Imprime el valor de a
	fmt.Println("b:", b, "c:",c, "d:", d) // Imprime el valor de b

	p:= createPointer() // Llama a la función createPointer y almacena el puntero en p
	fmt.Println("p:", *p ,"devuelve la variable del puntero a pesar que sea local de la función") // Imprime el valor apuntado por p (valor de a), se accede al valor usando el operador de desreferencia (*)

	s := make([]int, 5) // Crea un slice de enteros de longitud 5, se almacena en el heap (HEAP)
	fmt.Println(s)
}

func createPointer() *int {
	a:=42
	return &a // Devuelve la dirección de memoria de a, que se almacena en el heap (HEAP)
}
