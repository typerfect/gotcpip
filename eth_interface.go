package eth

import "io"

type Reader interface {
	io.Closer
}

type Writer interface {
	io.Closer
	Write([]byte) (int, error)
}
