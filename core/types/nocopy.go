package types

type NoCopy struct{}

// Lock is a no-op used by -copylocks checker from `go vet`.
func (*NoCopy) Lock()   {}
func (*NoCopy) Unlock() {}
