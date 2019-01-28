package main

import (
	"encoding/base64"
	"fmt"
	"lib/system"
	"lib/weibo"
	"net/url"
)

func main() {
	cookies, err := weibo.ParserCookie(*system.Account)
	system.OutputAllErros(err, true)
	
	if cookies == nil {
		username := url.QueryEscape(*system.Account)
		username = base64.StdEncoding.EncodeToString([]byte(username))
		
		weibo.Login(username, *system.Password)
	}
	
	//imgUrl := weibo.UploadImg("girls.jpg", cookies)
	imgUrl := weibo.UploadImg(*system.ImageUrl, cookies, *system.Nickname)
	fmt.Println(imgUrl)
}
