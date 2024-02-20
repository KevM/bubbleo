package color

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/kevm/bubbleo/navstack"
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
		pop := cmdize(navstack.PopNavigation{})
		selected := cmdize(ColorSelected{m.RGB})
		return m, tea.Sequence(pop, selected)
	case "ctrl+c":
		return m, tea.Quit
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
