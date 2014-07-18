package protocols

// protocols is an enum of possible protocols to use.

type Protocol int

const (
	PStdOut       Protocol = 0
	PProtoBuffers Protocol = 1
	PJson         Protocol = 2
)
