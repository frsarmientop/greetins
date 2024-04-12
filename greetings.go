package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

// Devuelve un saludo
func Hello(name string) (string, error) {
	var texto = "Nombre vacio"
	if name == "" {
		return "", errors.New(texto)
	}
	mensaje := fmt.Sprintf(randomFormat(), name)
	return mensaje, nil
}

func Hellos(names []string) (map[string]string, error) {
	masMensaje := make(map[string]string)

	for _, name := range names {
		mensaje, err := Hello(name)
		if err != nil {
			return nil, err
		}
		masMensaje[name] = mensaje

	}
	return masMensaje, nil
}

func randomFormat() string {
	formats := []string{
		"Hola %v Welcome",
		"Que bueno, %v",
		"Saludo %v! encantado de conocerte",
	}
	return formats[rand.Intn(len(formats))]
}
