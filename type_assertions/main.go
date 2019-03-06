package main

import (
	"./fromBool"
	"./fromString"
	"fmt"
)

func main() {
	b := true
	str := fromBool.ToString(b)
	fmt.Printf("%s (%T)\n", str, str)

	var i int = fromBool.ToInt(b)
	fmt.Printf("%d (%T)\n", i, i)
	var i8 int8 = fromBool.ToInt8(b)
	fmt.Printf("%d (%T)\n", i8, i8)
	var i16 int16 = fromBool.ToInt16(b)
	fmt.Printf("%d (%T)\n", i16, i16)
	var i32 int32 = fromBool.ToInt32(b)
	fmt.Printf("%d (%T)\n", i32, i32)
	var i64 int64 = fromBool.ToInt64(b)
	fmt.Printf("%d (%T)\n", i64, i64)

	var ui uint = fromBool.ToUint(b)
	fmt.Printf("%d (%T)\n", ui, ui)
	var ui8 uint8 = fromBool.ToUint8(b)
	fmt.Printf("%d (%T)\n", ui8, ui8)
	var ui16 uint16 = fromBool.ToUint16(b)
	fmt.Printf("%d (%T)\n", ui16, ui16)
	var ui32 uint32 = fromBool.ToUint32(b)
	fmt.Printf("%d (%T)\n", ui32, ui32)
	var ui64 uint64 = fromBool.ToUint64(b)
	fmt.Printf("%d (%T)\n", ui64, ui64)
	var uiptr uintptr = fromBool.ToUintptr(b)
	fmt.Printf("%d (%T)\n", uiptr, uiptr)

	var bt byte = fromBool.ToByte(b)
	fmt.Printf("%d (%T)\n", bt, bt)

	var rn rune = fromBool.ToRune(b)
	fmt.Printf("%d (%T)\n", rn, rn)

	var f32 float32 = fromBool.ToFloat32(b)
	fmt.Printf("%g (%T)\n", f32, f32)
	var f64 float64 = fromBool.ToFloat64(b)
	fmt.Printf("%g (%T)\n", f64, f64)

	var s = "true"
	bl, _ := fromString.ToBool(s)
	fmt.Printf("%t (%T)\n", bl, bl)
}
