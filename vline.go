package linedraw

import "os"

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
	print(os.Stdout, lineLength(h.length, vlineLight))
}

func (h *vline) Bold() {
	print(os.Stdout, lineLength(h.length, vlineBold))
}
