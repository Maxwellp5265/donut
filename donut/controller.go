package donut

import (
	"time"

	tea "charm.land/bubbletea/v2"
)

// 15ms means 60FPS.
const TimeStep = 15 * time.Millisecond

// Update handles incoming messages and updates the model accordingly.
func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tickMsg:
		m.step()
		m.sample(time.Time(msg))

		return m, tick()
	case tea.KeyMsg:
		switch msg.String() {
		case "c":
			m.color = nextColor(m.color)
		case "a":
			m.emoji = false
		case "e":
			m.emoji = true
		case "m":
			m.mute = !m.mute
		case "q", "ctrl+c":
			return m, tea.Quit
		}
	case tea.WindowSizeMsg:
		m.Resize(msg.Height, msg.Width)
	}

	return m, nil
}

// Init initializes the model and returns an initial command to start the animation.
func (m Model) Init() tea.Cmd {
	return tick()
}

type tickMsg time.Time

func tick() tea.Cmd {
	return tea.Tick(TimeStep, func(t time.Time) tea.Msg {
		return tickMsg(t)
	})
}
