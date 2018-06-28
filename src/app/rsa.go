package main

import (
	"fmt"
	"lib/encrypt"
	"encoding/base64"
)

func main() {
	str := "pDP4Fq+GXGn9qsARqBrdnC1G6alWtT3gDwRMkQ+IWka973TfL8IjRqncpdnP2RzFJB92y/PV4atQcXIAo0rq8Zhw4ozQF506OpSaYHw7715uso283x9hA/GHhys8KE91Mv7H8SKc8Fez9X4jdldr5d1HHAPmLUC/gUj4WV16zkY="
	data, err := base64.StdEncoding.DecodeString(str)
	origData, err := encrypt.RsaDecrypt(data)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(origData))
}
