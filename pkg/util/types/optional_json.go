package types

import "encoding/json"

func (o Optional[T]) MarshalJSON() ([]byte, error) {
	if !o.Present() {
		return []byte("null"), nil
	}
	return json.Marshal(o.v)
}

func (o *Optional[T]) UnmarshalJSON(b []byte) error {
	*o = Optional[T]{}
	var v *T
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	if v != nil {
		o.p = true
		o.v = *v
	}
	return nil
}
