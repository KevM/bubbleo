package color

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/kevm/bubbleo/menu"
)

type Model struct {
	RGB    string
	Sample string
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	switch msg.(tea.KeyMsg).String() {
	case "esc":
		return m, cmdize(menu.DismissSelected{})
	case "enter":
		return m, tea.Batch(cmdize(ColorSelected{RGB: m.RGB}), cmdize(menu.DismissSelected{}))
	}

	return m, nil
}

func (m Model) View() string {
	sample := lipgloss.NewStyle().
		Foreground(lipgloss.Color(m.RGB)).
		Render(m.Sample)

	return sample + "\n\n\n\n\n" + "enter: select, esc: back\n"
}

func cmdize[T any](t T) tea.Cmd {
	return func() tea.Msg {
		return t
	}
}
