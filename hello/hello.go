package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
	// Set properties of the predefined Logger, including
	// the log entry prefix and a flag to disable printing
	// the time, source file, and line number.
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	// Request a greeting message
	message, err := greetings.Hello("Jonny")
	// If an error was returned, print it to the console and
	// exit the program.
	validateError(err)

	// If no error was returned, print the returned message
	// to the console.
	fmt.Println(message)

	names := []string{"Jake", "Jeff", "Jerry"}
	messages, err := greetings.Hellos(names)

	// Similar to before, check that errs is not nil & handle if so.
	validateError(err)

	// For each name & message pair, print the name and message.
	for name, message := range messages {
		fmt.Println("name:", name, "\nmessage:", message)
	}

	// Or can just print map directly too.
	fmt.Println(messages)
}

func validateError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
