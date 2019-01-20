package main

import "math"

type Triangle struct {
	SideA float64
	SideB float64
	SideC float64
	Name  string
}

var triangleName func() string

func init() {
	triangleName = createName("Triangle")
}

func newTriangle(sideA, sideB, sideC float64) *Triangle {
	return &Triangle{
		SideA: sideA,
		SideB: sideB,
		SideC: sideC,
		Name:  triangleName(),
	}
}

func (t *Triangle) Perimeter() float64 {
	return t.SideA + t.SideB + t.SideC
}

func (t *Triangle) Squere() float64 {
	semiPerimeter := (t.SideA + t.SideB + t.SideC) / 2
	return math.Sqrt(semiPerimeter *
		(semiPerimeter - t.SideA) *
		(semiPerimeter - t.SideB) *
		(semiPerimeter - t.SideC))
}

func (r *Triangle) Title() string {
	return r.Name
}
