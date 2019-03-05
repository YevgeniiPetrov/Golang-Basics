package bool

import "strconv"

func ToString(b bool) string {
	return strconv.FormatBool(b)
}