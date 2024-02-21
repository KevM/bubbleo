package main

import (
	"fmt"
	"log"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/kevm/bubbleo/breadcrumb"
	"github.com/kevm/bubbleo/examples/deeper/artistcolors"
	"github.com/kevm/bubbleo/examples/deeper/data"
	"github.com/kevm/bubbleo/menu"
	"github.com/kevm/bubbleo/navstack"
	"github.com/kevm/bubbleo/window"
)

var docStyle = lipgloss.NewStyle()

type model struct {
	SelectedArtist string
	SelectedColor  string

	menu       menu.Model
	breadcrumb breadcrumb.Model
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
		m.menu.SetSize(msg)
		return m, nil
	}

	updatedmenu, cmd := m.menu.Update(msg)
	m.menu = updatedmenu.(menu.Model)
	return m, cmd
}

func (m model) View() string {
	bc := m.breadcrumb.View()
	menu := m.menu.View()
	return docStyle.Render(bc, menu)
}

func main() {
	artists := data.GetArtists()
	choices := make([]menu.Choice, len(artists))
	for i, a := range artists {
		choices[i] = menu.Choice{
			Title:       a.Name,
			Description: a.Description,
			Model:       artistcolors.New(a),
		}
	}

	maintop, mainside := docStyle.GetFrameSize()
	w := window.New(120, 25, maintop, mainside)
	ns := navstack.New(&w)

	// Add the breadcrumb height to the top offset
	bc := breadcrumb.New(&ns)
	bctop, bcside := bc.FrameStyle.GetFrameSize()
	w.TopOffset += bctop
	w.SideOffset += bcside

	title := "Choose an Artist:"
	m := model{
		menu:       menu.New(title, choices, nil),
		breadcrumb: bc,
	}

	ns.Push(navstack.NavigationItem{Model: m, Title: "main menu"})
	p := tea.NewProgram(ns, tea.WithAltScreen())

	finalns, err := p.Run()
	if err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}

	// the resulting model is a navstack. With the top model being the one that quit.
	topNavItem := finalns.(navstack.Model).Top()
	if topNavItem == nil {
		log.Printf("Nothing selected")
		os.Exit(1)
	}

	selected := topNavItem.Model.(model)
	log.Printf("You selected the color %s from the artist %s ", selected.SelectedColor, selected.SelectedArtist)
}
