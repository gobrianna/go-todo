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

	// for loop that iterates through taskList, returning the index from 1 and value of task
	for index, task := range taskList {
		fmt.Printf("%d. %s\n", index+1, task)
	}
}
