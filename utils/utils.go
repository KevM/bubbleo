// Package Utils is for bubblo utility functions
package utils

import tea "github.com/charmbracelet/bubbletea"

// Cmdize is a utility function to convert a given value into a `tea.Cmd`
func Cmdize[T any](t T) tea.Cmd {
	return func() tea.Msg {
		return t
	}
}
