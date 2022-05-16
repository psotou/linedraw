package linedraw

import "os"

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
	print(os.Stdout, lineLength(h.length, vdotLight))
}

func (h *vdotline) Bold() {
	print(os.Stdout, lineLength(h.length, vdotBold))
}
