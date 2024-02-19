package navstack

import tea "github.com/charmbracelet/bubbletea"

type NavigationItem struct {
	Title     string
	model     tea.Model
	modelInit *func() (tea.Model, tea.Cmd)
}

func NewNavFromModel(title string, model tea.Model) NavigationItem {
	return NavigationItem{
		Title: title,
		model: model,
	}
}

func NewNavFromFunc(title string, modelInit func() (tea.Model, tea.Cmd)) NavigationItem {
	return NavigationItem{
		Title:     title,
		modelInit: &modelInit,
	}
}

func (n NavigationItem) Init() tea.Cmd {
	if n.model != nil {
		return n.model.Init()
	}

	if n.modelInit != nil {
		initFunc := *n.modelInit
		model, cmd := initFunc()
		n.model = model
		return cmd
	}

	return nil
}

func (n NavigationItem) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	nm, cmd := n.model.Update(msg)
	n.model = nm
	return n, cmd
}

func (n NavigationItem) View() string {
	return n.model.View()
}
