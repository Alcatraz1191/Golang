package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (f *rot13Reader) Read(b []byte) (int, error){
	n, err := f.r.Read(b)
	for i,_ := range b{
		switch {
			case b[i] > 109:
				b[i] = b[i] - uint8(13)
			case b[i] > 97:
				b[i] = b[i] + uint8(13)
			case b[i] > 77:
				b[i] = b[i] + uint8(13)
			case b[i] > 65:
				b[i] = b[i] + uint8(13)
			default:
				break		
		}
	}
	return n, err
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}