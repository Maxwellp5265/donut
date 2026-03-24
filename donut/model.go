package donut

import (
	"time"
)

// Model represents the state of the donut animation.
type Model struct {
	fps
	grid  []Symbol
	depth []float64

	h, w int
	a, b int

	color byte
	emoji bool
	mute  bool
}

// NewModel creates a new Model with the specified height and width.
func NewModel(h, w int) *Model {
	return &Model{
		h:     h,
		w:     w,
		depth: make([]float64, DonutSize),
		grid:  make([]Symbol, DonutSize),
		color: pickColor(),
	}
}

// Size returns the total number of cells in the grid.
func (m Model) Size() int {
	return m.h * m.w
}

// Resize updates the dimensions of the model.
func (m *Model) Resize(h, w int) {
	m.h = max(1, h)
	m.w = max(1, w)
}

// step updates the model's state for the next frame of the animation.
func (m *Model) step() {
	clear(m.depth)
	clear(m.grid)

	cars, dw, xdiv := ASCII, DonutW, 2
	if m.emoji {
		cars, dw, xdiv = Emoji, DonutW*2, 4
	}

	for j := range sinTable7 {
		for i := range sinTable2 {
			si, ci := sinTable2[i], cosTable2[i]
			sj, cj := sinTable7[j], cosTable7[j]
			sa, ca := sinTable4[m.a], cosTable4[m.a]
			sb, cb := sinTable2[m.b], cosTable2[m.b]
			cjp2 := cj + 2
			z := 1 / (si*cjp2*sa + sj*ca + 5)
			t := si*cjp2*ca - sj*sa

			x := dw/xdiv + int(30*z*(ci*cjp2*cb-t*sb))
			y := DonutH/2 + int(15*z*(ci*cjp2*sb+t*cb))

			lumi := byte(clamp(
				int(9*((sj*sa-si*cj*ca)*cb-si*cj*sa-sj*ca-ci*cj*sb)),
				0, 11,
			))

			ii := x + DonutW*y

			if z > m.depth[ii] {
				m.depth[ii] = z
				m.grid[ii] = Symbol{
					rune: cars[lumi],
					RGBA: blend(m.color, lumi, 0.5),
				}
			}
		}
	}

	m.a = (m.a + 1) % len(sinTable4)
	m.b = (m.b + 1) % len(sinTable2)
}

type fps struct {
	frames int
	last   time.Time
	cur    float64
}

// FPS returns the current frames per second.
func (f *fps) FPS() float64 {
	return f.cur
}

func (f *fps) sample(now time.Time) {
	f.frames++

	if f.last.IsZero() {
		f.last = now
	}

	dt := now.Sub(f.last)

	if dt >= time.Second {
		f.cur = float64(f.frames) / dt.Seconds()
		f.frames = 0
		f.last = now
	}
}

func clamp(x, α, ω int) int {
	return max(α, min(x, ω))
}
