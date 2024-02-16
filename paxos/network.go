package paxos

type Network struct {
	// All communications represented in channels
	// Communication pattern: MPSC for separate channel

	PChans []chan AResponse
	AChans []chan PRequest
	// + Learners
}

var GlobalNetwork Network

func MakeNetwork(config *InitConfig) {
	pChans := make([]chan AResponse, config.KProposers)
	aChans := make([]chan PRequest, config.KAcceptors)
	for i := 0; i < config.KProposers; i++ {
		pChans[i] = make(chan AResponse)
	}
	for i := 0; i < config.KAcceptors; i++ {
		aChans[i] = make(chan PRequest)
	}
	GlobalNetwork = Network{PChans: pChans, AChans: aChans}
}

func GetNetwork() *Network {
	return &GlobalNetwork
}
