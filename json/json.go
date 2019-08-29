package json

import (
	"github.com/json-iterator/go"
)

var (
	json = jsoniter.ConfigCompatibleWithStandardLibrary

	Marshal         = json.Marshal
	MarshalToString = json.MarshalToString
	MarshalIndent   = json.MarshalIndent

	Unmarshal           = json.Unmarshal
	UnmarshalFromString = json.UnmarshalFromString

	NewDecoder = json.NewDecoder
	NewEncoder = json.NewEncoder
)

func MustMarshal(v interface{}) []byte {
	bs, _ := Marshal(v)
	return bs
}

func HumanMarshal(v interface{}) string {
	bs, _ := MarshalIndent(v, "", "  ")
	return string(bs)
}
