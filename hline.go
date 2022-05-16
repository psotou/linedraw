package linedraw

import (
	"fmt"
	"strings"
)

const (
	lineLight = "\u2500"
	lineBold  = "\u2501"
)

type hline struct {
	length int
}

func NewHLine(linelentgh int) *hline {
	return &hline{length: linelentgh}
}

func (h *hline) Light() {
	fmt.Print(lineLength(h.length, lineLight))
}

func (h *hline) Bold() {
	fmt.Print(lineLength(h.length, lineBold))
}

func lineLength(length int, lineType string) string {
	var linelen []string
	for i := 0; i < length+1; i++ {
		linelen = append(linelen, lineType)
	}
	return fmt.Sprintf("%v\n", strings.Join(linelen, ""))
}
