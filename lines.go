package linedraw

import (
	"os"
)

const (
	hlineLight       = "\u2500"
	hlineBold        = "\u2501"
	vlineLight       = "\u2502"
	vlineBold        = "\u2503"
	hdottedLineLight = "\u2504"
	hdottedLineBold  = "\u2505"
	vdottedLineLight = "\u2506"
	vdottedLineBold  = "\u2507"
)

type lines struct {
	Horizontal       lineProps
	Vertical         lineProps
	HorizontalDotted lineProps
	VerticalDotted   lineProps
}

type lineProps struct {
	length int
	light  string
	bold   string
}

// NewLine takes the length of the line to be printed to Stdout
func NewLine(length int) *lines {
	return &lines{
		Horizontal:       lineProps{length: length, light: hlineLight, bold: hlineBold},
		Vertical:         lineProps{length: length, light: vlineLight, bold: vlineBold},
		HorizontalDotted: lineProps{length: length, light: hdottedLineLight, bold: hdottedLineBold},
		VerticalDotted:   lineProps{length: length, light: vdottedLineLight, bold: vdottedLineBold},
	}
}

func (l *lineProps) Bold() {
	print(os.Stdout, lineLength(l.length, l.bold))
}

func (l *lineProps) Light() {
	print(os.Stdout, lineLength(l.length, l.light))
}
