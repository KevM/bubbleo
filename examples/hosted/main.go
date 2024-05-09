package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	settings "github.com/kevm/bubbleo/examples/hosted/settings"
	"github.com/kevm/bubbleo/navstack"
	"github.com/kevm/bubbleo/utils"
	"github.com/kevm/bubbleo/window"
)

type ShowSettings struct{}

type model struct {
	SelectedColor string
	navstack      navstack.Model
}

func (m model) Init() tea.Cmd {
	if m.SelectedColor == "" {
		return utils.Cmdize(ShowSettings{})
	}
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if m.navstack.Top() != nil {
		return m, m.navstack.Update(msg)
	}

	switch msg := msg.(type) {
	case ShowSettings:
		item := navstack.NavigationItem{Model: settings.New(), Title: "Settings"}
		cmd := m.navstack.Push(item)
		return m, cmd
	case settings.SettingsUpdated:
		m.SelectedColor = msg.Color
		return m, nil
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "esc", "q":
			return m, tea.Quit
		case "s":
			return m, utils.Cmdize(ShowSettings{})
		}
	}

	return m, nil
}

func (m model) View() string {
	if m.navstack.Top() != nil {
		return m.navstack.View()
	}

	return "Color selected: " + lipgloss.NewStyle().Foreground(lipgloss.Color(m.SelectedColor)).Render(m.SelectedColor)
}

func main() {
	window := window.New(100, 50, 0, 0)
	ns := navstack.New(&window).QuitOnEmptyStack(false) // hosted users likely want to avoid quitting when the stack is empty
	m := model{
		navstack: ns,
	}
	p := tea.NewProgram(m, tea.WithAltScreen())

	_, err := p.Run()
	if err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
