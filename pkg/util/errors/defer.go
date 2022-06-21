package errors

type DeferStack struct {
	fns []func() error
}

func NewDeferStack() *DeferStack {
	return &DeferStack{}
}

func (ds *DeferStack) DeferErr(fn func() error) {
	ds.fns = append(ds.fns, fn)
}

func (ds *DeferStack) Defer(fn func()) {
	ds.DeferErr(func() error {
		fn()
		return nil
	})
}

func (ds *DeferStack) Call() (err error) {
	fns := ds.fns
	for i := len(fns) - 1; i >= 0; i-- {
		err = Append(err, fns[i]())
	}
	return
}
