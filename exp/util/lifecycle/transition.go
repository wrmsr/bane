package lifecycles

type transition struct {
	old StateMask

	intermediate,
	failed,
	new State
}

var (
	construct = transition{
		old: New.Mask(),

		intermediate: Constructing,
		failed:       FailedConstructing,
		new:          Constructed,
	}

	start = transition{
		old: Constructed.Mask(),

		intermediate: Starting,
		failed:       FailedStarting,
		new:          Started,
	}

	stop = transition{
		old: Started.Mask(),

		intermediate: Stopping,
		failed:       FailedStopping,
		new:          Stopped,
	}

	destroy = transition{
		old: New.Mask() |
			Constructing.Mask() |
			FailedConstructing.Mask() |
			Constructed.Mask() |
			Starting.Mask() |
			FailedStarting.Mask() |
			Started.Mask() |
			Stopping.Mask() |
			FailedStopping.Mask() |
			Stopped.Mask(),

		intermediate: Destroying,
		failed:       FailedDestroying,
		new:          Destroyed,
	}
)
