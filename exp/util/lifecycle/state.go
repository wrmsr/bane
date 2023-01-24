package lifecycles

import "strings"

//

type State int8

const (
	New State = iota

	Constructing
	FailedConstructing
	Constructed

	Starting
	FailedStarting
	Started

	Stopping
	FailedStopping
	Stopped

	Destroying
	FailedDestroying
	Destroyed
)

func (s State) String() string {
	switch s {

	case New:
		return "NEW"

	case Constructing:
		return "CONSTRUCTING"
	case FailedConstructing:
		return "FAILED_CONSTRUCTING"
	case Constructed:
		return "CONSTRUCTED"

	case Starting:
		return "STARTING"
	case FailedStarting:
		return "FAILED_STARTING"
	case Started:
		return "STARTED"

	case Stopping:
		return "STOPPING"
	case FailedStopping:
		return "FAILED_STOPPING"
	case Stopped:
		return "STOPPED"

	case Destroying:
		return "DESTROYING"
	case FailedDestroying:
		return "FAILED_DESTROYING"
	case Destroyed:
		return "DESTROYED"

	}
	panic(s)
}

func ParseState(s string) State {
	switch s {

	case "NEW":
		return New

	case "CONSTRUCTING":
		return Constructing
	case "FAILED_CONSTRUCTING":
		return FailedConstructing
	case "CONSTRUCTED":
		return Constructed

	case "STARTING":
		return Starting
	case "FAILED_STARTING":
		return FailedStarting
	case "STARTED":
		return Started

	case "STOPPING":
		return Stopping
	case "FAILED_STOPPING":
		return FailedStopping
	case "STOPPED":
		return Stopped

	case "DESTROYING":
		return Destroying
	case "FAILED_DESTROYING":
		return FailedDestroying
	case "DESTROYED":
		return Destroyed

	}
	panic(s)
}

func (s State) IsStable() bool {
	switch s {
	case
		New,
		Started,
		Stopped,
		Destroyed:
		return true
	}
	return false
}

func (s State) IsFailed() bool {
	switch s {
	case
		FailedConstructing,
		FailedStarting,
		FailedStopping,
		FailedDestroying:
		return true
	}
	return false
}

//

type StateMask int16

func (s State) Mask() StateMask {
	return 1 << s
}

func (m StateMask) Contains(s State) bool {
	return (m & s.Mask()) != 0
}

func (m StateMask) String() string {
	var sb strings.Builder
	n := 0
	for i := New; i <= Destroyed; i++ {
		if m.Contains(i) {
			if n > 0 {
				sb.WriteString(" | ")
			}
			sb.WriteString(i.String())
			n++
		}
	}
	return sb.String()
}
