package weibo

import (
	"encoding/base64"
	"lib/system"
	"os"
	"io/ioutil"
	"lib/http"
	"regexp"
	"github.com/tidwall/gjson"
	"errors"
	"strconv"
	"math/rand"
)

func UploadImg(filePath string) string {
	file, err := os.OpenFile(filePath, os.O_RDONLY, os.ModePerm)
	defer file.Close()
	system.OutputAllErros(err, true)
	filContent, err := ioutil.ReadAll(file)
	
	imgUploadUrl := "http://picupload.service.weibo.com/interface/pic_upload.php?mime=image%2Fjpeg&data=base64&url=0&markpos=1&logo=&nick=0&marks=1&app=miniblog&cb=http://weibo.com/aj/static/upimgback.html?_wv=5&callback=STK_ijax_1111";
	postData := make(map[string]string)
	postData["b64_data"] = base64.StdEncoding.EncodeToString(filContent)
	
	uploadResult, errs := http.Request.Post(imgUploadUrl, postData)
	system.OutputAllErros(errs, true)
	
	reg := regexp.MustCompile(`.*?(\{.*)`)
	respJsonMatchResult := reg.FindAllStringSubmatch(uploadResult, -1)
	if len(respJsonMatchResult) == 0 {
		system.OutputAllErros(errors.New("上传图片失败"), true)
	}
	pid := gjson.Parse(respJsonMatchResult[0][1]).Get("data.pics.pic_1.pid").String()
	return getImgUrl(pid)
}
func getImgUrl(pid string) string {
	/*
	 *(($pid[9] === 'w' ? (crc32($pid) & 3) : (hexdec(substr($pid, 19, 2)) & 0xf)) + 1)
	 * 然而当前能访问的 cdn 编号只有 1 ~ 4，而且基本上任意的
	 * cdn 编号都能访问到同一资源，所以根据 pid 来判断 cdn 编号
	 * 当前实际上没啥意义了，有些实现甚至直接写死 cdn 编号
	 */
	
	return "https://ws" + strconv.Itoa(rand.Intn(4)) + ".sinaimg.cn/large/" + pid
}
