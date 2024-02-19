package main

import (
	"fmt"
	"log"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/kevm/bubbleo/examples/simple/color"
	"github.com/kevm/bubbleo/menu"
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
	red := menu.Choice{
		Title:       "Red Envy",
		Description: "Raindrops on roses",
		ModelInitFunc: func() (tea.Model, tea.Cmd) {
			return color.Model{RGB: "#FF0000", Sample: "❤️ Love makes the world go around ❤️"}, nil
		},
	}
	green := menu.Choice{
		Title:       "Green Grass",
		Description: "Green grows the grass over thy neighbors septic tank",
		ModelInitFunc: func() (tea.Model, tea.Cmd) {
			return color.Model{RGB: "#00FF00", Sample: "☘️ The luck you make for yourself ☘️"}, nil
		},
	}
	blue := menu.Choice{
		Title:       "Blue Shoes",
		Description: "But did he cry?! No!",
		ModelInitFunc: func() (tea.Model, tea.Cmd) {
			return color.Model{RGB: "#0000FF", Sample: "🧿 Never forget what it's like to feel young 🧿"}, nil
		},
	}

	choices := []menu.Choice{red, green, blue}

	title := "Colorful Choices"
	m := model{menu: menu.New(title, choices, nil, 80, 20)}

	p := tea.NewProgram(m, tea.WithAltScreen())

	result, err := p.Run()
	if err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}

	log.Printf("You selected the color: %s", result.(model).SelectedColor)
}
