package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

var randomFormat_FN = randomFormat

func Hello(name string) (string, error) {

	if name == "" || name == " " {
		return "", errors.New("el campo name se encuentra vacio")
	}
	message := fmt.Sprintf(randomFormat_FN(), name)
	return message, nil
}

func HelloPeople(names []string) (map[string]string, error) {
	messages := make(map[string]string)
	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}

		messages[name] = message
	}

	return messages, nil
}

func randomFormat() string {
	formats := []string{
		"Hola %v seas bienvenido",
		"%v que la fuerza te acompañe",
		"Saludos %v",
	}

	return formats[rand.Intn(len(formats))]
}
