package mapper

import (
	"strconv"
	"strings"
)

// JoinInt64SliceToString chuyển đổi slice int64 thành chuỗi với định dạng [id, id].
func JoinInt64SliceToString(slice []int64) string {
	strSlice := make([]string, 0, len(slice))
	for _, v := range slice {
		strSlice = append(strSlice, strconv.FormatInt(v, 10))
	}
	return "[" + strings.Join(strSlice, ", ") + "]"
}
