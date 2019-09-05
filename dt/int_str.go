package dt

import (
	"errors"
	"strconv"
)

type IntStr int64

var IntStrLengthError = errors.New("length must be more than 2")

func (is *IntStr) UnmarshalJSON(b []byte) error {
	if len(b) < 2 {
		return IntStrLengthError
	}
	i, err := strconv.ParseInt(string(b[1:len(b)-1]), 10, 64)
	if err == nil {
		*is = IntStr(i)
		return nil
	}
	return err
}

func (is IntStr) MarshalJSON() ([]byte, error) {
	return []byte("\"" + strconv.FormatInt(int64(is), 10) + "\""), nil
}
