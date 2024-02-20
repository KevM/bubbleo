package navstack

import (
	tea "github.com/charmbracelet/bubbletea"
)

type Closable interface {
	Close() error
}

type Model struct {
	stack []NavigationItem
}

func New() Model {
	model := Model{
		stack: []NavigationItem{},
	}

	return model
}

func (m *Model) Push(item NavigationItem) tea.Cmd {
	m.stack = append(m.stack, item)
	return item.Init()
}

func (m *Model) Pop() tea.Cmd {
	top := m.Top()
	if top == nil {
		return nil
	}

	if c, ok := top.Model.(Closable); ok {
		c.Close()
	}

	m.stack = m.stack[:len(m.stack)-1]
	top = m.Top()
	if top == nil {
		return nil
	}

	return top.Init()
}

func (m Model) Top() *NavigationItem {
	if len(m.stack) == 0 {
		return nil
	}

	top := m.stack[len(m.stack)-1]
	return &top
}

func (m *Model) Update(msg tea.Msg) tea.Cmd {
	top := m.Top()
	switch msg := msg.(type) {
	case ReloadCurrent:
		if top == nil {
			return nil
		}
		return top.Init()
	case PopNavigation:
		cmd := m.Pop()
		return cmd
	case PushNavigation:
		cmd := m.Push(msg.Item)
		return cmd
	}

	if top == nil {
		return nil
	}

	cmd := top.Update(msg)
	// m.stack[len(m.stack)-1] = um.(NavigationItem)
	return cmd
}

func (m Model) View() string {

	top := m.Top()
	if top == nil {
		return ""
	}

	return top.View()
}
