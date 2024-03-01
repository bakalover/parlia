package command

type Mul struct {
	val LogType
}

func (a Mul) Apply(val LogType) LogType {
	return a.val * val
}
