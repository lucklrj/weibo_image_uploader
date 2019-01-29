package http

import (
	"encoding/json"
	"fmt"
	"github.com/ddliu/go-httpclient"
	"lib/system"
	"net/http"
	"os"
	"regexp"
)

var (
	Request HttpRequest
)

type HttpRequest struct {
}

func (h *HttpRequest) Get(url string) (body string, errs error) {
	res, _ := GetHttpClient(url).Begin().Get(url)
	return res.ToString()
}

func (h *HttpRequest) Post(url string, postData map[string]string, save_cookie bool, cookies []*http.Cookie) (body string, errs error) {
	res, _ := GetHttpClient(url).Begin().WithCookie(cookies...).Post(url, postData)
	if save_cookie == true {
		saveCookie(url)
	}
	return res.ToString()
}
func GetHttpClient(url string) *httpclient.HttpClient {
	reg, _ := regexp.Compile(`^(https?:\/\/.*?)\/.*`)
	result := reg.FindAllStringSubmatch(url, -1)
	referer := result[0][1]
	
	return httpclient.NewHttpClient().Defaults(httpclient.Map{
		httpclient.OPT_REFERER:   referer,
		httpclient.OPT_USERAGENT: "Mozilla/5.0 (Macintosh; Intel Mac OS X 10.14; rv:64.0) Gecko/20100101 Firefox/64.0",
	})
}
func saveCookie(url string) {
	httpCookie := make(map[string]string)
	for _, cookie := range httpclient.Cookies(url) {
		httpCookie[cookie.Name] = cookie.Value
	}
	httpCookieBytes, _ := json.Marshal(httpCookie)
	file, err := os.OpenFile(system.GetCookName(), os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		fmt.Println(err.Error())
	}
	file.Write(httpCookieBytes)
	defer file.Close()
}
