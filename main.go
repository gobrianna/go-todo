package main

import "fmt"

// Main initializes to-do list and prints it to the console
func main() {
	fmt.Println("#### Todo List ####")

		var (
			taskOne = "Look at all these chickens!"
			taskTwo = "Learn basics of Go!"
			taskThree = "Build a mini-project using Go basics!"
			taskFour = "Water the plants."
			taskFive = "Play OoT."
			// taskList is a slice of strings that holds the to-do items; can only hold elements of the same type
			taskList = []string {taskOne, taskTwo, taskThree, taskFour, taskFive}
		)
		printTasks(taskList)
}

// Iterates through taskList, printing each task with its position (starting from 1)
// Expecting a slice of strings named 'taskList' as a parameter
func printTasks(taskList []string) {
		for index, task := range taskList {
			fmt.Printf("%d. %s\n", index+1, task)
		}
}
