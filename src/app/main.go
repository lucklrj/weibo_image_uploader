package main

import (
	"net/url"
	"encoding/base64"
	"lib/weibo"
	"fmt"
)

func main() {
	username := "sunny_lrj@yeah.net"
	username = url.QueryEscape(username)
	username = base64.StdEncoding.EncodeToString([]byte(username))
	
	password := "123asd123"
	weibo.Login(username, password)
	
	imgUrl := weibo.UploadImg("girls.jpg")
	fmt.Println(imgUrl)
}
