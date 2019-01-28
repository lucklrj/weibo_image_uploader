package main

import (
	"encoding/base64"
	"flag"
	"fmt"
	"github.com/fatih/color"
	"lib/system"
	"lib/weibo"
	"net/url"
	"os"
)

var (
	ImageUrl = flag.String("url", "", "图片位置")
	Account  = flag.String("account", "", "微博账号")
	Password = flag.String("password", "", "微博密码")
	Nickname = flag.String("nickname", "", "微博昵称")
)

func main() {
	cookies, err := weibo.ParserCookie(*Account)
	system.OutputAllErros(err, true)
	
	if cookies == nil {
		username := url.QueryEscape(*Account)
		username = base64.StdEncoding.EncodeToString([]byte(username))
		
		weibo.Login(username, *Password)
	}
	
	//imgUrl := weibo.UploadImg("girls.jpg", cookies)
	imgUrl := weibo.UploadImg(*ImageUrl, cookies, *Nickname)
	fmt.Println(imgUrl)
}

func init() {
	
	flag.Parse()
	if *ImageUrl == "" {
		color.Red("缺少url参数")
		os.Exit(0)
	}
	if *Account == "" {
		color.Red("缺少account参数")
		os.Exit(0)
	}
	if *Password == "" {
		color.Red("缺少password参数")
		os.Exit(0)
	}
}
