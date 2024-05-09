package color

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/kevm/bubbleo/navstack"
	"github.com/kevm/bubbleo/utils"
)

type SettingsUpdated struct {
	Color string
}

type Model struct {
	list list.Model
}

type ColorItem struct {
	title, desc string
}

func (i ColorItem) Title() string {
	return lipgloss.NewStyle().Foreground(lipgloss.Color(i.desc)).Render(i.title)
}
func (i ColorItem) Description() string { return i.desc }
func (i ColorItem) FilterValue() string { return i.title + i.desc }

func New() Model {
	items := []list.Item{
		ColorItem{title: "Red", desc: "#ff0000"},
		ColorItem{title: "Green", desc: "#00ff00"},
		ColorItem{title: "Blue", desc: "#0000ff"},
		ColorItem{title: "Yellow", desc: "#ffff00"},
		ColorItem{title: "Cyan", desc: "#00ffff"},
		ColorItem{title: "Magenta", desc: "#ff00ff"},
		ColorItem{title: "White", desc: "#ffffff"},
		ColorItem{title: "Black", desc: "#000000"},
	}

	return Model{
		list: list.New(items, list.NewDefaultDelegate(), 0, 0),
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.list.SetWidth(msg.Width)
		m.list.SetHeight(msg.Height)
	case tea.KeyMsg:
		switch msg.String() {
		case "q":
			return m, tea.Quit
		case "esc":
			pop := utils.Cmdize(navstack.PopNavigation{})
			return m, pop
		case "enter":
			selected := m.list.SelectedItem()
			item, ok := selected.(ColorItem)
			if !ok {
				return m, nil
			}
			pop := utils.Cmdize(navstack.PopNavigation{})
			cmd := utils.Cmdize(SettingsUpdated{Color: item.desc})
			return m, tea.Sequence(pop, cmd)
		}
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m Model) View() string {
	return m.list.View()
}
