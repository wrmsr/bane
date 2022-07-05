package container

import (
	"encoding/json"
)

//

func (m mapImpl[K, V]) MarshalJSON() ([]byte, error) {
	return json.Marshal(m.m)
}

func (m *mapImpl[K, V]) UnmarshalJSON(b []byte) error {
	return json.Unmarshal(b, &m.m)
}

//

func (m mutMapImpl[K, V]) MarshalJSON() ([]byte, error) {
	return json.Marshal(m.m)
}

func (m *mutMapImpl[K, V]) UnmarshalJSON(b []byte) error {
	return json.Unmarshal(b, &m.m)
}
