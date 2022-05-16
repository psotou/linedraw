package linedraw

import "fmt"

const (
	vlineLight = "\u2502"
	vlineBold  = "\u2503"
)

type vline struct {
	length int
}

func NewVLine(linelentgh int) *vline {
	return &vline{length: linelentgh}
}

func (h *vline) Light() {
	fmt.Print(lineLength(h.length, vlineLight))
}

func (h *vline) Bold() {
	fmt.Print(lineLength(h.length, vlineBold))
}
