package replica

import (
	"context"
	"net/http"
	"time"

	"github.com/bakalover/parlia/paxos"
	"github.com/bakalover/tate"
)

type Replica struct {
	log    []string
	server *http.Server
}

// --------------------------Replica----------------------------

func (r *Replica) Kill() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	r.server.Shutdown(ctx)
	r.log = nil
}

func (r *Replica) Step(stepTime time.Duration) {

	r.server = &http.Server{Addr: "todoport", Handler: nil}
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
