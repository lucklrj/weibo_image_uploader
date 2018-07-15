package main

import (
	"lib/weibo"
	"fmt"
	"net/url"
	"encoding/base64"
	"lib/system"
)

func main() {
	cookies, err := weibo.ParserCookie()
	system.OutputAllErros(err, true)
	
	if cookies == nil {
		username := "sunny_lrj@yeah.net"
		username = url.QueryEscape(username)
		username = base64.StdEncoding.EncodeToString([]byte(username))
		
		password := "123123"
		weibo.Login(username, password)
	}
	
	imgUrl := weibo.UploadImg("girls.jpg", cookies)
	fmt.Println(imgUrl)
}
