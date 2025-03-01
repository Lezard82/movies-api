package utils

import "strconv"

func ParseID(idStr string) (int64, error) {
	return strconv.ParseInt(idStr, 10, 64)
}
