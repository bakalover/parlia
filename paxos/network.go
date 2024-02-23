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

type NetRequest struct {
	data []byte
	req  RequestType
}

type Network struct {
	// All communications represented in channels
	// Communication pattern: MPSC for separate channel
	RequestChans  []chan []byte     //Client sends only data
	AcceptorChans []chan NetRequest // Need type to distinguish
	ProposerChans []chan NetRequest // Need type to distinguish
}

var GlobalNetwork Network

func MakeNetwork(config *InitConfig) {
	requestChans := make([]chan []byte, config.Kreplicas)
	acceptorChans := make([]chan NetRequest, config.Kreplicas)
	proposerChans := make([]chan NetRequest, config.Kreplicas)

	for i := 0; i < config.Kreplicas; i++ {
		requestChans[i] = make(chan []byte)
		acceptorChans[i] = make(chan NetRequest)
		proposerChans[i] = make(chan NetRequest)
	}

	GlobalNetwork = Network{requestChans, acceptorChans, proposerChans}
}

func GetNetwork() *Network {
	return &GlobalNetwork
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

func Serialize()   {}
func Deserialize() {}
