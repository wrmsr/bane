package container

import (
	"encoding/json"
)

//

func (m setImpl[T]) MarshalJSON() ([]byte, error) {
	return json.Marshal(m.m)
}

func (m *setImpl[T]) UnmarshalJSON(b []byte) error {
	return json.Unmarshal(b, &m.m)
}

//

func (m mutSetImpl[T]) MarshalJSON() ([]byte, error) {
	return json.Marshal(m.m)
}

func (m *mutSetImpl[T]) UnmarshalJSON(b []byte) error {
	return json.Unmarshal(b, &m.m)
}
