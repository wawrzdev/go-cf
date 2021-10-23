package encoding

type Codec struct{}

func (c Codec) Encode(v interface{}) ([]byte, error) {
	return marshal(v)
}

func (c Codec) Decode(b []byte, v interface{}) error {
	return unmarshal(b, v)
}

//Should marshal the interface into dotenv byte slice
func marshal(v interface{}) ([]byte, error) {
	return make([]byte, 0), nil
}

//Should unmarshal the dotenv byte slice into the interface
func unmarshal(b []byte, v interface{}) error {
	return nil
}
