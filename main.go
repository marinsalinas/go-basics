package main

import (
	"fmt"

	"github.com/marinsalinas/go-basics/model"
)

func main() {
	jumpers := model.GetList()
	for i := range jumpers {
		//Print with format.
		fmt.Printf("%+v\n", jumpers[i])
	}
}
