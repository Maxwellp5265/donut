package donut

import (
	"image/color"
	"strings"
	"testing"
)

func TestSymbolString(t *testing.T) {
	sym := Symbol{RGBA: color.RGBA{}, rune: '$'}
	str := sym.String()

	if str == "" {
		t.Error("String() returned empty string")
	}

	if len(str) == 1 || !strings.Contains(str, "$") {
		t.Errorf("String() = %q, want a color string containing '$'", str)
	}
}
