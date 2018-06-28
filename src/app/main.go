package main

import (
	"net/url"
	"encoding/base64"
	"github.com/parnurzeal/gorequest"
	"lib/system"
	"fmt"
	"sync"
)

var (
	Request HttpRequest
)

type HttpRequest struct {
	Request *gorequest.SuperAgent
	mu      sync.Mutex
}

func (h *HttpRequest) Get(url string) (body string, errs []error) {
	h.mu.Lock()
	_, body, errs = h.Request.Get(url).End()
	h.mu.Unlock()
	return body, errs
}
func (h *HttpRequest) Post(url string, postData map[string]string) (body string, errs []error) {
	h.mu.Lock()
	_, body, errs = h.Request.Post(url).Type("multipart").Send(postData).End()
	h.mu.Unlock()
	return body, errs
}

func init() {
	Request = HttpRequest{Request: gorequest.New()}
}
func main() {
	username := "sunny_lrj@yeaj.net"
	username = url.QueryEscape(username)
	username = base64.StdEncoding.EncodeToString([]byte(username))
	
	url := "http://login.sina.com.cn/sso/prelogin.php?entry=weibo&callback=sinaSSOController.preloginCallBack&su=" + username + "&rsakt=mod&checkpin=1&client=ssologin.js(v1.4.18)&_=1461819359582"
	html, errs := Request.Get(url)
	system.OutputAllErros(errs, true)
	
	fmt.Println(html)
	
}
