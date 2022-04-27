package log

import (
	"fmt"

	"github.com/gookit/color"
)

var (
	red    = color.FgRed.Render
	cyan   = color.FgLightCyan.Render
	yellow = color.Yellow.Render
	blue   = color.FgLightBlue.Render
)

func Section(value string) {
	fmt.Printf("%s\n", cyan(value))
	fmt.Print(red("--------------------\n"))
}

func Info(key string, value string) {
	fmt.Printf("%s %s %s", yellow(key), red(":"), blue(value))
	fmt.Println()
}
