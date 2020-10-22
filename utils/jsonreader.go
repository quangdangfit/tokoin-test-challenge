package utils

import (
	"encoding/json"
	"io/ioutil"
)

func ReadJsonFile(path string) ([]byte, error) {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

func Jsonify(m interface{}) string {
	b, _ := json.MarshalIndent(m, "", "  ")
	return string(b)
}
