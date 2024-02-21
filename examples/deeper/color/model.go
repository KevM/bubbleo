package color

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/kevm/bubbleo/navstack"
	"github.com/kevm/bubbleo/utils"
)

type Model struct {
	RGB    string
	Sample string
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "esc":
			return m, utils.Cmdize(navstack.PopNavigation{})
		case "enter":
			pop := utils.Cmdize(navstack.PopNavigation{})
			selected := utils.Cmdize(ColorSelected{RGB: m.RGB})
			return m, tea.Sequence(pop, selected)
		}
	}

	return m, nil
}

func (m Model) View() string {
	sample := lipgloss.NewStyle().
		Foreground(lipgloss.Color(m.RGB)).
		Render(m.Sample)

	return sample + "\n\n\n\n\n" + "enter: select, esc: back\n"
}
