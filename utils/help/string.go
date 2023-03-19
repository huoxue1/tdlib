package help

import (
	"fmt"
	"strconv"
	"unsafe"
)

func ToString(value any) string {

	if _, ok := value.(fmt.Stringer); ok {
		return value.(fmt.Stringer).String()
	}
	switch value.(type) {
	case int:
		return strconv.Itoa(value.(int))
	case int64:
		return strconv.FormatInt(value.(int64), 10)
	case float64:
		return strconv.FormatFloat(value.(float64), 'E', -1, 64)
	case string:
		return value.(string)
	case []byte:
		return ByteSliceToString(value.([]byte))
	default:
		return ""
	}
}

func StringToByteSlice(s string) []byte {

	tmp1 := (*[2]uintptr)(unsafe.Pointer(&s))

	tmp2 := [3]uintptr{tmp1[0], tmp1[1], tmp1[1]}

	return *(*[]byte)(unsafe.Pointer(&tmp2))

}

func ByteSliceToString(bytes []byte) string {

	return *(*string)(unsafe.Pointer(&bytes))

}
