// Package Breadcrumb is a component that consumes a pointer ot a navstack.Model
// and renders a breadcrumb trail. It is used to give the user a sense of where
// they are in the application.
package breadcrumb

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/kevm/bubbleo/navstack"
	"github.com/kevm/bubbleo/styles"
)

type BreadcrumbStyles struct {
	Frame     lipgloss.Style
	Delimiter string
}

type Model struct {
	Navstack *navstack.Model
	Styles   BreadcrumbStyles
}

func New(n *navstack.Model) Model {
	return Model{
		Navstack: n,
		Styles:   DefaultStyles(),
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
			b.WriteString(m.Styles.Delimiter)
		}
		b.WriteString(c)
	}
	crumbs := b.String()
	return m.Styles.Frame.Render(crumbs)
}

func DefaultStyles() BreadcrumbStyles {
	return BreadcrumbStyles{
		Frame:     styles.BreadCrumbFrameStyle,
		Delimiter: " > ",
	}
}
