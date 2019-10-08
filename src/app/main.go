package main

import (
	"encoding/base64"
	"github.com/fatih/color"
	"lib/http"
	"lib/system"
	"lib/weibo"
	"net/url"
	"strings"
)

func main() {
	cookies, err := weibo.ParserCookie(*system.Account)
	system.OutputAllErros(err, true)
	
	if cookies == nil {
		username := url.QueryEscape(*system.Account)
		username = base64.StdEncoding.EncodeToString([]byte(username))
		
		weibo.Login(username, *system.Password)
		cookies, err = weibo.ParserCookie(*system.Account)
		system.OutputAllErros(err, true)
	}
	//多张图
	imgs :=make([]string,0)
	if strings.Contains(*system.ImageUrl,","){
		imgUrls  :=strings.Split(*system.ImageUrl,",")
		for _,url :=range imgUrls{
			imgs=append(imgs, weibo.UploadImg(url, cookies, *system.Nickname))
		}
	}else{
		imgs=append(imgs, weibo.UploadImg(*system.ImageUrl, cookies, *system.Nickname))
	}

	//http上传到远端
	if *system.PostUrl !="" {
		content :=""
		for _,url :=range imgs{
			content = content +"<img src='"+url+"' />"
		}
		postData := make(map[string]string)
		postData["title"] = *system.Title
		postData["content"] = content

		uploadResult, errs := http.Request.Post(*system.PostUrl, postData,false,nil)
		system.OutputAllErros(errs, true)
		color.Green("result:",uploadResult)
	}else{
		color.Green("result:",imgs)
	}

}
