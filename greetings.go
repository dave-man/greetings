package greetings

import (
	"errors"
	"fmt"
)

// Hello is a fucking greeting function
func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("empty name")
	}

	message := fmt.Sprintf("hi %v", name)
	return message, nil
}
