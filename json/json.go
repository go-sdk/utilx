package json

import (
	"encoding/json"
)

var (
	Marshal       = json.Marshal
	MarshalIndent = json.MarshalIndent

	Unmarshal = json.Unmarshal

	NewDecoder = json.NewDecoder
	NewEncoder = json.NewEncoder
)

func MarshalToString(v interface{}) (string, error) {
	bs, err := Marshal(v)
	return string(bs), err
}

func UnmarshalFromString(str string, v interface{}) error {
	return Unmarshal([]byte(str), v)
}

func MustMarshal(v interface{}) []byte {
	bs, _ := Marshal(v)
	return bs
}

func MustMarshalToString(v interface{}) string {
	bs, _ := Marshal(v)
	return string(bs)
}

func HumanMarshal(v interface{}) []byte {
	bs, _ := MarshalIndent(v, "", "  ")
	return bs
}

func ReMarshal(bs []byte) []byte {
	var data interface{}
	_ = Unmarshal(bs, &data)
	return MustMarshal(data)
}
