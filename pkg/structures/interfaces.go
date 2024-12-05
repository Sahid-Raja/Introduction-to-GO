package structures

import (
	"fmt"
	"math"
)

type geometry interface {
	iArea() float64
	iPrimeter() float64
}

type circle struct {
	radius float64
}

func (c *circle) iArea() float64 {
	return math.Pi * (math.Pow(float64(c.radius), 2))
}

func (c *circle) iPrimeter() float64 {
	return 2 * math.Pi * c.radius
}

func (r *rectangle) iPrimeter() float64 {
	return 2 * (float64(r.breadth) + float64(r.length))
}

func Interfaces() {
	c := circle{5}
	r := rectangle{10, 20}
	fmt.Printf("Area is %v \n", c.iArea())
	fmt.Printf("Perimeter is %v \n", c.iPrimeter())
	fmt.Printf("Area of Rectange is %v \n", r.area())
	fmt.Printf("Length of Rectange is %v \n", r.iPrimeter())
	fmt.Print(r)
}
