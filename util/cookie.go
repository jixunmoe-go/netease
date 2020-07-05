package util

import (
	"net/http"
)

func ParseCookie(cookieStr string) map[string]interface{} {
	header := http.Header{}
	header.Add("Cookie", cookieStr)

	req := http.Request{Header: header}
	cookies := req.Cookies()

	result := map[string]interface{}{}
	size := len(cookies)
	for i := 0; i < size; i++ {
		result[cookies[i].Name] = cookies[i].Value
	}

	return result
}
