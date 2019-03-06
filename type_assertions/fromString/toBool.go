package fromString

import (
	"strconv"
)

func ToBool(str string) (b bool, err error) {
	return strconv.ParseBool(str)
}