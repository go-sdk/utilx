package json

import (
	"testing"

	"github.com/stretchr/testify/assert"
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

var dataStr = `{"A":true,"B":9999,"C":".."}`

func TestMarshalToString(t *testing.T) {
	str, err := MarshalToString(data)
	assert.Nil(t, err)
	assert.Equal(t, dataStr, str)
}

func TestUnmarshalFromString(t *testing.T) {
	v := &Data{}
	err := UnmarshalFromString(dataStr, v)
	assert.Nil(t, err)
	assert.Equal(t, dataStr, string(MustMarshal(v)))
}

func TestMustMarshal(t *testing.T) {
	assert.Equal(t, dataStr, string(MustMarshal(data)))
}

func TestMustMarshalToString(t *testing.T) {
	assert.Equal(t, dataStr, MustMarshalToString(data))
}

func TestHumanMarshal(t *testing.T) {
	t.Log(HumanMarshal(data))
}

func TestReMarshal(t *testing.T) {
	assert.Equal(t, dataStr, string(ReMarshal([]byte(`{"A"  :true,"B" :      9999,"C":".."    }`))))
}
