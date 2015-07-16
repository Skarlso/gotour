package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (rot rot13Reader) Read(b []byte) (n int, err error) {
	n, err = rot.r.Read(b)
	for i, value := range b[:n] {
        if value <= 'Z' && value >='A'  {
            b[i] = (value - 'A' + 13)%26 + 'A'
        } else if value >= 'a' && value <= 'z' {
            b[i] = (value - 'a' + 13)%26 + 'a'
        } 
	}
	return n, err
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}

