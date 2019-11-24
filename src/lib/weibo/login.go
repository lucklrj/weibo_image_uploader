package weibo

import (
	"crypto/rand"
	"crypto/rsa"
	"encoding/hex"
	"errors"
	"github.com/fatih/color"
	"github.com/tidwall/gjson"
	"io/ioutil"
	"lib/http"
	"lib/system"
	"math/big"
	http2 "net/http"
	"os"
	"regexp"
	"strings"
)

type preLoginData struct {
	servertime string
	pcid       string
	nonce      string
	pubkey     string
	rsakv      string
	exectime   string
	showpin    string
	su         string
	sp         string
	door       string
}

func Login(username string, password string) {
	color.Green("获取登陆信息")
	preLoginData := preLogin(username)
	if preLoginData.showpin == "1" {
		color.Green("请输入验证码")
	}

	color.Green("加密密码")
	sp := getPassword(preLoginData, password)
	preLoginData.su = username
	preLoginData.sp = sp

	color.Green("开始登陆系统")
	loginSubmit(preLoginData)
	color.Green("登陆成功")

}

func preLogin(username string) preLoginData {
	preLoginUrl := "http://login.sina.com.cn/sso/prelogin.php?entry=weibo&callback=sinaSSOController.preloginCallBack&su=" + username + "&rsakt=mod&checkpin=1&client=ssologin.js(v1.4.18)&_=1461819359582"
	result, errs := http.Request.Get(preLoginUrl, nil)
	system.OutputAllErros(errs, true)

	result = strings.Replace(result, "sinaSSOController.preloginCallBack(", "", -1)
	result = result[0 : len(result)-1]

	htmlJson := gjson.Parse(result)
	servertime := htmlJson.Get("servertime").String()
	pcid := htmlJson.Get("pcid").String()
	nonce := htmlJson.Get("nonce").String()
	pubkey := htmlJson.Get("pubkey").String()
	rsakv := htmlJson.Get("rsakv").String()
	exectime := htmlJson.Get("exectime").String()
	showpin := htmlJson.Get("exectime").String()

	return preLoginData{servertime: servertime, pcid: pcid, nonce: nonce, pubkey: pubkey, rsakv: rsakv, exectime: exectime, showpin: showpin}
}

func getPassword(preLoginData preLoginData, password string) string {
	int := new(big.Int)
	int.SetString(preLoginData.pubkey, 16)

	pub := rsa.PublicKey{
		N: int,
		E: 65537,
	}
	encryString := preLoginData.servertime + "\t" + preLoginData.nonce + "\n" + password

	encryResult, _ := rsa.EncryptPKCS1v15(rand.Reader, &pub, []byte(encryString))
	return hex.EncodeToString(encryResult)
}

func loginSubmit(preLoginData preLoginData) {
	postUrl := "http://login.sina.com.cn/sso/login.php?client=ssologin.js(v1.4.18)"
	postData := make(map[string]string)
	postData["entry"] = "weibo"
	postData["gateway"] = "1"
	postData["from"] = ""
	postData["savestate"] = "7"
	postData["useticket"] = "1"
	postData["pagerefer"] = ""
	postData["vsnf"] = "1"
	postData["su"] = preLoginData.su
	postData["servertime"] = preLoginData.servertime
	postData["nonce"] = preLoginData.nonce
	postData["pwencode"] = "rsa2"
	postData["rsakv"] = preLoginData.rsakv
	postData["sp"] = preLoginData.sp
	postData["sr"] = "1440*900"
	postData["encoding"] = "UTF-8"
	postData["prelt"] = "0"
	postData["url"] = "http://weibo.com/ajaxlogin.php?framelogin=1&callback=parent.sinaSSOController.feedBackUrlCallBack"
	postData["returntype"] = "META"
	//postData["door"] = "123"

	cookies := make([]*http2.Cookie, 0)
	postResult, errs := http.Request.Post(postUrl, postData, true, cookies)
	system.OutputAllErros(errs, true)

	reg := regexp.MustCompile(`(https:\/\/passport.*?)\'`)
	links := reg.FindAllStringSubmatch(postResult, -1)
	if len(links) == 0 {
		color.Red("账号，密码错误，无法登陆")
		os.Exit(0)
	}
	getTokenUrl := links[0][1]
	tokenResult, err := http.Request.Get(getTokenUrl, nil)
	system.OutputAllErros(err, true)

	//"uniqueid":"1575892744",
	reg = regexp.MustCompile(`("uniqueid":".*?",)`)
	matchResult := reg.FindAllStringSubmatch(tokenResult, -1)
	if len(matchResult) == 0 {
		system.OutputAllErros(errors.New("无法获取登陆token"), true)
	}
}

func ParserCookie(account string) ([]*http2.Cookie, error) {
	cookiePath := system.GetCookName()

	cookieFile, err := os.OpenFile(cookiePath, os.O_RDONLY|os.O_CREATE, os.ModePerm)
	cookieContent, err := ioutil.ReadAll(cookieFile)

	if len(cookieContent) == 0 {
		return nil, nil
	}
	cookieContentJson := gjson.ParseBytes(cookieContent).Value().(map[string]interface{})
	if len(cookieContentJson) == 0 {
		return nil, nil
	}
	cookies := make([]*http2.Cookie, 0)
	for key, val := range cookieContentJson {
		cookies = append(cookies, &http2.Cookie{
			Name:  key,
			Value: val.(string),
		})
	}
	color.Green("解析cookie成功")
	return cookies, err
}

func DeleteCookie(account string) {
	os.Remove(system.GetCookName())
}
