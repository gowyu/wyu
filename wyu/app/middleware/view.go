package middleware

import (
	"wyu/configs"
)

func TviewURL(url string, route string, security ...bool) (srcRH string) {
	strPrefixRH := "http://"

	if len(security) != 0 && security[0] == true {
		strPrefixRH = "https://"
	}

	if route == "" {
		srcRH = strPrefixRH + url
	} else {
		strRH, ok := configs.YuRoutes[route]

		if ok {
			srcRH = strPrefixRH + url + strRH
		} else {
			srcRH = strPrefixRH + url + route
		}
	}

	return
}
