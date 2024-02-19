package artistcolors

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/kevm/bubbleo/examples/deeper/color"
	"github.com/kevm/bubbleo/examples/deeper/data"
	"github.com/kevm/bubbleo/menu"
	"github.com/kevm/bubbleo/navstack"
)

type Model struct {
	Artist data.Artist
	Nav    *navstack.Model

	menu menu.Model
}

func New(a data.Artist, n *navstack.Model) Model {

	choices := make([]menu.Choice, len(a.Paintings))
	for i, p := range a.Paintings {
		for _, c := range p.Colors {
			choices[i] = menu.Choice{
				Title:       c.RGB,
				Description: c.Sample,
				Model: color.Model{
					RGB:    c.RGB,
					Sample: c.Sample,
				},
			}
		}
	}

	return Model{
		Artist: a,
		Nav:    n,
		menu:   menu.New("Artist Colors", choices, nil, nil),
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	switch msg := msg.(type) {
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
	// sample := lipgloss.NewStyle().
	// 	Foreground(lipgloss.Color(m.Artist.Name)).
	// 	Render(m.Artist.Description)

	// return sample + "\n\n\n\n\n" + "enter: select, esc: back\n"
	return m.menu.View()
}

func cmdize[T any](t T) tea.Cmd {
	return func() tea.Msg {
		return t
	}
}
