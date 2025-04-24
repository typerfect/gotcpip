package physical

import "errors"

type loopbackIO struct {
	recvQueue chan []byte
}

var (
	globalLoopbackIO *loopbackIO
)

func NewLoopbackIO() *loopbackIO {
	return &loopbackIO{
		recvQueue: make(chan []byte, recvQueueSize),
	}
}

func (loop *loopbackIO) Write(data []byte) (int, error) {
	loop.recvQueue <- data
	return len(data), nil
}

func (loop *loopbackIO) Read() ([]byte, error) {
	data := <-loop.recvQueue
	if len(data) == 0 {
		return nil, errors.New("loop recvQueue channel has been close")
	}

	return data, nil
}

func (loop *loopbackIO) Close() error {
	close(loop.recvQueue)
	return nil
}

func (loop *loopbackIO) getInput() (chan []byte, error) {
	return loop.recvQueue, nil
}
