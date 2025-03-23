package main

import "fmt"

// Main initializes a to-do list and prints it to the console
func main() {
	taskOne := "Look at all these chickens!"
	taskTwo := "Learn basics of Go!"
	taskThree := "Build a mini-project using Go basics!"
	taskFour := "Water the plants."
	taskFive := "Play OoT."

	// taskList is a slice of strings that holds the to-do items; can only hold elements of the same type
	taskList := []string {taskOne, taskTwo, taskThree, taskFour, taskFive}

	fmt.Println("#### Todo List ####")
	// Prints taskList to the console
	fmt.Println("Here's whacha gotta do:\n", taskList)
}
