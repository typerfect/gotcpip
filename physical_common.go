package physical

const (
	recvQueueSize = 5000
	frameMaxSize  = 1526
)

type IOIndex int

const (
	LoopbackIOIndex IOIndex = 1
	ExternalIOIndex IOIndex = 1
)
