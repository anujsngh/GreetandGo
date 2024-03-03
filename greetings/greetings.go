// server
package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

// Hello returns a greeting for the named person.
func Hello(name string) (string, error) {
	// If no name was given, return an error with a message.
	if name == "" {
		return name, errors.New("empty name")
	}

	// Create a message using a random format.
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

// Hellos returns a hashmap {name: greeting} for each name.
func Hellos(names []string) (map[string]string, error) {
	// A map to associate names with messages.
	messages := make(map[string]string)

	// Loop through the received slice of names, calling
	// the Hello function to get a message for each name.
	for _, name := range names {
		message, err := Hello(name)

		if err != nil {
			return nil, err
		}

		// In the map, associate the retrieved message with the name.
		messages[name] = message
	}

	return messages, nil
}

// randomFormat returns one of a set of greeting messages. The returned
// message is selected at random.
func randomFormat() string {
	// A slice of message formats.
	formats := []string{
		"Hey %v, wishing you a splendid day ahead!",
		"Hi there, %v! Hope your day is filled with joy and laughter.",
		"Hello, %v! Sending you warm greetings and positive vibes.",
		"Greetings, %v! May your day be as wonderful as you are.",
		"Hey %v, just dropping by to say hello and spread some cheer!",
		"Hi %v! Here's a friendly wave to brighten your day.",
		"Hello, %v! Wishing you a day filled with smiles and good times.",
		"Hey there, %v! Hope today treats you exceptionally well.",
		"Hi %v, just wanted to send a quick hello and wish you happiness.",
		"Hello, %v! May your day be filled with all the good things life has to offer.",
	}

	// Return a randomly selected message format by specifying
	// a random index for the slice of formats.
	return formats[rand.Intn(len(formats))]
}
