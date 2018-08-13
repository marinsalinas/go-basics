/*A package definition is always
the very first thing in a Go source file.
*/
package main

//Package used by the program must be explicity
//imported after the package definition
import "fmt"

//The main() function is the entry point for all
//Go programs. It must have this name.
func main() {
	fmt.Println("Hello world 🐹 !")
}
