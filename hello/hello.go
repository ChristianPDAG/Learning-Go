package main
//
import "fmt"
// importar paquete para usar la función Go() que imprime un mensaje
import "rsc.io/quote"

func main() {
    fmt.Println("Hello, World!")
    fmt.Println(quote.Go()) // Imprime un mensaje de Go de la librería rsc.io/quote

}


