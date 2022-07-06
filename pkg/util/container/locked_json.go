package container

import "encoding/json"

//

func (m *lockedMapImpl[K, V]) MarshalJSON() ([]byte, error) {
	return json.Marshal(m.m)
}

func (m *lockedMapImpl[K, V]) UnmarshalJSON(b []byte) error {
	return json.Unmarshal(b, &m.m)
}

//

func (m *lockedMutMapImpl[K, V]) MarshalJSON() ([]byte, error) {
	return json.Marshal(m.m)
}

func (m *lockedMutMapImpl[K, V]) UnmarshalJSON(b []byte) error {
	return json.Unmarshal(b, &m.m)
}
