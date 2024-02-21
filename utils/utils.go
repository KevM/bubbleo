package utils

import tea "github.com/charmbracelet/bubbletea"

func Cmdize[T any](t T) tea.Cmd {
	return func() tea.Msg {
		return t
	}
}
