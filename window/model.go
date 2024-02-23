// Package Windows holds the dimensions of the container window.
// The offsets communicate parts of the window that are already in use.
// This component is used by shell to adjust child components based on the window size minus the offsets.
package window

import tea "github.com/charmbracelet/bubbletea"

type Model struct {
	Width      int
	Height     int
	TopOffset  int
	SideOffset int
}

// New creates a new window model the dimensions are usualy the starter default with
// future tea.WindowSizeMsg messages updating the height and width dimensions.
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

// Update updates the window model with the new window size.
func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.Width = msg.Width
		m.Height = msg.Height
	}
	return m, nil
}

// There is no default view.
func (m Model) View() string {
	return ""
}

// GetWindowSizeMsg returns a tea.WindowSizeMsg with the current window size minus the offsets.
func (m Model) GetWindowSizeMsg() tea.WindowSizeMsg {
	return tea.WindowSizeMsg{Width: m.Width - m.SideOffset, Height: m.Height - m.TopOffset}
}
