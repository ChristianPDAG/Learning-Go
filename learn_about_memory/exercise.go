package main

import (
	"fmt"
	"runtime"
)

type Cliente struct {
	nombre string
}

func crearCliente(nombre string) *Cliente {
	return &Cliente{nombre}
}

func main() {
	c1 := crearCliente("Pedro")
	fmt.Println("Cliente 1:", c1.nombre)

	// Eliminamos referencia
	c1 = nil 

	// Forzamos el Garbage Collector
	runtime.GC()

	// Volvemos a acceder a c1 (esto es un test)
	fmt.Println("Cliente 1 después de GC:", c1)
}
