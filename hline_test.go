package linedraw

import (
	"bytes"
	"testing"
)

func TestHLine(t *testing.T) {
	t.Run("should print \u2501 (a horizontal bold line)", func(t *testing.T) {
		boldLine := "\u2501"

		var output bytes.Buffer
		print(&output, boldLine)

		got := output.String()
		want := boldLine

		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}
