package container

import (
	"encoding/json"
	"reflect"

	its "github.com/wrmsr/bane/pkg/util/iterators"
	ju "github.com/wrmsr/bane/pkg/util/json"
	bt "github.com/wrmsr/bane/pkg/util/types"
)

//

func (m SliceMap[K, V]) MarshalJSON() ([]byte, error) {
	ro := make(ju.RawObject, len(m.s))
	for i, kv := range m.s {
		ks, err := ju.MarshalAsText(reflect.ValueOf(kv.K))
		if err != nil {
			return nil, err
		}

		vb, err := json.Marshal(kv.V)
		if err != nil {
			return nil, err
		}

		ro[i] = bt.KvOf[string, json.RawMessage](ks, vb)
	}
	return json.Marshal(ro)
}

func (m *SliceMap[K, V]) UnmarshalJSON(b []byte) error {
	var ro ju.RawObject
	if err := json.Unmarshal(b, &ro); err != nil {
		return err
	}

	ret, err := ju.UnmarshalRawObject[K, V](ro)
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
