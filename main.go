package main

import (
	"fmt"
	"os"
)

func main() {
	//An array with the program arguments, starting
	//with the name fo the executable file and
	//followed by any user-supplied arguments.
	if len(os.Args) > 1 {
		fmt.Println(os.Args[1])
	} else {
		//No argument passed.
		fmt.Println("Hello world 🐹 !")
	}
}

//run the program with:
//$go run main.go "Some String"
