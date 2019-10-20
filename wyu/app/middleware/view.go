package middleware

import "wyu/configs"

func TviewURL(url string, index string) string {
	var srcRH string = ""
	var strPrefixRH string = "http://"

	if index == "" {
		return strPrefixRH + url
	}

	strRH, ok := configs.WYuRouteHttp[index]
	if ok {
		srcRH = strPrefixRH + url + strRH
	} else {
		srcRH = strPrefixRH + url + index
	}

	return srcRH
}
