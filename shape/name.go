package main

import "strconv"

func createName(shape string) func() string {
	number := 0
	return func() string {
		number++
		return shape + " " + strconv.Itoa(number)
	}
}
