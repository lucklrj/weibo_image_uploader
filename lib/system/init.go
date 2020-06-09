package system

import (
	"io/ioutil"
	"os"
)

func GetCookName() string {

	_, err := ioutil.ReadDir("cookie")
	if err != nil {
		os.Mkdir("cookie", 0777)
	}
	return "cookie/default.txt"

}
