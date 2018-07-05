package system

import (
	"reflect"
	"github.com/fatih/color"
	"os"
	"fmt"
)

func OutputAllErros(errs interface{}, end bool) {
	if errs != nil {
		if reflect.TypeOf(errs).String() == "[]error" && len(errs.([]error)) > 0 {
			fmt.Println("in2")
			for _, err := range errs.([]error) {
				color.Red(err.Error())
			}
			if end == true {
				os.Exit(0)
			}
		} else if reflect.TypeOf(errs).String() == "error" {
			fmt.Println("in3")
			color.Red(errs.(error).Error())
			if end == true {
				os.Exit(0)
			}
		}
		
	}
}
