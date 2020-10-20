package utils

import (
	"encoding/json"
	"io/ioutil"
)

func ReadJsonFile(path string) ([]map[string]interface{}, error) {
	var data []map[string]interface{}
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func Jsonify(m interface{}) string {
	b, _ := json.MarshalIndent(m, "", "  ")
	return string(b)
}
