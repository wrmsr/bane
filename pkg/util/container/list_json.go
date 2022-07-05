package container

import (
	"encoding/json"
)

//

func (l listImpl[T]) MarshalJSON() ([]byte, error) {
	return json.Marshal(l.s)
}

func (l *listImpl[T]) UnmarshalJSON(b []byte) error {
	return json.Unmarshal(b, &l.s)
}

//

func (l mutListImpl[T]) MarshalJSON() ([]byte, error) {
	return json.Marshal(l.s)
}

func (l *mutListImpl[T]) UnmarshalJSON(b []byte) error {
	return json.Unmarshal(b, &l.s)
}
