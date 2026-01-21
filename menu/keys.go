package menu

import (
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/kevm/bubbleo/navstack"
)

type KeyMap struct {
	Select key.Binding
	Back   key.Binding
	Help   key.Binding
	Quit   key.Binding
}

var DefaultKeyMap = KeyMap{
	Select: key.NewBinding(
		key.WithKeys("enter"),
		key.WithHelp("enter", "Select current choice"),
	),
	Back: key.NewBinding(
		key.WithKeys("esc"),
		key.WithHelp("esc", "Back to previous view"),
	),
	Help: key.NewBinding(
		key.WithKeys("?", "h"),
		key.WithHelp("? / h", "toggle help"),
	),
	Quit: key.NewBinding(
		key.WithKeys("q", "ctrl+c"),
		key.WithHelp("q / ctrl+c", "quit"),
	),
}

// ShortHelp returns keybindings to be shown in the mini help view. It's part
// of the key.Map interface.
func (k KeyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.Help, k.Quit}
}

// FullHelp returns keybindings for the expanded help view. It's part of the
// key.Map interface.
func (k KeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{}, {
			k.Help, k.Quit, k.Back, k.Select,
		},
	}
}

func (m Model) handleKeyMsg(keyMsg tea.KeyMsg, msg tea.Msg) (tea.Model, tea.Cmd) {
	if m.autoHideHelp && m.help.ShowAll && !key.Matches(keyMsg, DefaultKeyMap.Help) {
		m.help.ShowAll = false // toggle help view
		switch {               //override escape to only close help
		case keyMsg.String() == tea.KeyEscape.String():
			return m, nil
		}
	}

	switch {
	case key.Matches(keyMsg, DefaultKeyMap.Help):
		m.help.ShowAll = !m.help.ShowAll
	case key.Matches(keyMsg, DefaultKeyMap.Quit):
		return m, tea.Quit
	case key.Matches(keyMsg, DefaultKeyMap.Back):
		return m, navstack.PopNavigationCmd()
	case key.Matches(keyMsg, DefaultKeyMap.Select):
		choice, ok := m.list.SelectedItem().(choiceItem)
		if ok {
			return m.SelectChoice(choice.key)
		}
	default:
		l, cmd := m.list.Update(msg)
		m.list = l
		return m, cmd
	}

	return m, nil
}
