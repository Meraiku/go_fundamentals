package tasks

import "fmt"

func GreetPerson(name string) string {
	if name == "" {
		name = "World"
	}

	return fmt.Sprintf("Hello, %s!", name)
}
