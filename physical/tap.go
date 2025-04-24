package physical

import (
	"errors"

	"github.com/hsheth2/water"
	"github.com/typerfect/gotcpip/log"
)

type tapIO struct {
	ifce      *water.Interface
	recvQueue chan []byte
}

const (
	tapName = "tap0"
)

var (
	globalTapIO *tapIO
)

func NewTapIO() *tapIO {
	ifce, err := water.NewTAP(tapName)
	if err != nil {
		log.Log.Fatalf("new tap failed, %v", err)
	}

	err = ifce.SetPersistent(true)
	if err != nil {
		log.Log.Fatalf("tap set persistent failed,%v", err)
	}

	tap := &tapIO{
		ifce:      ifce,
		recvQueue: make(chan []byte, recvQueueSize),
	}

	// go tap read
	go func() {
		for {
			data, err := tap.readIO()
			if err != nil {
				log.Log.Errorf("tap readio failed, %v", err)
				continue
			}

			select {
			case tap.recvQueue <- data:
			default:
				log.Log.Warn("tap io recvQueue overflow")
			}
		}
	}()

	return tap
}

func (tap *tapIO) readIO() ([]byte, error) {
	buf := make([]byte, frameMaxSize)

	n, err := tap.ifce.Read(buf)
	if err != nil {
		return nil, err
	}
	return buf[:n], nil
}

func (tap *tapIO) Write(data []byte) (int, error) {
	n, err := tap.ifce.Write(data)
	if err != nil {
		return 0, err
	}

	if n != len(data) {
		return n, errors.New("tap did not write all data")
	}

	return n, nil
}

func (tap *tapIO) Read() ([]byte, error) {
	return <-tap.recvQueue, nil
}

func (tap *tapIO) Close() error {
	return tap.ifce.Close()
}

func (tap *tapIO) getInput() (chan []byte, error) {
	return tap.recvQueue, nil
}
