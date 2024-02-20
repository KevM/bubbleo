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
	"github.com/kevm/bubbleo/window"
)

var docStyle = lipgloss.NewStyle()

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
	ns := navstack.New()
	w := window.New(10, 20, 6)
	m := model{
		menu:   menu.New(title, choices, nil, &w),
		window: &w,
	}
	ns.Push(navstack.NavigationItem{Model: m, Title: "main menu"})

	p := tea.NewProgram(ns, tea.WithAltScreen())

	finalns, err := p.Run()
	if err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}

	result := finalns.(navstack.Model).Top().Model.(model)
	log.Printf("You selected the color: %s", result.SelectedColor)
}
