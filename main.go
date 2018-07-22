package main

// Import the fmt for formatting strings
// Import os so we can read environment variables from the system
import (
	"fmt"
	"os"
)

// create a function named getUserName that returns a single string
func getUserName() string {
	// On Linux and Mac the username is available as simply USER
	username := os.Getenv("USER")
	// check if we got an empty string and attempt to get USERNAME instead which would be available on Windows
	if username == "" {
		username = os.Getenv("USERNAME")
	}
	// Return whatever we received
	return username
}

// entrypoint function for our app
func main() {
	fmt.Printf("Hello, %s\n", getUserName())
}
