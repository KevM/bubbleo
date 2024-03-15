package paintingcolors

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/kevm/bubbleo/examples/deeper/color"
	"github.com/kevm/bubbleo/examples/deeper/data"
	"github.com/kevm/bubbleo/menu"
	"github.com/kevm/bubbleo/navstack"
	"github.com/kevm/bubbleo/utils"
)

type Model struct {
	Painting data.Painting

	menu menu.Model
}

func New(painting data.Painting) Model {

	choices := []menu.Choice{}
	for _, c := range painting.Colors {
		choice := menu.Choice{
			Title:       lipgloss.NewStyle().Foreground(lipgloss.Color(c.RGB)).Render(c.RGB),
			Description: c.Sample,
			Model: color.Model{
				RGB:    c.RGB,
				Sample: c.Sample,
			},
		}
		choices = append(choices, choice)
	}

	title := fmt.Sprintf(" 🎨 Colors featured in %s", painting.Title)
	menu := menu.New(title, choices, nil)

	return Model{
		Painting: painting,
		menu:     menu,
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
		result := PaintingColorSelected{
			Painting: m.Painting.Title,
			Color:    msg.RGB,
		}
		cmd := utils.Cmdize(result)
		return m, tea.Sequence(pop, cmd)
	}

	updatedmenu, cmd := m.menu.Update(msg)
	m.menu = updatedmenu.(menu.Model)
	return m, cmd
}

func (m Model) View() string {
	return m.menu.View()
}
