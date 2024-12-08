package utils

import "github.com/fatih/color"

var (
	BlueColor    = color.New(color.FgBlue).SprintFunc()
	GreenColor   = color.New(color.FgGreen).SprintFunc()
	RedColor     = color.New(color.FgRed).SprintFunc()
	YellowColor  = color.New(color.FgYellow).SprintFunc()
	MagentaColor = color.New(color.FgMagenta).SprintFunc()
)
