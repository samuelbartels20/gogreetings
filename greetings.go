package gogreetings

import "errors"

func Hello(name string) (string, error) {

	// If no name was given, return an error with a message.
	if name == "" {
		return "", errors.New("empty name")
	}
	// Return a greeting that embeds the name in a message.
	message := "Hi, " + name + ". Welcome!"
	return message, nil
}
