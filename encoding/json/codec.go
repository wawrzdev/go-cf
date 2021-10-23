package encoding

import "encoding/json"

type Codec struct{}

func (c Codec) Encode(v interface{}) ([]byte, error) {
	return json.MarshalIndent(v, "", "  ")
}

func (c Codec) Decode(b []byte, v interface{}) error {
	return json.Unmarshal(b, v)
}
