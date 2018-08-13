package main

import (
	"fmt"
	"os"
)

func main() {
	//The := operator tells Go to automatically find out the data type
	//for the value returned from os.Args.
	args := os.Args

	//Data type is declared prior to assigment.
	//Given a value, we must determine how much space is needed to store
	//this value for later reuse.
	var message string

	if len(args) > 1 {
		message = args[1]
	} else {
		//No argument passed.
		message = "Hello world 🐹 !"
	}
	fmt.Println(message)
}
