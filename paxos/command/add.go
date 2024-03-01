package command

type Add struct {
	val LogType
}

func (a Add) Apply(val LogType) LogType {
	return a.val + val
}
