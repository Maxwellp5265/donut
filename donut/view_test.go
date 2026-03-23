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
		b.Run(
			fmt.Sprintf("%dx%d", sz.w, sz.h),

			func(b *testing.B) {
				m := NewModel(sz.h, sz.w)
				m.emoji = true

				for range 10 {
					m.step()
				}

				for b.Loop() {
					m.View()
				}
			},
		)
	}
}
