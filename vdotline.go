package linedraw

import "fmt"

const (
	vdotLight = "\u2506"
	vdotBold  = "\u2507"
)

type vdotline struct {
	length int
}

func NewVDotLine(linelentgh int) *vdotline {
	return &vdotline{length: linelentgh}
}

func (h *vdotline) Light() {
	fmt.Print(lineLength(h.length, vdotLight))
}

func (h *vdotline) Bold() {
	fmt.Print(lineLength(h.length, vdotBold))
}
