package http

import (
	"encoding/json"
	"fmt"
	"github.com/ddliu/go-httpclient"
	"lib/system"
	"net/http"
	"os"
)

var (
	Request HttpRequest
)

type HttpRequest struct {
}

func (h *HttpRequest) Get(url string) (body string, errs error) {
	res, _ := httpclient.Begin().Get(url)
	return res.ToString()
}

func (h *HttpRequest) Post(url string, postData map[string]string, save_cookie bool, cookies []*http.Cookie) (body string, errs error) {
	res, _ := httpclient.Begin().WithCookie(cookies...).Post(url, postData)
	if save_cookie == true {
		saveCookie(url)
	}
	return res.ToString()
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
