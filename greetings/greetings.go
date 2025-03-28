package greetings

import (
	"fmt"
	"errors"
	"math/rand"
)

// name: parameter || string: parameter type || string: return type
// string, error permite devolver 2 valores, string y error
func Hello(name string) (string, error) {
	// si no se ingresó nombre, devuelve mensaje de error
    if name == "" {
        return "", errors.New("empty name")
    }
	// se añadió la funcion randomFormat() para devolver un mensaje aleatorio
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

// randomFormat returns one of a set of greeting messages. The returned
// message is selected at random.
func randomFormat() string {
    // A slice of message formats.
    formats := []string{
        "Hi, %v. Welcome!",
        "Great to see you, %v!",
        "Hail, %v! Well met!",
    }

    // Return a randomly selected message format by specifying
    // a random index for the slice of formats.
    return formats[rand.Intn(len(formats))]
}

