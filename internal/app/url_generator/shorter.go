package url_generator

import (
	"encoding/base64"
)

func EncodeURL(url string) string {
	str := base64.StdEncoding.EncodeToString([]byte(url))
	return str
}

func DecodeURL(encodeUrl string) (string, error) {
	data, err := base64.StdEncoding.DecodeString(encodeUrl)
	if err != nil {
		return "", err
	}
	return string(data), nil
}
