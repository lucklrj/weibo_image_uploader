package http

import (
	"github.com/parnurzeal/gorequest"
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
