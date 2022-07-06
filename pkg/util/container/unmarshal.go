package container

type initUnmarshal interface {
	initUnmarshal(a ...any)
}

func InitUnmarshal(o any, a ...any) {
	o.(initUnmarshal).initUnmarshal(a...)
}
