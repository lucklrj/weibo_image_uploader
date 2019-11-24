package system

import (
	"github.com/fatih/color"
	"os"
	"reflect"
)

func OutputAllErros(errs interface{}, end bool) {
	if errs != nil {
		if reflect.TypeOf(errs).String() == "[]error" && len(errs.([]error)) > 0 {
			for _, err := range errs.([]error) {
				color.Red(err.Error())
			}
		} else if reflect.TypeOf(errs).String() == "error" {
			color.Red(errs.(error).Error())
		} else {
			color.Red(errs.(error).Error())
		}
		if end == true {
			os.Exit(0)
		}
	}
}
