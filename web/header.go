package web

import (
	"net/http"
)

func setLoginHeader(request *http.Request) {
	loginHeaders := map[string]string{
		"Host":                      "ids.xidian.edu.cn",
		"Connection":                "keep-alive",
		"Pragma":                    "no-cache",
		"Cache-Control":             "no-cache",
		"Upgrade-Insecure-Requests": "1",
		"Origin":                    "http://ids.xidian.edu.cn",
		"Content-Type":              "application/x-www-form-urlencoded",
		"User-Agent":                "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/88.0.4324.96 Safari/537.36 Edg/88.0.705.56",
		"Accept":                    "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9",
		"Referer":                   "http://ids.xidian.edu.cn/authserver/login",
		"Accept-Language":           "zh-CN,zh;q=0.9,en;q=0.8,en-GB;q=0.7,en-US;q=0.6",
		"Accept-Encoding":           "gzip, deflate",
	}

	for eachKey, eachValue := range loginHeaders {
		request.Header.Add(eachKey, eachValue)
	}

	return
}
