package system

import (
	"flag"
	"github.com/fatih/color"
	"os"
)

var (
	ImageUrl = flag.String("url", "", "图片位置")
	Account  = flag.String("account", "", "微博账号")
	Password = flag.String("password", "", "微博密码")
	Nickname = flag.String("nickname", "", "微博昵称")
)

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

func GetCookName() string {
	if *Account == "" {
		return "cookie.txt"
	} else {
		return *Account + ".txt"
	}
}
