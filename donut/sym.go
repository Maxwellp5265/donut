package donut

import (
	"image/color"
	"sync"

	"charm.land/lipgloss/v2"
)

// Symbol represents a pair of color and character.
type Symbol struct {
	color.RGBA
	rune
}

// String returns the formatted ASCII string representation of the Symbol.
func (s Symbol) String() string {
	str, ok := getCache(s)

	if !ok {
		style := getStyle()
		str = style.Foreground(s.RGBA).Render(string(s.rune))

		setCache(s, str)
		putStyle(style)
	}

	return str
}

var stringCache struct {
	syms map[Symbol]string
	sync.RWMutex
}

func getCache(s Symbol) (string, bool) {
	stringCache.RLock()

	str, ok := stringCache.syms[s]

	stringCache.RUnlock()

	return str, ok
}

func setCache(s Symbol, str string) {
	stringCache.Lock()

	stringCache.syms[s] = str

	stringCache.Unlock()
}

var stylePool sync.Pool

func init() {
	stringCache.syms = make(map[Symbol]string)
	stylePool.New = func() any {
		return new(lipgloss.NewStyle())
	}
}

func getStyle() *lipgloss.Style {
	return stylePool.Get().(*lipgloss.Style)
}

func putStyle(s *lipgloss.Style) {
	stylePool.Put(s)
}
