// Package donut implements a rotating ASCII donut using the Bubble Tea
// framework.
// See https://www.a1k0n.net/2011/07/20/donut-math.html for the original algorithm.
package donut

// DonutW and DonutH define the width and height of the donut grid with some
// margin to avoid clipping over the edges of the rendering canvas.
const (
	DonutW    = 84 // 80 + margin
	DonutH    = 26 // 22 + margin
	DonutSize = DonutW * DonutH
)

var (
	// Character set for rendering the donut, ordered from darkest to brightest.
	// There are intentionally left editable.
	ASCII = []rune(".,-~:;=!*#$@")
	Emoji = []rune("🍗🍫🍩🥨🍕🥔🍔🧀🥐🥯🍞🍚")
)
