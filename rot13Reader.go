package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (rot *rot13Reader) Read(b []byte) (n int, err error) {
	n, err = rot.r.Read(b)
	for i, v := range b {
		if (v >=65 && v <= 77) || (v >= 97 && v <= 109) {
			b[i] += 13
		} else if (v >= 78 && v <= 90) || (v >= 110 && v <= 122) {
			b[i] -= 13
		}
	}
	return
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}

