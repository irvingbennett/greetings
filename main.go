// Package greetings returns a sting
package greetings

import "fmt"

// Hello returns a for the named person
func Hello(name string) string {
	// Return greeting that embeds the name in a message.
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message
}
