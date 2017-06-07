package memeater

import (
	"bytes"
	"strings"
)

var (
	buf  = bytes.NewBufferString("")
	data = strings.Repeat("0", 8*1024*1024)
)

func Eat() int {
	buf.WriteString(data)
	return buf.Len()
}
