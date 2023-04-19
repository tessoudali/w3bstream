package stringsx

import (
	"math/rand"
	"time"
)

var (
	visibleChars    = []byte("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	visibleCharsLen = len(visibleChars)
)

func GenRandomVisibleString(length int) string {
	if length == 0 {
		length = 16
	}
	result := make([]byte, 0, length)
	rand.Seed(time.Now().UnixNano() + int64(rand.Intn(100)))
	for i := 0; i < length; i++ {
		idx := rand.Intn(visibleCharsLen)
		result = append(result, visibleChars[idx])
	}
	return string(result)
}

func SubStringWithLength(str string, start, length int) string {
	rs := []rune(str)
	strLen := len(rs)

	if start < 0 {
		start += strLen
	}
	if start < 0 {
		start = 0
	}
	if start >= strLen || length <= 0 {
		return ""
	}
	end := start + length
	if end > strLen {
		end = strLen
	}
	return string(rs[start:end])
}
