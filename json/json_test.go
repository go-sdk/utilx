package json

import (
	"testing"
)

type Data struct {
	A bool
	B int64
	C string
}

var data = &Data{
	A: true,
	B: 9999,
	C: "..",
}

func TestMarshalToString(t *testing.T) {
	t.Log(MarshalToString(data))
}

func TestMustMarshal(t *testing.T) {
	t.Log(string(MustMarshal(data)))
}

func TestHumanMarshal(t *testing.T) {
	t.Log(HumanMarshal(data))
}
