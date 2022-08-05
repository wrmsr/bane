package behave

type State[E any] interface {
	Enter(entity E)
	Update(entity E)
	Exit(entity E)
	OnMessage(entity E, telegram Telegram) bool
}
