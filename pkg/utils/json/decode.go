package json

import (
	"encoding/json"

	bt "github.com/wrmsr/bane/pkg/utils/types"
)

type Object []bt.Kv[string, json.RawMessage]

//type objectDecoder struct {
//	ks []string
//	vs map[string]json.RawMessage
//}

//func decodeObject(dec *json.Decoder, o *objectDecoder) error {
//	hasKey := make(map[string]struct{}, len(o.vs))
//	for {
//		token, err := dec.Token()
//		if err != nil {
//			return err
//		}
//		if delim, ok := token.(json.Delim); ok && delim == '}' {
//			return nil
//		}
//		key := token.(string)
//		if _, ok := hasKey[key]; ok {
//			// duplicate key
//			for j, k := range o.ks {
//				if k == key {
//					copy(o.ks[j:], o.ks[j+1:])
//					break
//				}
//			}
//			o.ks[len(o.ks)-1] = key
//		} else {
//			hasKey[key] = struct{}{}
//			o.ks = append(o.ks, key)
//		}
//
//		token, err = dec.Token()
//		if err != nil {
//			return err
//		}
//		if delim, ok := token.(json.Delim); ok {
//			switch delim {
//			case '{':
//				if values, ok := o.vs[key].(map[string]json.RawMessage); ok {
//					newMap := objectDecoder{
//						ks: make([]string, 0, len(values)),
//						vs: values,
//					}
//					if err = decodeObject(dec, &newMap); err != nil {
//						return err
//					}
//					o.vs[key] = newMap
//				} else if oldMap, ok := o.vs[key].(objectDecoder); ok {
//					newMap := objectDecoder{
//						ks: make([]string, 0, len(oldMap.vs)),
//						vs: oldMap.vs,
//					}
//					if err = decodeObject(dec, &newMap); err != nil {
//						return err
//					}
//					o.vs[key] = newMap
//				} else if err = decodeObject(dec, &objectDecoder{}); err != nil {
//					return err
//				}
//			case '[':
//				if values, ok := o.vs[key].([]json.RawMessage); ok {
//					if err = decodeObjectSlice(dec, values); err != nil {
//						return err
//					}
//				} else if err = decodeObjectSlice(dec, []json.RawMessage{}); err != nil {
//					return err
//				}
//			}
//		}
//	}
//}

//func decodeObjectSlice(dec *json.Decoder, s []json.RawMessage) error {
//	for i := 0; ; i++ {
//		token, err := dec.Token()
//		if err != nil {
//			return err
//		}
//		if delim, ok := token.(json.Delim); ok {
//			switch delim {
//			case '{':
//				if i < len(s) {
//					if values, ok := s[i].(map[string]json.RawMessage); ok {
//						newMap := objectDecoder{
//							ks: make([]string, 0, len(values)),
//							vs: values,
//						}
//						if err = decodeObject(dec, &newMap); err != nil {
//							return err
//						}
//						s[i] = newMap
//					} else if oldMap, ok := s[i].(objectDecoder); ok {
//						newMap := objectDecoder{
//							ks: make([]string, 0, len(oldMap.vs)),
//							vs: oldMap.vs,
//						}
//						if err = decodeObject(dec, &newMap); err != nil {
//							return err
//						}
//						s[i] = newMap
//					} else if err = decodeObject(dec, &objectDecoder{}); err != nil {
//						return err
//					}
//				} else if err = decodeObject(dec, &objectDecoder{}); err != nil {
//					return err
//				}
//			case '[':
//				if i < len(s) {
//					if values, ok := s[i].([]json.RawMessage); ok {
//						if err = decodeObjectSlice(dec, values); err != nil {
//							return err
//						}
//					} else if err = decodeObjectSlice(dec, []json.RawMessage{}); err != nil {
//						return err
//					}
//				} else if err = decodeObjectSlice(dec, []json.RawMessage{}); err != nil {
//					return err
//				}
//			case ']':
//				return nil
//			}
//		}
//	}
//}
