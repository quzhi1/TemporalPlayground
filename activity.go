package app

import (
	"fmt"
)

func ComposeGreeting(name string) (string, error) {
	greeting := fmt.Sprintf("Hello %s!", name)
	return greeting, nil
}

func ComposeGoodbye(name string) (string, error) {
	goodbye := fmt.Sprintf("Goodbye %s!", name)
	return goodbye, nil
}
