package main

type Proposer interface {
	Prepare(n BallotNumber) // + Slot N
	Accept(b Ballot)
}
