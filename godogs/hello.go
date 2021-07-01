package main

import "fmt"

// Hello returns a greeting for the named person.
func hello(name string) string {
    message := fmt.Sprintf("Hi, %v. Welcome!", name)
    return message
}