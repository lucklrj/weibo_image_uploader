package http

import (
	"encoding/json"
	"net/http"
	"os"
	"regexp"

	"github.com/ddliu/go-httpclient"
	"github.com/fatih/color"

	"github.com/lucklrj/weibo_image_uploader/lib/system"
)

var (
	Request HttpRequest
)

type HttpRequest struct {
}

func (h *HttpRequest) Get(url string, cookies []*http.Cookie) (body string, errs error) {
	if cookies == nil {
		cookies = make([]*http.Cookie, 0)
	}
	res, _ := GetHttpClient(url).Begin().WithCookie(cookies...).Get(url)
	return res.ToString()
}

func (h *HttpRequest) Post(url string, postData map[string]string, save_cookie bool, cookies []*http.Cookie) (body string, errs error) {
	hc := GetHttpClient(url)
	res, err := hc.Begin().WithCookie(cookies...).Post(url, postData)
	if err != nil {
		//color.Red(err.Error())
		return "", err
	}
	if save_cookie == true {
		saveCookie(hc, url)
	}
	return res.ToString()
}
func GetHttpClient(url string) *httpclient.HttpClient {
	reg, _ := regexp.Compile(`^(https?:\/\/.*?)\/.*`)
	result := reg.FindAllStringSubmatch(url, -1)
	referer := result[0][1]

	return httpclient.NewHttpClient().Defaults(httpclient.Map{
		httpclient.OPT_REFERER:    referer,
		httpclient.OPT_USERAGENT:  "Mozilla/5.0 (Macintosh; Intel Mac OS X 10.14; rv:64.0) Gecko/20100101 Firefox/64.0",
		httpclient.OPT_UNSAFE_TLS: true,
	})
}
func saveCookie(hc *httpclient.HttpClient, url string) {
	httpCookie := make(map[string]string)
	for _, cookie := range hc.Cookies(url) {
		httpCookie[cookie.Name] = cookie.Value
	}
	httpCookieBytes, _ := json.Marshal(httpCookie)
	file, err := os.OpenFile(system.GetCookName(), os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		color.Red(err.Error())
	}
	file.Write(httpCookieBytes)
	defer file.Close()
}
