package container

import (
	"encoding/json"
)

//

func (s setImpl[T]) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.m)
}

func (s *setImpl[T]) UnmarshalJSON(b []byte) error {
	return json.Unmarshal(b, &s.m)
}

//

func (s mutSetImpl[T]) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.m)
}

func (s *mutSetImpl[T]) UnmarshalJSON(b []byte) error {
	return json.Unmarshal(b, &s.m)
}
