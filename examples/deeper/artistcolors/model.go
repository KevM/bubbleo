package artistcolors

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/kevm/bubbleo/examples/deeper/color"
	"github.com/kevm/bubbleo/examples/deeper/data"
	"github.com/kevm/bubbleo/menu"
	"github.com/kevm/bubbleo/navstack"
	"github.com/kevm/bubbleo/window"
)

type Model struct {
	Artist data.Artist
	Nav    *navstack.Model

	menu   menu.Model
	window *window.Model
}

func New(a data.Artist, n *navstack.Model, window *window.Model) Model {

	choices := []menu.Choice{}
	for _, p := range a.Paintings {
		for _, c := range p.Colors {
			choice := menu.Choice{
				Title:       c.RGB,
				Description: c.Sample,
				Model: color.Model{
					RGB:    c.RGB,
					Sample: c.Sample,
				},
			}
			choices = append(choices, choice)
		}
	}

	menu := menu.New("Artist Colors", choices, nil, window, nil)
	menu.SetSize(window)

	return Model{
		Artist: a,
		Nav:    n,
		menu:   menu,
		window: window,
	}
}

func (m Model) Init() tea.Cmd {
	m.menu.SetSize(m.window)
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.menu.SetSize(m.window)
	case color.ColorSelected:
		cmd := cmdize(ArtistSelected{Name: m.Artist.Name, Color: msg.RGB})
		pop := cmdize(navstack.PopNavigation{})
		return m, tea.Batch(pop, cmd)
	}

	updatedmenu, cmd := m.menu.Update(msg)
	m.menu = updatedmenu.(menu.Model)
	return m, cmd
}

func (m Model) View() string {
	return m.menu.View()
}

func cmdize[T any](t T) tea.Cmd {
	return func() tea.Msg {
		return t
	}
}
