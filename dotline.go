package linedraw

import "fmt"

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
	fmt.Print(lineLength(h.length, dotLight))
}

func (h *dotline) Bold() {
	fmt.Print(lineLength(h.length, dotBold))
}
