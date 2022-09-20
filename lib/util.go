package lib

import (
	"strconv"
	"strings"
)

func IDConveryChannelID(id int64) int64 {
	i, _ := strconv.ParseInt("-100"+strconv.FormatInt(id, 10), 10, 64)
	return i
}

// ChannelIDConvertID
/* @Description: 去掉-100前缀
*  @param id
*  @return int64
 */
func ChannelIDConvertID(id int64) int64 {
	i, _ := strconv.ParseInt(strings.TrimPrefix(strconv.FormatInt(id, 10), "-100"), 10, 64)
	return i
}
