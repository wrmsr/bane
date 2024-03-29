package behave

import (
	bt "github.com/wrmsr/bane/core/types"
)

//

type State[E, T any] interface {
	Equals(o State[E, T]) bool
	Enter(entity bt.Optional[E])
	Update(entity bt.Optional[E])
	Exit(entity bt.Optional[E])
	OnMessage(entity bt.Optional[E], telegram Telegram[T]) bool
}

//

type StateMachine[E, T any] struct {
	generatedIdentity

	owner bt.Optional[E]

	current, previous, global bt.Optional[State[E, T]]
}

func NewStateMachine[E, T any](owner bt.Optional[E], initial, global bt.Optional[State[E, T]]) *StateMachine[E, T] {
	return &StateMachine[E, T]{
		owner: owner,

		current: initial,
		global:  global,
	}
}

func (sm *StateMachine[E, T]) Owner() bt.Optional[E] { return sm.owner }

func (sm *StateMachine[E, T]) Current() bt.Optional[State[E, T]]  { return sm.current }
func (sm *StateMachine[E, T]) Previous() bt.Optional[State[E, T]] { return sm.previous }
func (sm *StateMachine[E, T]) Global() bt.Optional[State[E, T]]   { return sm.global }

var _ Telegraph[int] = &StateMachine[string, int]{}

func (sm *StateMachine[E, T]) Update() {
	if sm.global.Present() {
		sm.global.Value().Update(sm.owner)
	}
	if sm.current.Present() {
		sm.current.Value().Update(sm.owner)
	}
}

func (sm *StateMachine[E, T]) Change(ns State[E, T]) {
	sm.previous = sm.current
	sm.change(ns)
}

func (sm *StateMachine[E, T]) change(ns State[E, T]) {
	if sm.current.Present() {
		sm.current.Value().Exit(sm.owner)
	}
	sm.current = bt.Just(ns)
	if sm.current.Present() {
		sm.current.Value().Enter(sm.owner)
	}
}

func (sm *StateMachine[E, T]) Revert() bool {
	if !sm.previous.Present() {
		return false
	}
	sm.Change(sm.previous.Value())
	return true
}

func (sm *StateMachine[E, T]) IsIn(state State[E, T]) bool {
	return sm.current.Present() && sm.current.Value().Equals(state)
}

func (sm *StateMachine[E, T]) HandleMessage(telegram Telegram[T]) bool {
	if sm.current.Present() && sm.current.Value().OnMessage(sm.owner, telegram) {
		return true
	}
	if sm.global.Present() && sm.global.Value().OnMessage(sm.owner, telegram) {
		return true
	}
	return false
}

//

/*
type StackStateMachine(StateMachine[E, StateT]):

    def __init__(
            self,
            owner: ta.Optional[E] = None,
            *,
            initial: ta.Optional[StateT] = None,
    ) -> None:
        super().__init__(owner, initial=initial)

        sm.stack: ta.List[StateT] = []

    @property
    def previous(self) -> ta.Optional[StateT]:
        return sm.stack[-1] if sm.stack else None

    def change(self, ns: StateT, *, push: bool = True) -> None:
        if push and sm.current is not None:
            sm.stack.append(sm.current)
        sm.change(ns)

    def revert(self) -> bool:
        if not sm.stack:
            return False
        self.change(sm.stack.pop(), push=False)
        return True
*/
