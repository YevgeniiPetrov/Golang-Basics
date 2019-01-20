package main

import "fmt"

func main() {
	shapes := make([]Shape, 0, 10)

	for i := 1; i < 6; i++ {
		j := float64(i)
		shapes = append(shapes, newRectangle(j, j+1))
		shapes = append(shapes, newTriangle(j, j+1, j+2))
	}

	abooutShapes(shapes)
}

func abooutShapes(shapes []Shape) {
	for _, shape := range shapes {
		fmt.Printf("%v: Perimeter = %f, Squere = %f\n",
			shape.Title(), shape.Perimeter(), shape.Squere())
	}
}
