package dt

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/go-sdk/utilx/json"
)

var is = IntStr(100)

func TestIntStr_MarshalJSON(t *testing.T) {
	assert.Equal(t, `"100"`, string(json.MustMarshal(is)))
}

func TestIntStr_UnmarshalJSON(t *testing.T) {
	assert.Nil(t, json.UnmarshalFromString(`"1"`, &is))
	assert.Equal(t, IntStr(1), is)

	assert.NotNil(t, json.UnmarshalFromString(`1`, &is))
}
