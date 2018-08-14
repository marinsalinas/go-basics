package main

import (
	"fmt"
	"os"
	"time"

	"github.com/marinsalinas/go-basics/greet"
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

	//Add Greeting
	hourOfDay := time.Now().Hour()
	greeting, err := greet.GetGreeting(hourOfDay)

	//if error is not nil, then some error must have ocurred!
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1) //Exit program with errors
	}

	fmt.Println(greeting)
}
