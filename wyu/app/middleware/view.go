package middleware

import (
	"wyu/configs"
)

func TviewURL(url string, route string, security ...bool) string {
	var srcRH string = ""
	var strPrefixRH string = ""

	if len(security) != 0 && security[0] == true {
		strPrefixRH = "https://"
	} else {
		strPrefixRH = "http://"
	}

	if route == "" {
		return strPrefixRH + url
	}

	strRH, ok := configs.YuRoutes[route]
	if ok {
		srcRH = strPrefixRH + url + strRH
	} else {
		srcRH = strPrefixRH + url + route
	}

	return srcRH
}
