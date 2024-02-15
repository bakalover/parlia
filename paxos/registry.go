package paxos

type Registry struct {
	// All communications represented in channels
	// Communication pattern: MPSC for separate channel

	PChans []<-chan AResponse
	AChans []<-chan PRequest
	// + Learners
}

func MakeRegistry(config *InitConfig) *Registry {
	pChans := make([]<-chan AResponse, config.KProposers)
	aChans := make([]<-chan PRequest, config.KAcceptors)
	for i := 0; i < config.KProposers; i++ {
		pChans[i] = make(<-chan AResponse)
	}
	for i := 0; i < config.KAcceptors; i++ {
		aChans[i] = make(<-chan PRequest)
	}
	return &Registry{PChans: pChans, AChans: aChans}
}
