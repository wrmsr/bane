package behave

type State[E, T any] interface {
	Enter(entity E)
	Update(entity E)
	Exit(entity E)
	OnMessage(entity E, telegram Telegram[T]) bool
}
