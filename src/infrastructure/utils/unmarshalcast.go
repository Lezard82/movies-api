package utils

import (
	"encoding/json"
)

func UnmarshalCast(data string) ([]string, error) {
	var cast []string
	err := json.Unmarshal([]byte(data), &cast)
	if err != nil {
		return []string{}, err
	}
	return cast, nil
}
