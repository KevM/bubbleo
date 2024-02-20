package window

import tea "github.com/charmbracelet/bubbletea"

type Model struct {
	Width     int
	Height    int
	TopOffset int
}

func New(width, height, topOffset int) Model {
	return Model{
		Width:     width,
		Height:    height,
		TopOffset: topOffset,
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.Width = msg.Width
		m.Height = msg.Height
	}
	return m, nil
}

func (m Model) View() string {
	return ""
}
