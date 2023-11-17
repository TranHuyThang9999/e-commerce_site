package mapper

import (
	"fmt"
	"regexp"
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

// SplitStringToInt64Slice chuyển đổi chuỗi có định dạng "[id, id]" thành slice int64.
func SplitStringToInt64Slice(input string) ([]int64, error) {
	// Sử dụng biểu thức chính quy để trích xuất các số từ chuỗi.
	re := regexp.MustCompile(`\d+`)
	matches := re.FindAllString(input, -1)

	// Chuyển đổi các chuỗi số thành int64.
	result := make([]int64, len(matches))
	for i, match := range matches {
		val, err := strconv.ParseInt(match, 10, 64)
		if err != nil {
			return nil, fmt.Errorf("không thể chuyển đổi chuỗi thành số nguyên: %v", err)
		}
		result[i] = val
	}

	return result, nil
}
