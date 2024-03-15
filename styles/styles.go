// Package styles provides default styles for the bubbleo components
package styles

import (
	"github.com/charmbracelet/lipgloss"
)

var (
	ListTitleStyle       = lipgloss.NewStyle().MarginLeft(2).Foreground(lipgloss.Color("230")).Bold(true)
	BreadCrumbFrameStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("229")).Margin(1)
	HelpStyle            = lipgloss.NewStyle().Padding(1, 2)
)
