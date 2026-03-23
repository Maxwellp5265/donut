package donut

import (
	"fmt"
	"testing"
	"time"
)

func TestNewModel(t *testing.T) {
	m := NewModel(10, 20)

	if m.h != 10 || m.w != 20 {
		t.Errorf("expected h=10, w=20, got h=%d, w=%d", m.h, m.w)
	}
}

func TestModelResize(t *testing.T) {
	m := NewModel(5, 5)
	m.Resize(8, 12)

	if m.h != 8 || m.w != 12 {
		t.Errorf("expected h=8, w=12, got h=%d, w=%d", m.h, m.w)
	}
}

func TestModelSize(t *testing.T) {
	m := NewModel(7, 9)

	if m.Size() != 63 {
		t.Errorf("expected size 63, got %d", m.Size())
	}
}

func TestClamp(t *testing.T) {
	tests := []struct {
		x, a, w, want int
	}{
		{5, 1, 10, 5},
		{0, 1, 10, 1},
		{15, 1, 10, 10},
		{10, 10, 10, 10},
	}
	for _, tt := range tests {
		got := clamp(tt.x, tt.a, tt.w)

		if got != tt.want {
			t.Errorf("clamp(%d, %d, %d) = %d; want %d", tt.x, tt.a, tt.w, got, tt.want)
		}
	}
}

func TestBlendColors(t *testing.T) {
	cid := 1
	lum := 2

	out := blend(cid, lum, 0.5)

	if out.R != 188 || out.G != 124 || out.B != 60 || out.A != 255 {
		t.Errorf("unexpected blend result: %+v", out)
	}
}

func TestFPS(t *testing.T) {
	var f fps

	now := time.Now()
	f.sample(now)

	if f.FPS() != 0 {
		t.Errorf("expected FPS 0, got %f", f.FPS())
	}
	// Simulate 2 seconds and 9 frames
	f.last = now.Add(-2 * time.Second)
	f.frames = 9
	f.sample(now)

	if f.FPS() < 4.9 || f.FPS() > 5.1 {
		t.Errorf("expected FPS ~5, got %f", f.FPS())
	}
}

func TestModelStep(t *testing.T) {
	m := NewModel(10, 20)
	m.step()
	// After step, grid should be filled with Symbol, depth updated
	nonEmpty := 0
	for _, s := range m.grid {
		if s.rune != 0 && s.rune != ' ' {
			nonEmpty++
		}
	}
	if nonEmpty == 0 {
		t.Errorf("expected some non-empty grid cells after step")
	}
}

func BenchmarkStep(b *testing.B) {
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

				for b.Loop() {
					m.step()
				}
			},
		)
	}
}
