package main

type Rectangle struct {
	SideA float64
	SideB float64
	Name  string
}

var rectangleName func() string

func init() {
	rectangleName = createName("Rectangle")
}

func newRectangle(sideA, sideB float64) *Rectangle {
	return &Rectangle{
		SideA: sideA,
		SideB: sideB,
		Name:  rectangleName(),
	}
}

func (r *Rectangle) Perimeter() float64 {
	return (r.SideA + r.SideB) * 2
}

func (r *Rectangle) Squere() float64 {
	return r.SideA * r.SideB
}

func (r *Rectangle) Title() string {
	return r.Name
}
