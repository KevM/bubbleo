// Package styles provides default styles for the bubbleo components
package styles

import (
	"github.com/charmbracelet/lipgloss"
)

var (
	foregroundColor = lipgloss.AdaptiveColor{Light: "#343433", Dark: "#C1C6B2"}
	backgroundColor = lipgloss.AdaptiveColor{Light: "#D9DCCF", Dark: "#353533"}

	ListTitleStyle       = lipgloss.NewStyle().MarginLeft(2).Background(backgroundColor).Foreground(foregroundColor).Bold(true)
	BreadCrumbFrameStyle = lipgloss.NewStyle().Background(backgroundColor).Foreground(foregroundColor).Margin(1)
)
