package match

import "sync/atomic"

//

type Capture struct {
	seq int64
}

var captureSeq atomic.Int64

func NewCapture() *Capture {
	return &Capture{
		seq: captureSeq.Add(1),
	}
}

//

type Captures struct {
	capture *Capture
	value   any
	next    *Captures
}

func MakeCaptures(c *Capture, v any) *Captures {
	if v == nil {
		return nil
	}
	return &Captures{
		capture: c,
		value:   v,
	}
}

func (cs *Captures) AddAll(o *Captures) *Captures {
	if cs == nil {
		return o
	}
	return &Captures{
		capture: cs.capture,
		value:   cs.value,
		next:    cs.next.AddAll(o),
	}
}
