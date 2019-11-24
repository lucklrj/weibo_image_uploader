package weibo

import (
	"lib/http"
	http2 "net/http"
	"strings"
)

func Ping(cookies []*http2.Cookie) bool {

	pingUrl := "https://weibo.com/"
	pingResult, _ := http.Request.Get(pingUrl, cookies)
	return strings.Contains(pingResult, "['islogin']='1';")
	return false
}
