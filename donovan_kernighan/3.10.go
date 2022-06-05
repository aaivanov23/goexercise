package main

import (
	"bytes"
	"fmt"
)

func main() {

	fmt.Println(comma("12345678918947183289"))
}

func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}

	var buf bytes.Buffer
	t := len(s) % 3

	if t != 0 {
		buf.WriteString(s[:t])
		buf.WriteByte(',')
	}

	for i := t; i < len(s); i += 3 {
		buf.WriteString(s[i : i+3])
		buf.WriteByte(',')
	}

	buf.Truncate(buf.Len() - 1)
	return buf.String()
}
