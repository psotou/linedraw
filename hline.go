package linedraw

import (
	"os"
)

const (
	hlineLight = "\u2500"
	hlineBold  = "\u2501"
)

type hline struct {
	length int
}

func NewHLine(linelentgh int) *hline {
	return &hline{length: linelentgh}
}

func (h *hline) Light() {
	print(os.Stdout, lineLength(h.length, hlineLight))
}

func (h *hline) Bold() {
	print(os.Stdout, lineLength(h.length, hlineBold))
}
