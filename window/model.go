package window

import tea "github.com/charmbracelet/bubbletea"

type Model struct {
	Width      int
	Height     int
	TopOffset  int
	SideOffset int
}

func New(width, height, topOffset int, sideOffset int) Model {
	return Model{
		Width:      width,
		Height:     height,
		TopOffset:  topOffset,
		SideOffset: sideOffset,
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

func (m Model) GetWindowSizeMsg() tea.WindowSizeMsg {
	return tea.WindowSizeMsg{Width: m.Width - m.SideOffset, Height: m.Height - m.TopOffset}
}
