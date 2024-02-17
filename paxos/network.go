package paxos

type Prequest struct {
	Kslot  uint64
	Ballot BallotNumber
}

type RequestType uint8

const (
	Prepare  = 0
	Accept   = 1
	Promise  = 2
	Commited = 3
)

type NetRequest struct {
	data []byte
	t    RequestType
}

type Network struct {
	// All communications represented in channels
	// Communication pattern: MPSC for separate channel
	RequestChans  []chan NetRequest
	AcceptorChans []chan NetRequest
	ProposerChans []chan NetRequest
}

var GlobalNetwork Network

func MakeNetwork(config *InitConfig) {
	requestChans := make([]chan NetRequest, config.Kreplicas)
	acceptorChans := make([]chan NetRequest, config.Kreplicas)
	proposerChans := make([]chan NetRequest, config.Kreplicas)

	for i := 0; i < config.Kreplicas; i++ {
		requestChans[i] = make(chan NetRequest)
		acceptorChans[i] = make(chan NetRequest)
		proposerChans[i] = make(chan NetRequest)
	}

	GlobalNetwork = Network{requestChans, acceptorChans, proposerChans}
}

func GetNetwork() *Network {
	return &GlobalNetwork
}

func (net *Network) Broadcast(req Prequest, t RequestType) {

	var ser []byte
	// Serialize request
	if t == Prepare || t == Accept {
		for _, ch := range net.AcceptorChans {
			ch <- NetRequest{ser, t}
		}
	} else {
		for _, ch := range net.ProposerChans {
			ch <- NetRequest{ser, t}
		}
	}
}
