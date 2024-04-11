package artistpaintings

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/kevm/bubbleo/examples/deeper/data"
	"github.com/kevm/bubbleo/examples/deeper/paintingcolors"
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
		choice := menu.Choice{
			Title:       p.Title,
			Description: p.Description,
			Model:       paintingcolors.New(p),
		}
		choices = append(choices, choice)
	}

	title := fmt.Sprintf(" 🖼️  Paintings by %s", a.Name)
	menu := menu.New(title, choices, nil)

	return Model{
		Artist: a,
		menu:   menu,
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Close() error {
	// This is called when the model is pushed and popped from the navigation stack.
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.menu.SetSize(msg)
	case paintingcolors.PaintingColorSelected:
		pop := utils.Cmdize(navstack.PopNavigation{})
		cmd := utils.Cmdize(ArtistPaintingColorSelected{Name: m.Artist.Name, Painting: msg.Painting, Color: msg.Color})
		return m, tea.Sequence(pop, cmd)
	}

	updatedmenu, cmd := m.menu.Update(msg)
	m.menu = updatedmenu.(menu.Model)
	return m, cmd
}

func (m Model) View() string {
	return m.menu.View()
}
