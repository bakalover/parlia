package command

type LogType uint64

type Log struct {
	Sprefix []LogType    // Prefix snapshot
	actions []Applicable // Mutations after snapshot point
}

func (l *Log) Reduce() {
	// Apply some commands
}
