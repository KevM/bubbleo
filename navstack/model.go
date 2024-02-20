package navstack

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/kevm/bubbleo/window"
)

type Closable interface {
	Close() error
}

type Model struct {
	stack  []NavigationItem
	window *window.Model
}

func New(w *window.Model) Model {
	model := Model{
		stack:  []NavigationItem{},
		window: w,
	}

	return model
}

func (m Model) Init() tea.Cmd {
	top := m.Top()
	if top == nil {
		return nil
	}

	return top.Init()
}

func (m *Model) Push(item NavigationItem) tea.Cmd {

	nim, cmd := item.Model.Update(m.window.GetWindowSizeMsg())
	item.Model = nim

	m.stack = append(m.stack, item)
	return tea.Batch(cmd, item.Init())
}

func (m *Model) Pop() tea.Cmd {
	top := m.Top()
	if top == nil {
		return tea.Quit // should not happen
	}

	if c, ok := top.Model.(Closable); ok {
		c.Close()
	}

	cmds := []tea.Cmd{}
	nim, cmd := top.Model.Update(m.window.GetWindowSizeMsg())
	top.Model = nim
	cmds = append(cmds, cmd, top.Init())

	m.stack = m.stack[:len(m.stack)-1]
	top = m.Top()
	if top == nil {
		return tea.Quit
	}

	return tea.Batch(cmds...)
}

func (m Model) Top() *NavigationItem {
	if len(m.stack) == 0 {
		return nil
	}

	top := m.stack[len(m.stack)-1]
	return &top
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	top := m.Top()
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.window.Height = msg.Height
		m.window.Width = msg.Width
	case ReloadCurrent:
		if top == nil {
			return m, nil
		}
		return m, top.Init()
	case PopNavigation:
		cmd := m.Pop()
		return m, cmd
	case PushNavigation:
		cmd := m.Push(msg.Item)
		return m, cmd
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
