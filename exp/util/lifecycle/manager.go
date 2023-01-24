package lifecycles

type managerEntry struct {
	controller

	dependencies map[*managerEntry]struct{}
	dependents   map[*managerEntry]struct{}
}

type Manager struct {
	controller

	m map[*Lifecycle]*managerEntry
}

func NewManager() *Manager {
	m := &Manager{
		m: make(map[*Lifecycle]*managerEntry),
	}
	// FIXME: ...
	m.lc = m
	return m
}

var dependableStates = New.Mask() |
	Constructing.Mask() |
	Constructed.Mask() |
	Starting.Mask() |
	Started.Mask()

func (m *Manager) addInternal(lc *Lifecycle, deps []*Lifecycle) (*managerEntry, error) {
	if !dependableStates.Contains(m.st) {
		return nil, StateError{Current: m.st, Expected: dependableStates}
	}

	var e *managerEntry
	if e = m.m[lc]; m == nil {
		e = &managerEntry{
			controller: controller{
				lc: lc,
			},
		}
		m.m[lc] = e
	}

	if len(deps) > 0 {
		e.dependencies = make(map[*managerEntry]struct{})
		for _, dep := range deps {
			de, err := m.addInternal(dep, nil)
			if err != nil {
				return nil, err
			}
			e.dependencies[de] = struct{}{}
			if de.dependents == nil {
				de.dependents = make(map[*managerEntry]struct{})
			}
			de.dependents[e] = struct{}{}
		}
	}

	return e, nil
}

func (m *Manager) Add(lc *Lifecycle, deps []*Lifecycle) error {
	if !dependableStates.Contains(m.st) {
		return StateError{Current: m.st, Expected: dependableStates}
	}

	m.mtx.Lock()
	defer m.mtx.Unlock()

	e, err := m.addInternal(lc, deps)
	if err != nil {
		return err
	}

	if m.st >= Constructing {
		rec := func(e *managerEntry) error {
			if e.st 
		}
	}
}

/*
           if self.state.phase >= LifecycleStates.CONSTRUCTING.phase:
               def rec(e):
                   if e.controller.state.phase < LifecycleStates.CONSTRUCTED.phase:
                       for dep in e.dependencies:
                           rec(dep)
                       e.controller.lifecycle_construct()
               rec(entry)
           if self.state.phase >= LifecycleStates.STARTING.phase:
               def rec(e):
                   if e.controller.state.phase < LifecycleStates.STARTED.phase:
                       for dep in e.dependencies:
                           rec(dep)
                       e.controller.lifecycle_start()
               rec(entry)

           return entry

   def construct(self) -> None:
       self.lifecycle_construct()

   def _do_lifecycle_construct(self) -> None:
       def rec(entry: LifecycleManager.Entry) -> None:
           for dep in entry.dependencies:
               rec(dep)
           if entry.controller.state.is_failed:
               raise LifecycleStateException(entry.controller)
           if entry.controller.state.phase < LifecycleStates.CONSTRUCTED.phase:
               entry.controller.lifecycle_construct()
       for entry in self._entries_by_lifecycle.values():
           rec(entry)

   def start(self) -> None:
       self.lifecycle_start()

   def _do_lifecycle_start(self) -> None:
       def rec(entry: LifecycleManager.Entry) -> None:
           for dep in entry.dependencies:
               rec(dep)
           if entry.controller.state.is_failed:
               raise LifecycleStateException(entry.controller)
           if entry.controller.state.phase < LifecycleStates.CONSTRUCTED.phase:
               entry.controller.lifecycle_construct()
           if entry.controller.state.phase < LifecycleStates.STARTED.phase:
               entry.controller.lifecycle_start()
       for entry in self._entries_by_lifecycle.values():
           rec(entry)

   def stop(self) -> None:
       self.lifecycle_stop()

   def _do_lifecycle_stop(self) -> None:
       def rec(entry: LifecycleManager.Entry) -> None:
           for dep in entry.dependents:
               rec(dep)
           if entry.controller.state.is_failed:
               raise LifecycleStateException(entry.controller)
           if entry.controller.state is LifecycleStates.STARTED:
               entry.controller.lifecycle_stop()
       for entry in self._entries_by_lifecycle.values():
           rec(entry)

   def destroy(self) -> None:
       self.lifecycle_destroy()

   def _do_lifecycle_destroy(self) -> None:
       def rec(entry: LifecycleManager.Entry) -> None:
           for dep in entry.dependents:
               rec(dep)
           if entry.controller.state.phase < LifecycleStates.DESTROYED.phase:
               entry.controller.lifecycle_destroy()
       for entry in self._entries_by_lifecycle.values():
           rec(entry)
*/
