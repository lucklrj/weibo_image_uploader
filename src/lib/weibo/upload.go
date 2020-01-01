package weibo

import (
	"encoding/base64"
	"errors"
	"github.com/fatih/color"
	"github.com/tidwall/gjson"
	"io/ioutil"
	"lib/http"
	"lib/system"
	"math/rand"
	http2 "net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func UploadImg(filePath string, cookies []*http2.Cookie, nickname string) string {
	fileContent := []byte("")
	var isRemote bool
	if strings.HasPrefix(filePath, "http://") == true || strings.HasPrefix(filePath, "https://") == true {
		isRemote = true
		fileContentString, errs := http.Request.Get(filePath, nil)
		fileContent = []byte(fileContentString)
		system.OutputAllErros(errs, true)
	} else {
		isRemote = false
		file, err := os.OpenFile(filePath, os.O_RDONLY, os.ModePerm)
		defer file.Close()
		system.OutputAllErros(err, true)
		fileContent, _ = ioutil.ReadAll(file)
	}

	imgUploadUrl := "https://picupload.service.weibo.com/interface/pic_upload." + "php?mime=image%2Fjpeg&data=base64&url=0&markpos=1&logo=&nick=" + nickname + "&marks=1&app=miniblog&cb=http://weibo.com/aj/static/upimgback.html?_wv=5&callback=STK_ijax_1111"
	postData := make(map[string]string)
	postData["b64_data"] = base64.StdEncoding.EncodeToString([]byte(fileContent))

	retryMaxNum := 5

	for i := 0; i <= retryMaxNum; i++ {
		uploadResult, err := http.Request.Post(imgUploadUrl, postData, false, cookies)
		if err != nil {
			color.Red(err.Error())
			color.Green("开始重试：" + filePath)
			continue
		} else {
			reg := regexp.MustCompile(`.*?(\{.*)`)
			respJsonMatchResult := reg.FindAllStringSubmatch(uploadResult, -1)

			code := gjson.Parse(respJsonMatchResult[0][1]).Get("code").String()
			if code != "A00006" {
				system.OutputAllErros(errors.New(filePath+":上传图片失败"), false)
				color.Green("开始重试：" + filePath)
			}
			pid := gjson.Parse(respJsonMatchResult[0][1]).Get("data.pics.pic_1.pid").String()
			if pid != "" {
				color.Green(filePath + " 上传成功")
				if isRemote == false {
					os.Remove(filePath)
				}
			}
			return getImgUrl(pid)
		}
	}
	return ""
}
func getImgUrl(pid string) string {

	/*
	 *(($pid[9] === 'w' ? (crc32($pid) & 3) : (hexdec(substr($pid, 19, 2)) & 0xf)) + 1)
	 * 然而当前能访问的 cdn 编号只有 1 ~ 4，而且基本上任意的
	 * cdn 编号都能访问到同一资源，所以根据 pid 来判断 cdn 编号
	 * 当前实际上没啥意义了，有些实现甚至直接写死 cdn 编号
	 */
	return "https://ws" + strconv.Itoa(rand.Intn(3)+1) + ".sinaimg.cn/large/" + pid
}
