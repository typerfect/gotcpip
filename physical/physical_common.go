package physical

import "io"

type physicalIO interface {
	Write([]byte)(int, error)
	Read()([]byte, error)
	io.Closer
	getInput()(chan []byte, error)
}


const (
	recvQueueSize = 5000
	frameMaxSize = 1526
)
