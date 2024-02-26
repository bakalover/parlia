package replica

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/bakalover/parlia/paxos"
	"github.com/bakalover/tate"
)

type Replica struct {
	log     []string
	server  *http.Server
	PortGen *paxos.Generator
	port    string
}

// --------------------------Replica----------------------------

func (r *Replica) Kill() {
	ctx, cancel := context.WithCancel(context.Background())
	r.PortGen.InvalidatePort(r.port)
	defer cancel()
	r.server.Shutdown(ctx)
	r.log = nil
}

func (r *Replica) Step(stepTime time.Duration) {
	r.port = fmt.Sprintf(":%s", r.PortGen.GeneratePort())
	r.server = &http.Server{Addr: r.port, Handler: nil}
	serverRoutine := tate.NewNursery(nil)

	serverRoutine.Add(func(c *tate.Linker) {
		r.server.ListenAndServe()
	})

	time.AfterFunc(stepTime, func() {
		r.Kill()
	})

	serverRoutine.Join()
}

func (r *Replica) Apply(command string) {
	//Prepare
}

// --------------------------Replica----------------------------

// --------------------------Acceptor---------------------------

func (r *Replica) Promise(b paxos.Ballot) {

}

func (r *Replica) Accepted(n paxos.BallotNumber) {

}

// --------------------------Acceptor---------------------------

// --------------------------Proposer---------------------------

func (r *Replica) Prepare(b paxos.Ballot) {

}

func (r *Replica) Accept(n paxos.BallotNumber) {

}

// --------------------------Proposer---------------------------
