package main

import (
	"encoding/base64"
	"net/url"
	"strings"

	"github.com/fatih/color"
	_ "gopkg.in/gographics/imagick.v2/imagick"

	"github.com/lucklrj/weibo_image_uploader/lib/system"
	weibo "github.com/lucklrj/weibo_image_uploader/lib/weibo"
)

func main() {
	account := ""
	password := ""
	imageUrl := ""
	nickname := ""

	cookies, err := weibo.ParserCookie(account)
	system.OutputAllErros(err, true)

	var newLogin bool
	if cookies == nil {
		newLogin = true
	} else {
		//查看cookie是否已失效
		if weibo.Ping(cookies) == false {
			color.Red("cookie已失效,开始重新登录")
			weibo.DeleteCookie(account)
			newLogin = true
		} else {
			color.Green("cookie还在有效时间范围内")
			newLogin = false
		}
	}
	if newLogin == true {
		username := url.QueryEscape(account)
		username = base64.StdEncoding.EncodeToString([]byte(username))

		weibo.Login(username, password)
		cookies, err = weibo.ParserCookie(account)
		system.OutputAllErros(err, true)
	}
	//多张图
	remoteUrl := ""
	imgs := make([]string, 0)
	if strings.Contains(imageUrl, ",") {
		imgUrls := strings.Split(imageUrl, ",")
		for _, url := range imgUrls {
			remoteUrl = weibo.UploadImg(url, cookies, nickname)
			if remoteUrl != "" {
				imgs = append(imgs, remoteUrl)
			}

		}
	} else if imageUrl != "" {
		remoteUrl = weibo.UploadImg(imageUrl, cookies, nickname)
		if remoteUrl != "" {
			imgs = append(imgs, remoteUrl)
		}
	}
}
