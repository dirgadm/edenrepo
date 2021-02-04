package edenrepo

import (
	"net/url"
	"strings"
)

type QueryString struct {
	QueryString string
}

func queryStringBuilder(v url.Values) (string, error) {
	var str strings.Builder
	var e error
	if len(v) == 0 {
		str.WriteString("?")
		str.WriteString("perpage=1")
	} else {
		var i = 0
		for k, j := range v {
			if i == 0 {
				str.WriteString("?")
				str.WriteString(k + "=" + j[0])
				i++
			} else {
				str.WriteString("&")
				str.WriteString(k + "=" + j[0])
			}
		}
		if v.Get("perpage") == "" {
			str.WriteString("&")
			str.WriteString("perpage=1")
		}
	}

	return str.String(), e
}
