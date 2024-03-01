package command

type Set struct {
	val LogType
}

func (s Set) Apply(val LogType) LogType {
	return s.val
}
