package main


type Acceptor interface {
	Promise(b Ballot)
	Accepted(n BallotNumber)
}
