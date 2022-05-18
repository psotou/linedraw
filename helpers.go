package linedraw

import (
	"fmt"
	"io"
	"strings"
)

func print(w io.Writer, line string) {
	fmt.Fprintf(w, line)
}

func lineLength(length int, lineType string) string {
	var linelen []string
	for i := 0; i < length; i++ {
		linelen = append(linelen, lineType)
	}
	return fmt.Sprintf("%v", strings.Join(linelen, ""))
}
