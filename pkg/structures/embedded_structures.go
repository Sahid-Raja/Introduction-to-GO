package structures

import "fmt"

type name struct {
	firstName, middleName, lastName string
}

func (n *name) getName() string {
	return fmt.Sprintf("%v %v %v", n.firstName, n.middleName, n.lastName)
}

type person struct {
	name
	age int
}

func EmbeddedStructures() {
	p := person{
		name: name{
			firstName:  "Sahid",
			middleName: "Raja",
			lastName:   "Ansari",
		},
		age: 10,
	}
	fmt.Println(p)
	fmt.Printf("Hi @%v you are currently %v\n", p.lastName, p.age)
	fmt.Printf("Hi %v!!! you are currently %v\n", p.name.firstName, p.age)
	fmt.Println(p.getName())

	//TODO:
	type details interface {
		getName() string
	}

	var d details = &p
	fmt.Println(d.getName())
}
