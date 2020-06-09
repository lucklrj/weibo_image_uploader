package weibo

import (
	http2 "net/http"
	"strings"

	"github.com/lucklrj/weibo_image_uploader/lib/http"
)

func Ping(cookies []*http2.Cookie) bool {

	pingUrl := "https://weibo.com/"
	pingResult, _ := http.Request.Get(pingUrl, cookies)
	return strings.Contains(pingResult, "retcode=6102")
}
