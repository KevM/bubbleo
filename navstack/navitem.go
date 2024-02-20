package navstack

import tea "github.com/charmbracelet/bubbletea"

type NavigationItem struct {
	Title string
	Model tea.Model
}

func (n NavigationItem) Init() tea.Cmd {
	return n.Model.Init()
}

func (n NavigationItem) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	nm, cmd := n.Model.Update(msg)
	n.Model = nm
	return n, cmd
}

func (n NavigationItem) View() string {
	return n.Model.View()
}
