package main
//
import (
    "fmt"
    "log"
    "example.com/greetings" // importar el paquete greetings
)
// importar paquete para usar la función Go() que imprime un mensaje
import "rsc.io/quote"

func main() {
    fmt.Println("Hello, World!")
    fmt.Println(quote.Go()) // Imprime un mensaje de Go de la librería rsc.io/quote

    //message:= greetings.Hello("John") // Llamar a la función Hello del paquete greetings
    //fmt.Println(message) // Imprimir el mensaje devuelto por la función Hello
    log.SetPrefix("greetings: ")
    log.SetFlags(0)

    // Request a greeting message.
    message, err := greetings.Hello("Gladys")
    // If an error was returned, print it to the console and
    // exit the program.
    if err != nil {
        log.Fatal(err)
    }

    // If no error was returned, print the returned message
    // to the console.
    fmt.Println(message)
}


