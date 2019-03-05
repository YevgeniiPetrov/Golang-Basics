package main

import (
	"./bool"
	"fmt"
)

func main() {
	b := false
	str := bool.ToString(b)
	fmt.Printf("%s (%T)\n", str, str)

	var i int = bool.ToInt(b)
	fmt.Printf("%d (%T)\n", i, i)
	var i8 int8 = bool.ToInt8(b)
	fmt.Printf("%d (%T)\n", i8, i8)
	var i16 int16 = bool.ToInt16(b)
	fmt.Printf("%d (%T)\n", i16, i16)
	var i32 int32 = bool.ToInt32(b)
	fmt.Printf("%d (%T)\n", i32, i32)
	var i64 int64 = bool.ToInt64(b)
	fmt.Printf("%d (%T)\n", i64, i64)
}
