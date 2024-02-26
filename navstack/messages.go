package navstack

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/kevm/bubbleo/utils"
)

// ReloadCurrent is a message that can be sent to the menu model to reload the currently selected menu choice
type ReloadCurrent struct{}

// PopNavigation is a message that can be sent to the menu model to de-select the currently selected menu choice
type PopNavigation struct{}

type PushNavigation struct {
	Item NavigationItem
}

func PopNavigationCmd() tea.Cmd {
	return utils.Cmdize(PopNavigation{})
}

func PushNavigationCmd(item NavigationItem) tea.Cmd {
	return utils.Cmdize(PushNavigation{Item: item})
}
