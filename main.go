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

	//Ways to declare for loops:
	//Print numbers from 0 to 4.
	for i := 0; i < 5; i++ { //i++ increments by 1
		fmt.Println(i)
	}

	//for loop with a single condition also prints 5
	i := 0
	isLessThanFive := true

	for isLessThanFive {
		if i >= 5 {
			isLessThanFive = false
		}
		fmt.Println(i)
		i++
	}

	//for loop with no arguments
	j := 0

	for { //most commonly used in waiting process
		if j >= 5 {
			break
		}

		fmt.Println(j)
		j++
	}

	printLangsWithSlices()

}

func printLangs() {
	//Declaring arrays with manual declaration
	//we must set a specific size.
	var langs [3]string //Cannot hold more than 3 values of type string.

	//Remember: index count starts at 0.
	langs[0] = "Spanish"
	langs[1] = "French"
	langs[2] = "English"
	//langs[3] = "Other Lang" raises an out of bound error.

	fmt.Println(langs)
}

func printLangsWithSlices() {
	//Leaving out max elements creates a slice with a zero value of nil.
	var langs []string

	//Empty slice can be appended to using the built-in function
	//append(), which takes two arguments: a slice and
	//a variable number of elements.
	langs = append(langs, "Spanish")
	langs = append(langs, "French")
	langs = append(langs, "English")

	fmt.Println(langs)

}

func getLangs() []string {
	//Slice Literals (Array with initial values)
	langs := []string{"Spanish", "French", "English"}
	return langs
}

func printLangsWithSliceLiterals() {
	langs := getLangs()
	fmt.Println(langs)
}
