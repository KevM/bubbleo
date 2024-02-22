package shell

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/kevm/bubbleo/breadcrumb"
	"github.com/kevm/bubbleo/navstack"
	"github.com/kevm/bubbleo/utils"
	"github.com/kevm/bubbleo/window"
)

type Model struct {
	Navstack   *navstack.Model
	Breadcrumb breadcrumb.Model

	window *window.Model
}

func New() Model {
	w := window.New(120, 30, 0, 0)
	ns := navstack.New(&w)
	bc := breadcrumb.New(&ns)

	return Model{
		Navstack:   &ns,
		Breadcrumb: bc,

		window: &w,
	}
}

func (m Model) Init() tea.Cmd {

	w, h := m.Breadcrumb.FrameStyle.GetFrameSize()
	m.window.SideOffset = w
	m.window.TopOffset = h

	return utils.Cmdize(m.window.GetWindowSizeMsg())
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	cmd := m.Navstack.Update(msg)
	return m, cmd
}

func (m Model) View() string {
	bc := m.Breadcrumb.View()
	nav := m.Navstack.View()
	return lipgloss.NewStyle().Render(bc, nav)
}
