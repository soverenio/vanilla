package bstream

import (
	"io"
)

var Zero io.Reader = ZeroReader{}

type ZeroReader struct{}

func (z ZeroReader) Read(p []byte) (n int, err error) {
	for i := 0; i < len(p); i++ {
		p[i] = 0
	}
	return len(p), nil
}
