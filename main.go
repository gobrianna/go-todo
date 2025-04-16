package main

import "fmt"

// Displaying usage of global variables
var (
taskOne = "Look at all these chickens!"
taskTwo = "Learn basics of Go!"
taskThree = "Build a mini-project using Go basics!"
taskFour = "Water the plants."
taskFive = "Play OoT."
)
// taskList is a slice of strings that holds the to-do items; can only hold elements of the same type
var taskList = []string {taskOne, taskTwo, taskThree, taskFour, taskFive}

// Main initializes a to-do list and prints it to the console
func main() {
	fmt.Println("#### Todo List ####")
	printTasks()
}

// Iterates through taskList, printing each task with its position (starting from 1)
func printTasks() {
		for index, task := range taskList {
			fmt.Printf("%d. %s\n", index+1, task)
		}
}
