package output

import "github.com/fatih/color"

func PrintError(err error, errString string) {
	color.Red("%s: %s", errString, err.Error())
}
