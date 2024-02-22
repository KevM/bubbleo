package breadcrumb

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/kevm/bubbleo/navstack"
	"github.com/kevm/bubbleo/styles"
)

type Model struct {
	Navstack   *navstack.Model
	FrameStyle lipgloss.Style
}

func New(n *navstack.Model) Model {
	return Model{
		Navstack:   n,
		FrameStyle: styles.BreadCrumbFrameStyle.Copy().Width(120).Height(1),
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return m, nil
}

func (m Model) View() string {
	b := strings.Builder{}

	for i, c := range m.Navstack.StackSummary() {
		if i != 0 {
			b.WriteString(" > ")
		}
		b.WriteString(c)
	}
	crumbs := b.String()
	return m.FrameStyle.Render(crumbs)
}
