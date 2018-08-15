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

	langs := getLangs()
	printEachLang(langs)

	gopher1 := &gopher{name: "Marin", age: 24}
	fmt.Println(gopher1.jump())
	validateAge(gopher1)
	fmt.Println(gopher1.isAdult)
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
func printLangsWithSliceLiterals() {
	langs := getLangs()
	fmt.Println(langs)
}

func getLangs() []string {
	//Slice Literals (Array with initial values)
	langs := []string{"Spanish", "French", "English"}
	return langs
}

func printEachLang(langs []string) {
	//The index for each element is returned on each run of the loop.
	for index := range langs {
		//We can now safety use the index to fetch each element from the slice.
		fmt.Printf("language[%d] = %s\n", index, langs[index])
	}
}

type jumper interface {
	//Method expected to be present in all
	//types that implement this interface
	jump() string
}

//A struct begins with the type keyword that
// indicates a new type is about to be declared.
//Then the name of the struct and ends with the 'struct' primitive type.
//Also must be placed outside the main function.
type gopher struct {
	//Properties are variables internal to the struct
	name    string
	age     int
	isAdult bool
}

//This is how we specify an explicit receiver for this function.
func (g gopher) jump() string {
	//we can access to gopher properties.
	if g.age < 65 {
		return g.name + " can jump HIGH"
	}
	return g.name + " can still jump"
}

//The horse struct implements jumper interface specifying
// jump method.
type horse struct {
	name   string
	weight float64
}

func (h horse) jump() string {
	if h.weight > 2500 {
		return "It's too heavy, can't jump"
	}
	return "I will jump!"
}

func getList() []jumper {
	phil := &gopher{name: "Phil", age: 30}
	noodles := &gopher{name: "Noodles", age: 90}
	gil := &horse{name: "Gil", weight: 2400.40}
	return []jumper{phil, noodles, gil}
}

//passing structs by reference
//g is a memory reference of the original struct data.
// * indicates a pointer to the original value.
func validateAge(g *gopher) {
	//assings true to the copy of the data.
	g.isAdult = g.age >= 18
}

// func main() {
// 	//In order to assing a struct reference to a new variable,
// 	//we use the & operator to return a pointer.
// 	marin := &gopher{name: "Marin", age: 24}
// 	validateAge(marin) // Passes a reference to the original struct.
// 	fmt.Println(marin) //Will prints age:true, it will change.
// }
