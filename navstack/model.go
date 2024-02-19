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

func (m Model) Push(item NavigationItem) (tea.Model, tea.Cmd) {
	m.stack = append(m.stack, item)
	return m, item.Init()
}

func (m Model) Pop() (tea.Model, tea.Cmd) {
	top := m.Top()
	if top == nil {
		return m, nil
	}

	if c, ok := top.model.(Closable); ok {
		c.Close()
	}

	m.stack = m.stack[:len(m.stack)-1]
	top = m.Top()
	if top == nil {
		return m, nil
	}

	return m, top.Init()
}

func (m Model) Top() *NavigationItem {
	if len(m.stack) == 0 {
		return nil
	}

	top := m.stack[len(m.stack)-1]
	return &top
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	top := m.Top()
	switch msg.(type) {
	case Reload:
		if top == nil {
			return m, nil
		}
		return m, top.Init()
	case Dismiss:
		return m.Pop()
	}

	if top == nil {
		return m, nil
	}

	um, cmd := top.Update(msg)
	m.stack[len(m.stack)-1] = um.(NavigationItem)
	return m, cmd
}

func (m Model) View() string {

	top := m.Top()
	if top == nil {
		return ""
	}

	return top.View()
}
