package main

import (
	"./bool"
	"fmt"
)

func main() {
	b := true
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

	var ui uint = bool.ToUint(b)
	fmt.Printf("%d (%T)\n", ui, ui)
	var ui8 uint8 = bool.ToUint8(b)
	fmt.Printf("%d (%T)\n", ui8, ui8)
	var ui16 uint16 = bool.ToUint16(b)
	fmt.Printf("%d (%T)\n", ui16, ui16)
	var ui32 uint32 = bool.ToUint32(b)
	fmt.Printf("%d (%T)\n", ui32, ui32)
	var ui64 uint64 = bool.ToUint64(b)
	fmt.Printf("%d (%T)\n", ui64, ui64)
	var uiptr uintptr = bool.ToUintptr(b)
	fmt.Printf("%d (%T)\n", uiptr, uiptr)

	var bt byte = bool.ToByte(b)
	fmt.Printf("%d (%T)\n", bt, bt)

	var rn rune = bool.ToRune(b)
	fmt.Printf("%d (%T)\n", rn, rn)

	var f32 float32 = bool.ToFloat32(b)
	fmt.Printf("%g (%T)\n", f32, f32)
	var f64 float64 = bool.ToFloat64(b)
	fmt.Printf("%g (%T)\n", f64, f64)
}
