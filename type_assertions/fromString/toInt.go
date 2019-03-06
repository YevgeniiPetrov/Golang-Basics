package fromString

import (
	"strconv"
)

func ToInt(str string) (i int, err error) {
	return strconv.Atoi(str)
}

func ToInt8(str string, base int, bitSize int) (i int8, err error) {
	i64, err := ToInt64(str, base, bitSize)
	i = int8(i64)
	return
}

func ToInt16(str string, base int, bitSize int) (i int16, err error) {
	i64, err := ToInt64(str, base, bitSize)
	i = int16(i64)
	return
}

func ToInt32(str string, base int, bitSize int) (i int32, err error) {
	i64, err := ToInt64(str, base, bitSize)
	i = int32(i64)
	return
}

func ToInt64(str string, base int, bitSize int) (i int64, err error) {
	return strconv.ParseInt(str, base, bitSize)
}

func toInt64(str string, base int, bitSize int)
