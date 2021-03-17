// Package greetings returns a sting
package greetings

import (
	"errors"
	"fmt"
)

// Hello returns a for the named person
func Hello(name string) (string, error) {
	// If no name gas given, return an error with a message.
	if name == "" {
		return "", errors.New("empty name")
	}
	// Return greeting that embeds the name in a message.
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message, nil
}
