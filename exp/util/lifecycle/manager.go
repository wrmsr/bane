package lifecycles

import (
	"sync"

	"github.com/wrmsr/bane/pkg/util/maps"
)

type managerEntry struct {
	Controller

	dependencies maps.Set[*managerEntry]
	dependents   maps.Set[*managerEntry]
}

type Manager struct {
	mtx sync.Mutex
	m   map[*Lifecycle]*managerEntry
}

func NewManager() *Manager {
	return &Manager{
		m: make(map[*Lifecycle]*managerEntry),
	}
}

/*
class LifecycleManager(AbstractLifecycle):

    class Entry(dc.Pure):
        controller: LifecycleController
        dependencies: ta.Set['LifecycleManager.Entry'] = dc.field(default_factory=ocol.IdentitySet)
        dependents: ta.Set['LifecycleManager.Entry'] = dc.field(default_factory=ocol.IdentitySet)

    def __init__(
            self,
            *,
            lock: lang.DefaultLockable = None,
    ) -> None:
        super().__init__()

        self._lock = lang.default_lock(lock, True)

        self._entries_by_lifecycle: ta.MutableMapping[Lifecycle, LifecycleManager.Entry] = ocol.IdentityKeyDict()

    @property
    def controller(self: LifecycleManagerT) -> 'LifecycleController[LifecycleManagerT]':
        return self.lifecycle_controller

    @property
    def state(self) -> LifecycleState:
        return self.lifecycle_controller.state

    @staticmethod
    def _get_controller(lifecycle: Lifecycle) -> LifecycleController:
        if isinstance(lifecycle, LifecycleController):
            return lifecycle
        elif isinstance(lifecycle, AbstractLifecycle):
            return lifecycle.lifecycle_controller
        elif isinstance(lifecycle, Lifecycle):
            return LifecycleController(lifecycle)
        else:
            raise TypeError(lifecycle)

    def _add_internal(self, lifecycle: Lifecycle, dependencies: ta.Iterable[Lifecycle]) -> Entry:
        check.state(self.state.phase < LifecycleStates.STOPPING.phase and not self.state.is_failed)

        check.isinstance(lifecycle, Lifecycle)
        try:
            entry = self._entries_by_lifecycle[lifecycle]
        except KeyError:
            controller = self._get_controller(lifecycle)
            entry = self._entries_by_lifecycle[lifecycle] = LifecycleManager.Entry(controller)

        for dep in dependencies:
            check.isinstance(dep, Lifecycle)
            dep_entry = self._add_internal(dep, [])
            entry.dependencies.add(dep_entry)
            dep_entry.dependents.add(entry)

        return entry

    def add(
            self,
            lifecycle: Lifecycle,
            dependencies: ta.Iterable[Lifecycle] = (),
    ) -> Entry:
        check.state(self.state.phase < LifecycleStates.STOPPING.phase and not self.state.is_failed)

        with self._lock():
            entry = self._add_internal(lifecycle, dependencies)

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
