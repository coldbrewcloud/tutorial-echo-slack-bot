package utils

import "encoding/json"

func ToJSON(v interface{}) string {
	data, err := json.Marshal(v)
	if err != nil {
		return "(error) " + err.Error()
	}
	return string(data)
}
