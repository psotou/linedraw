package linedraw

import (
	"bytes"
	"testing"
)

func TestPrint(t *testing.T) {
	tt := []struct {
		description string
		want        string
	}{
		{"should print a horizontal bold line \u2501", "\u2501"},
		{"should print a horizontal light line \u2500", "\u2500"},
	}

	for _, tc := range tt {
		t.Run(tc.description, func(t *testing.T) {
			var output bytes.Buffer
			print(&output, tc.want)
			if tc.want != output.String() {
				t.Errorf("got  %vwant %v", output.String(), tc.want)
			}
		})
	}
}

func TestLineLenght(t *testing.T) {
	tt := []struct {
		description string
		linetype    string
		want        string
	}{
		{"should return a dotted vertical bold line of lenght 4", "\u2507", "\u2507\u2507\u2507\u2507\n"},
		{"should return a dotted vertical light line of lenght 4", "\u2506", "\u2506\u2506\u2506\u2506\n"},
	}

	for _, tc := range tt {
		t.Run(tc.description, func(t *testing.T) {
			got := lineLength(4, tc.linetype)
			if got != tc.want {
				t.Errorf("\ngot  %vwant %v", got, tc.want)
			}
		})
	}
}
