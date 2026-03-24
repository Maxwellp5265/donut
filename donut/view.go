package donut

import (
	"strconv"
	"strings"

	tea "charm.land/bubbletea/v2"
)

// View renders the current state of the model as a string.
func (m *Model) View() tea.View {
	var sb strings.Builder

	// Preallocate the string builder.
	size := clamp(m.Size(), 32768, 155648)
	if m.emoji {
		size = max(size, 81920)
	}

	sb.Grow(size)

	if !m.mute {
		m.writeHeader(&sb)
	}

	// Emojis are rendered twice as big as ascii chars.
	vpad := (m.h - DonutH) / 2
	if m.emoji {
		vpad = (m.h - 2*DonutH) / 2
	}

	hpad, sp := (m.w-DonutW)/2, " "
	if m.emoji {
		hpad, sp = (m.w-2*DonutW)/2, "  "
	}

	for range vpad {
		sb.WriteByte('\n')
	}

	idx := func(r, c int) int {
		return r*DonutW + c
	}

	for i := range min(DonutH, m.h) {

		// Each row of emojis is rendered twice to maintain the aspect ratio.
		n := 1
		if m.emoji {
			n = 2
		}

		for range n {
			for range hpad {
				sb.WriteByte(' ')
			}

			for j := range min(DonutW, m.w) {
				if m.depth[idx(i, j)] == 0 {
					sb.WriteString(sp)
					continue
				}

				s := m.grid[idx(i, j)]
				sb.WriteString(s.String())
			}
			sb.WriteByte('\n')
		}
	}

	v := tea.NewView(sb.String())
	v.AltScreen = true

	return v
}

func (m *Model) writeHeader(sb *strings.Builder) {
	const (
		header0 = "fps [q]uit [m]ute [c]olor [e]moji"
		header1 = "fps [q]uit [m]ute [a]scii"
	)

	var buf [24]byte

	s := strconv.AppendFloat(buf[:0], m.FPS(), 'f', 0, 64)

	if len(s) < 2 {
		sb.WriteByte(' ')
	}

	sb.Write(s)

	if m.emoji {
		sb.WriteString(header1)
	} else {
		sb.WriteString(header0)
	}

	sb.WriteByte('\n')
}
