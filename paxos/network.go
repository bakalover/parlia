package paxos

// type Prequest struct {
// 	Kslot  uint64
// 	Ballot BallotNumber
// }

type RequestType uint8

const (
	Prepare  = 0
	Accept   = 1
	Promise  = 2
	Commited = 3
	Init     = 4
)

type BroadcastType uint8

const (
	Acceptors = 0
	Proposers = 1
)

const NetBufSize = 1024

// Internal cluster network message unit
type NetRequest struct {
	data []byte
	req  RequestType
}

// Just bunch of wires | Global communication point
// All communications represented in channels
// Communication pattern: MPSC for each channel
type Network struct {
	// Client sends only data
	// Outer part of Network
	RequestChans []chan []byte

	// Inner part of Network
	AcceptorChans []chan NetRequest
	ProposerChans []chan NetRequest
}

var WWW Network

func MakeNetwork(config *InitConfig) {
	requestChans := make([]chan []byte, config.Kreplicas)
	acceptorChans := make([]chan NetRequest, config.Kreplicas)
	proposerChans := make([]chan NetRequest, config.Kreplicas)

	for i := 0; i < config.Kreplicas; i++ {
		requestChans[i] = make(chan []byte, NetBufSize)
		acceptorChans[i] = make(chan NetRequest, NetBufSize)
		proposerChans[i] = make(chan NetRequest, NetBufSize)
	}

	WWW = Network{requestChans, acceptorChans, proposerChans}
}

func GetWWW() *Network {
	return &WWW
}

func (net *Network) Broadcast(data []byte, b BroadcastType, req RequestType) {
	if b == Acceptors {
		for _, ch := range net.AcceptorChans {
			ch <- NetRequest{data, req}
		}
	} else {
		for _, ch := range net.ProposerChans {
			ch <- NetRequest{data, req}
		}
	}
}
