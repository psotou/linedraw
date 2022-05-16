package linedraw

import "os"

const (
	dotLight = "\u2504"
	dotBold  = "\u2505"
)

type dotline struct {
	length int
}

func NewDotLine(linelentgh int) *dotline {
	return &dotline{length: linelentgh}
}

func (h *dotline) Light() {
	print(os.Stdout, lineLength(h.length, dotLight))
}

func (h *dotline) Bold() {
	print(os.Stdout, lineLength(h.length, dotBold))
}
