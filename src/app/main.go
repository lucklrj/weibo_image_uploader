package main

import (
	"net/url"
	"encoding/base64"
	"github.com/parnurzeal/gorequest"
	"fmt"
	"sync"
	"lib/system"
	"strings"
	"github.com/tidwall/gjson"
	"lib/encrypt"
	"os"
)

var (
	Request HttpRequest
)

type HttpRequest struct {
	Request *gorequest.SuperAgent
	mu      sync.Mutex
}

func (h *HttpRequest) Get(url string) (body string, errs []error) {
	h.mu.Lock()
	_, body, errs = h.Request.Get(url).End()
	h.mu.Unlock()
	return body, errs
}
func (h *HttpRequest) Post(url string, postData map[string]string) (body string, errs []error) {
	h.mu.Lock()
	_, body, errs = h.Request.Post(url).Type("multipart").Send(postData).End()
	h.mu.Unlock()
	return body, errs
}

func init() {
	Request = HttpRequest{Request: gorequest.New()}
}
func main() {
	//处理用户名
	username := "sunny_lrj@yeaj.net"
	username = url.QueryEscape(username)
	username = base64.StdEncoding.EncodeToString([]byte(username))
	
	password := "123asd123"
	
	url := "http://login.sina.com.cn/sso/prelogin.php?entry=weibo&callback=sinaSSOController.preloginCallBack&su=" + username + "&rsakt=mod&checkpin=1&client=ssologin.js(v1.4.18)&_=1461819359582"
	html, errs := Request.Get(url)
	system.OutputAllErros(errs, true)
	
	//获取预登陆地址
	html = strings.Replace(html, "sinaSSOController.preloginCallBack(", "", -1)
	html = html[0 : len(html)-1]
	
	//解析预处理的数据
	/*{
		"retcode": 0,
		"servertime": 1530284135,
		"pcid": "gz-8a0121d67c090b79760432dde062834ad3d6",
		"nonce": "CFZW7J",
		"pubkey": "EB2A38568661887FA180BDDB5CABD5F21C7BFD59C090CB2D245A87AC253062882729293E5506350508E7F9AA3BB77F4333231490F915F6D63C55FE2F08A49B353F444AD3993CACC02DB784ABBB8E42A9B1BBFFFB38BE18D78E87A0E41B9B8F73A928EE0CCEE1F6739884B9777E4FE9E88A1BBE495927AC4A799B3181D6442443",
		"rsakv": "1330428213",
		"exectime": 5
	}*/
	htmlJson := gjson.Parse(html)
	retcode := htmlJson.Get("retcode").String()
	servertime := htmlJson.Get("servertime").String()
	pcid := htmlJson.Get("pcid").String()
	nonce := htmlJson.Get("nonce").String()
	pubkey := htmlJson.Get("pubkey").String()
	rsakv := htmlJson.Get("rsakv").String()
	exectime := htmlJson.Get("exectime").String()
	
	fmt.Println(retcode, servertime, pcid, nonce, pubkey, rsakv, exectime)
	
	//开始登陆
	//'entry' => 'weibo',
	//'gateway' => '1',
	//	'from' => '',
	//	'savestate' => '7',
	//	'useticket' => '1',
	//	'pagerefer' => '',
	//	'vsnf' => '1',
	//	'su' => base64_encode(urlencode($this->username)),
	//'service' => 'miniblog',
	//'servertime' => $data['servertime'],
	//	'nonce' => $data['nonce'],
	//	'pwencode' => 'rsa2',
	//'rsakv' => $data['rsakv'],
	//// 加密用户登入密码
	//	'sp' => bin2hex(rsa_encrypt($msg, '010001', $data['pubkey'])),
	//'sr' => '1440*900',
	//'encoding' => 'UTF-8',
	//// 该参数为加载 preLogin 页面到提交登入表单的间隔时间
	//// 此处使用 float 是为了兼容 32 位系统
	//'prelt' => (int)round((microtime(true) - $data['preloginTime']) * 1000),
	//'url' => 'http://weibo.com/ajaxlogin.php?'
	//. 'framelogin=1&callback=parent.sinaSSOController.feedBackUrlCallBack',
	//'returntype' => 'META'
	
	//pubkeyByte, err := hex.DecodeString(pubkey)
	//if err != nil {
	//	fmt.Println(err)
	//	os.Exit(0)
	//}
	//encrypt.PublicKey = pubkeyByte
	msg := []byte(servertime + "\t" + nonce + "\n" + password)
	sp, err := encrypt.RsaEncrypt(msg)
	fmt.Println(err)
	fmt.Println(sp)
	
	os.Exit(0)
	postUrl := "http://login.sina.com.cn/sso/login.php?client=ssologin.js(v1.4.18)"
	postData := make(map[string]string)
	postData["entry"] = "weibo"
	postData["gateway"] = "1"
	postData["from"] = ""
	postData["savestate"] = "7"
	postData["useticket"] = "1"
	postData["pagerefer"] = ""
	postData["vsnf"] = "1"
	postData["su"] = username
	postData["servertime"] = servertime
	postData["nonce"] = nonce
	postData["pwencode"] = "rsa2"
	postData["rsakv"] = rsakv
	postData["sp"] = string(sp)
	postData["sr"] = "1440*900"
	postData["encoding"] = "UTF-8"
	postData["prelt"] = "0"
	postData["url"] = "http://weibo.com/ajaxlogin.php?framelogin=1&callback=parent.sinaSSOController.feedBackUrlCallBack"
	postData["returntype"] = "META"
	postHtml, errs := Request.Post(postUrl, postData)
	system.OutputAllErros(errs, true)
	
	fmt.Println(postHtml)
	
}
