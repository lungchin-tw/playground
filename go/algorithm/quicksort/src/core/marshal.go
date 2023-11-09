package core

import "encoding/json"

func ToJSONString(m interface{}) string {
	b, _ := json.Marshal(m)
	return string(b)
}
