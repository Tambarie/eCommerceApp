package helper

import "encoding/base64"

func ConvertToBase64(string2 string) string {
	data := []byte(string2)
	base64 := base64.StdEncoding.EncodeToString(data)
	return base64
}
