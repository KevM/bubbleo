package shell

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/kevm/bubbleo/breadcrumb"
	"github.com/kevm/bubbleo/navstack"
	"github.com/kevm/bubbleo/window"
)

type Model struct {
	Navstack   *navstack.Model
	Breadcrumb breadcrumb.Model
}

func New() Model {
	w := window.New(120, 30, 2, 0)
	ns := navstack.New(&w)
	bc := breadcrumb.New(&ns)

	return Model{
		Navstack:   &ns,
		Breadcrumb: bc,
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	// cmds := []tea.Cmd{}

	// updatedbc, cmd := m.Breadcrumb.Update(msg)
	// m.Breadcrumb = updatedbc.(breadcrumb.Model)
	// cmds = append(cmds, cmd)

	cmd := m.Navstack.Update(msg)
	// cmds = append(cmds, cmd)

	return m, cmd //tea.Batch(cmds...)
}

func (m Model) View() string {
	bc := m.Breadcrumb.View()
	nav := m.Navstack.View()
	return lipgloss.NewStyle().Render(bc, nav)
}
