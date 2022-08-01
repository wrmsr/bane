package container

import (
	"encoding/json"

	its "github.com/wrmsr/bane/pkg/util/iterators"
	ju "github.com/wrmsr/bane/pkg/util/json"
)

//

func (m SliceMap[K, V]) MarshalJSON() ([]byte, error) {
	return MapMarshalJson[K, V](m)
}

func (m *SliceMap[K, V]) UnmarshalJSON(b []byte) error {
	ret, err := ju.UnmarshalObjectFields[K, V](b)
	if err != nil {
		return err
	}

	*m = NewSliceMap[K, V](its.OfSlice(ret))
	return nil
}

//

func (m MutSliceMap[K, V]) MarshalJSON() ([]byte, error) {
	return json.Marshal(m.m)
}

func (m *MutSliceMap[K, V]) UnmarshalJSON(b []byte) error {
	return json.Unmarshal(b, &m.m)
}
