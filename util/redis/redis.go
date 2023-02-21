package util_redis

import "strconv"

func ParseCount(s string) int64 {
	if s != "" {
		count, _ := strconv.ParseInt(s, 10, 0)
		return count
	}
	return 0
}
