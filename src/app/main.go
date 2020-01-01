package main

import (
	"encoding/base64"
	"github.com/fatih/color"
	imgtype "github.com/shamsher31/goimgtype"
	"gopkg.in/gographics/imagick.v3/imagick"

	"io/ioutil"
	"lib/http"
	"lib/system"
	"lib/weibo"
	"net/url"
	"os"
	"strings"
)

func main() {
	cookies, err := weibo.ParserCookie(*system.Account)
	system.OutputAllErros(err, true)

	var newLogin bool
	if cookies == nil {
		newLogin = true
	} else {
		//查看cookie是否已失效
		if weibo.Ping(cookies) == false {
			color.Red("cookie已失效,开始重新登录")
			weibo.DeleteCookie(*system.Account)
			newLogin = true
		} else {
			color.Green("cookie还在有效时间范围内")
			newLogin = false
		}
	}
	if newLogin == true {
		username := url.QueryEscape(*system.Account)
		username = base64.StdEncoding.EncodeToString([]byte(username))

		weibo.Login(username, *system.Password)
		cookies, err = weibo.ParserCookie(*system.Account)
		system.OutputAllErros(err, true)
	}
	//多张图
	remoteUrl := ""
	imgs := make([]string, 0)
	if strings.Contains(*system.ImageUrl, ",") {
		imgUrls := strings.Split(*system.ImageUrl, ",")
		for _, url := range imgUrls {
			remoteUrl = weibo.UploadImg(url, cookies, *system.Nickname)
			if remoteUrl != "" {
				imgs = append(imgs, remoteUrl)
			}

		}
	} else if *system.ImageUrl != "" {
		remoteUrl = weibo.UploadImg(*system.ImageUrl, cookies, *system.Nickname)
		if remoteUrl != "" {
			imgs = append(imgs, remoteUrl)
		}
	}
	if *system.IsCutBottom == "y" {
		imagick.Initialize()
		defer imagick.Terminate()
	}

	//兼容目录
	if *system.ImageDir != "" {
		files, err := ioutil.ReadDir(*system.ImageDir)
		if err != nil {
			color.Red(err.Error())
		} else {
			for _, file := range files {
				filePath := *system.ImageDir + file.Name()
				_, err := imgtype.Get(filePath)

				if err != nil {
					color.Red(filePath + " 不是图片")
				} else {
					if *system.IsCutBottom == "y" {
						mw := imagick.NewMagickWand()
						mw.ReadImage(filePath)
						imageWidth := mw.GetImageWidth()
						imageHeight := mw.GetImageHeight()

						mw.CropImage(imageWidth, imageHeight-30, 0, 0)
						mw.SetImageCompressionQuality(95)
						mw.WriteImage(filePath)
						mw.Destroy()
					}
					remoteUrl = weibo.UploadImg(filePath, cookies, *system.Nickname)
					if remoteUrl != "" {
						imgs = append(imgs, remoteUrl)
					}
				}
			}
		}
	}

	if len(imgs) == 0 {
		color.Red("没有文件上传")
		os.Exit(0)
	}
	//http上传到远端
	if *system.PostUrl != "" {
		content := ""
		for _, url := range imgs {
			content = content + "<img src='" + url + "' />"
		}
		postData := make(map[string]string)
		postData["title"] = *system.Title
		postData["content"] = content

		uploadResult, errs := http.Request.Post(*system.PostUrl, postData, false, nil)
		system.OutputAllErros(errs, true)
		color.Green("result:", uploadResult)
	} else {
		color.Green("result:", imgs)
	}

}
