package structures

import "fmt"

type student struct {
	name, faculty_no string
	age              int
}

var newAge int = 0

func createStudent(name string) *student {
	newAge++
	s := student{name: name, age: newAge}

	return &s
}

func Structures() {
	fmt.Println(student{"Sahid Raja Ansari", "20COB335", 23})
	fmt.Println(student{name: "Babbar Singh", age: 23, faculty_no: "23BH335"})
	fmt.Println(student{name: "Khan Sahab"})
	fmt.Println(&student{name: "Ann", age: 40})
	fmt.Println(createStudent("John Snow"))

	newStudnet := createStudent("Danerys")
	fmt.Printf("%v ", newStudnet)

	var num = 1
	var ptr *int = &num
	println(ptr)

	college := struct {
		name string
		id   int
	}{
		"Aligarh Muslim University",
		12,
	}

	fmt.Println(college)
}
