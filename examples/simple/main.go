package main

import (
	"fmt"
	"log"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/kevm/bubbleo/examples/simple/color"
	"github.com/kevm/bubbleo/menu"
	"github.com/kevm/bubbleo/navstack"
	"github.com/kevm/bubbleo/shell"
)

var docStyle = lipgloss.NewStyle()

type model struct {
	SelectedColor string
	menu          menu.Model
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case color.ColorSelected:
		m.SelectedColor = msg.RGB
		return m, tea.Quit
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
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
	return docStyle.Render(m.menu.View())
}

func main() {
	red := menu.Choice{
		Title:       "Red Envy",
		Description: "Raindrops on roses",
		Model:       color.Model{RGB: "#FF0000", Sample: "❤️ Love makes the world go around ❤️"},
	}

	green := menu.Choice{
		Title:       "Green Grass",
		Description: "Green grows the grass over thy neighbors septic tank",
		Model:       color.Model{RGB: "#00FF00", Sample: "☘️ The luck you make for yourself ☘️"},
	}

	blue := menu.Choice{
		Title:       "Blue Shoes",
		Description: "But did he cry?! No!",
		Model:       color.Model{RGB: "#0000FF", Sample: "🧿 Never forget what it's like to feel young 🧿"},
	}

	choices := []menu.Choice{red, green, blue}

	title := "Colorful Choices"
	// top, side := docStyle.GetFrameSize()
	// w := window.New(120, 25, top, side)
	// ns := navstack.New(&w)
	m := model{
		menu: menu.New(title, choices, nil),
	}
	s := shell.New()
	s.Navstack.Push(navstack.NavigationItem{Model: m, Title: "Colors"})
	p := tea.NewProgram(s, tea.WithAltScreen())

	finalshell, err := p.Run()
	if err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}

	topNavItem := finalshell.(shell.Model).Navstack.Top()
	if topNavItem == nil {
		log.Printf("Nothing selected")
		os.Exit(1)
	}

	selected := topNavItem.Model.(model)
	log.Printf("You selected the color: %s", selected.SelectedColor)
}
