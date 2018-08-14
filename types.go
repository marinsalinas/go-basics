package main

import "fmt"

func printDataTypes() {
	//Here are a few built-in data tyles commonly found in most Go programs.
	i := 1                                //int
	st := "Some String"                   //string
	isValid := true                       //bool (true, false)
	f := 3.1416                           //float
	as := []string{"un", "deux", "trois"} // array of string

	//Print values
	fmt.Print(i, st, isValid, f, as)
}
