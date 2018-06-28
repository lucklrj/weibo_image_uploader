package system

import (
	"github.com/fatih/color"
	"os"
)

func OutputAllErros(errs []error, end bool) {
	for _, err := range (errs) {
		color.Red(err.Error())
	}
	if end == true {
		os.Exit(0)
	}
}
