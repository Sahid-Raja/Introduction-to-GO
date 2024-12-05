package structures

import "fmt"

type rectangle struct {
	length, breadth int
}

func (r *rectangle) area() int {
	return r.length * r.breadth
}

func (r rectangle) perimeter() int {
	return 2 * (r.length + r.breadth)
}

func Methods() {
	r := rectangle{10, 20}
	fmt.Println("Area of the rectangle is :", r.area())
	fmt.Println("Perimeter of the rectangle is :", r.perimeter())
}
