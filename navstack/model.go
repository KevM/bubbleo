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

	wmsg := m.window.GetWindowSizeMsg()
	nim, cmd := item.Model.Update(wmsg)
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

func (m Model) StackSummary() []string {
	summary := []string{}
	for _, item := range m.stack {
		summary = append(summary, item.Title)
	}

	return summary
}

func (m *Model) Update(msg tea.Msg) tea.Cmd {
	top := m.Top()
	switch msg := msg.(type) {
	case tea.WindowSizeMsg: // update the window size based on offsets
		if top == nil {
			return nil
		}
		m.window.Height = msg.Height
		m.window.Width = msg.Width
		msg.Width = m.window.Width - m.window.SideOffset
		msg.Height = m.window.Height - m.window.TopOffset
		um, cmd := top.Update(msg)
		m.stack[len(m.stack)-1] = um.(NavigationItem)
		return cmd
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
	default:
		if top == nil {
			return nil
		}
		um, cmd := top.Update(msg)
		m.stack[len(m.stack)-1] = um.(NavigationItem)
		return cmd
	}
}

func (m Model) View() string {

	top := m.Top()
	if top == nil {
		return ""
	}

	return top.View()
}
