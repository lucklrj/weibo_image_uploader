package system

import (
	"flag"
	"github.com/fatih/color"
	"io/ioutil"
	"os"
	"strings"
)

var (
	ImageUrl = flag.String("url", "", "图片位置")
	ImageDir = flag.String("dir", "", "图片文件夹")
	Account  = flag.String("account", "", "微博账号")
	Password = flag.String("password", "", "微博密码")
	Nickname = flag.String("nickname", "", "微博昵称")

	PostUrl = flag.String("post_url", "", "传送到远端地址")
	Title   = flag.String("title", "", "标题")
)

func init() {

	flag.Parse()
	if *ImageUrl == "" && *ImageDir == "" {
		color.Red("缺少url或dir")
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

	if *ImageDir != "" && strings.HasSuffix(*ImageDir, "/") == false {
		*ImageDir = *ImageDir + "/"
	}
}

func GetCookName() string {

	_, err := ioutil.ReadDir("cookie")
	if err != nil {
		os.Mkdir("cookie", 0777)
	}
	if *Account == "" {
		return "cookie/cookie.txt"
	} else {
		return "cookie/" + *Account + ".txt"
	}
}
