package utils

import (
	"encoding/json"
)

func MarshalCast(cast []string) (string, error) {
	data, err := json.Marshal(cast)
	if err != nil {
		return "[]", err
	}
	return string(data), nil
}
