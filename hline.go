package linedraw

import (
	"fmt"
	"io"
	"os"
	"strings"
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

func print(w io.Writer, line string) {
	fmt.Fprintf(w, line)
}

func lineLength(length int, lineType string) string {
	var linelen []string
	for i := 0; i < length+1; i++ {
		linelen = append(linelen, lineType)
	}
	return fmt.Sprintf("%v\n", strings.Join(linelen, ""))
}
