package donut

import (
	"fmt"
	"testing"
)

func BenchmarkView(b *testing.B) {
	sizes := []struct {
		w, h int
	}{
		{80, 24},
		{120, 40},
		{160, 48},
		{320, 90},
		{480, 135},
		{640, 270},
		{3829, 700},
	}

	for _, sz := range sizes {
		for _, emoji := range []bool{false, true} {
			name := fmt.Sprintf("%dx%d/emoji=%v", sz.w, sz.h, emoji)
			b.Run(name, func(b *testing.B) {
				m := NewModel(sz.h, sz.w)
				m.emoji = emoji

				for range 10 {
					m.step()
				}

				for b.Loop() {
					m.View()
				}
			})
		}
	}
}
