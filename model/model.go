package model

type jumper interface {
	//Method expected to be present in all
	//types that implement this interface
	Jump() string //Capitalized alows us to use the method outside the package.
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
func (g gopher) Jump() string {
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

func (h horse) Jump() string {
	if h.weight > 2500 {
		return "It's too heavy, can't jump"
	}
	return "I will jump!"
}

//GetList :[]jumper can be used as the return type,
//and may contain gopher and jumper elements within
func GetList() []jumper {
	phil := &gopher{name: "Phil", age: 30}
	noodles := &gopher{name: "Noodles", age: 90}
	gil := &horse{name: "Gil", weight: 2400.40}
	return []jumper{phil, noodles, gil}
}

//passing structs by reference
//g is a memory reference of the original struct data.
// * indicates a pointer to the original value.
func ValidateAge(g *gopher) {
	//assings true to the copy of the data.
	g.isAdult = g.age >= 18
}
