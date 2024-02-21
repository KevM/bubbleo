package artistcolors

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/kevm/bubbleo/examples/deeper/color"
	"github.com/kevm/bubbleo/examples/deeper/data"
	"github.com/kevm/bubbleo/menu"
	"github.com/kevm/bubbleo/navstack"
	"github.com/kevm/bubbleo/utils"
)

type Model struct {
	Artist data.Artist

	menu menu.Model
}

func New(a data.Artist) Model {

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

	menu := menu.New("Artist Colors", choices, nil)

	return Model{
		Artist: a,
		menu:   menu,
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.menu.SetSize(msg)
	case color.ColorSelected:
		pop := utils.Cmdize(navstack.PopNavigation{})
		cmd := utils.Cmdize(ArtistSelected{Name: m.Artist.Name, Color: msg.RGB})
		return m, tea.Sequence(pop, cmd)
	}

	updatedmenu, cmd := m.menu.Update(msg)
	m.menu = updatedmenu.(menu.Model)
	return m, cmd
}

func (m Model) View() string {
	return m.menu.View()
}
