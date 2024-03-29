// client
package main

import (
	"fmt"
	"log"
	"math/rand"

	"example.com/greetings"
)

func main() {
	// Set properties of the predefined Logger, including
	// the log entry prefix and a flag to disable printing
	// the time, source file, and line number.
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	// A slice of names.
	names := []string{
		"Emma",
		"Lucas",
		"Olivia",
		"Ethan",
		"Ava",
		"Benjamin",
		"Sophia",
		"Jackson",
		"Mia",
		"Liam",
	}

	// generate a random slice of names.
	namesSlice := names[:rand.Intn(len(names))]

	// Request greeting messages for the names.
	messages, err := greetings.Hellos(namesSlice)

	// If an error was returned, print it to the console and
	// exit the program.
	if err != nil {
		log.Fatal(err)
	}

	// If no error was returned, print the returned map of
	// messages to the console.
	for _, message := range messages {
		fmt.Println(message)
	}
}
