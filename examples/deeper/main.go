package main

import (
	"fmt"
	"log"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/kevm/bubbleo/examples/deeper/artistcolors"
	"github.com/kevm/bubbleo/examples/deeper/data"
	"github.com/kevm/bubbleo/menu"
	"github.com/kevm/bubbleo/navstack"
)

var docStyle = lipgloss.NewStyle()

type model struct {
	SelectedArtist string
	SelectedColor  string
	menu           menu.Model
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case artistcolors.ArtistSelected:
		m.SelectedArtist = msg.Name
		m.SelectedColor = msg.Color
		return m, tea.Quit
	case tea.KeyMsg:
		if msg.String() == "ctrl+c" {
			return m, tea.Quit
		}
	case tea.WindowSizeMsg:
		h, v := docStyle.GetFrameSize()
		m.menu.SetSize(msg.Width-h, msg.Height-v)
		return m, nil
	}

	updatedmenu, cmd := m.menu.Update(msg)
	m.menu = updatedmenu.(menu.Model)
	return m, cmd
}

func (m model) View() string {
	return docStyle.Render(m.menu.View())
}

func main() {
	ns := navstack.New()
	artists := data.GetArtists()
	choices := make([]menu.Choice, len(artists))
	for i, a := range artists {
		choices[i] = menu.Choice{
			Title:       a.Name,
			Description: a.Description,
			Model:       artistcolors.New(a, &ns),
		}
	}

	title := "Choose an Artist:"
	m := model{menu: menu.New(title, choices, nil, &ns)}

	p := tea.NewProgram(m, tea.WithAltScreen())

	result, err := p.Run()
	if err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}

	resultModel := result.(model)

	log.Printf("You selected the color %s from the artist %s ", resultModel.SelectedColor, resultModel.SelectedArtist)
}
