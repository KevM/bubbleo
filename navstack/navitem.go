package navstack

import tea "github.com/charmbracelet/bubbletea"

// NavigationItem is a component that represents an item in the navigation stack.
// The top most item on the stack is rendered.
type NavigationItem struct {
	Title string
	Model tea.Model
}

// Init is called when the item is pushed onto the stack.
func (n NavigationItem) Init() tea.Cmd {
	return n.Model.Init()
}

// Update receives messages when the item is on top of the stack.
func (n NavigationItem) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	nm, cmd := n.Model.Update(msg)
	n.Model = nm
	return n, cmd
}

// View is calledn when the item is on top of the stack.
func (n NavigationItem) View() string {
	return n.Model.View()
}
