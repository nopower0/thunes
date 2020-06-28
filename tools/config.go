package tools

import (
	"encoding/json"
	"os"
)

func LoadConfigFromFile(path string, i interface{}) error {
	if f, err := os.Open(path); err != nil {
		return err
	} else if err := json.NewDecoder(f).Decode(i); err != nil {
		return err
	}
	return nil
}
