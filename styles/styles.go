package styles

import (
	"github.com/charmbracelet/lipgloss"
)

var (
	foregroundColor = lipgloss.AdaptiveColor{Light: "#343433", Dark: "#C1C6B2"}
	backgroundColor = lipgloss.AdaptiveColor{Light: "#D9DCCF", Dark: "#353533"}

	AppStyle       = lipgloss.NewStyle().Padding(2, 5)
	ListStyle      = lipgloss.NewStyle().Margin(1, 2)
	ListItemStyle  = lipgloss.NewStyle().PaddingLeft(4)
	ListTitleStyle = lipgloss.NewStyle().MarginLeft(2).Background(backgroundColor).Foreground(foregroundColor).Bold(true)

	BreadCrumbFrameStyle = lipgloss.NewStyle().Background(backgroundColor).Foreground(foregroundColor).Padding(1)
)
