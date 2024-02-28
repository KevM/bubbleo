// Package Navstack manages a stack of NavigationItems which can be pushed or popped from the stack.
// The top most stack navigation item is used by [BubbleTea] to Update and renders it's View.
// When pushing and popping items from the stack, the new view to be presented is sent a tea.WindowSizeMsg
// to ensure it's view can be presented correctly. When the last item is popped from the stack the application will quit.
// NavigationItem models which implement the Closable interface will have their Close method called when they are popped from the stack.
// This is useful for cleaning up resources that may not be garbage collected when a view a no longer needed.
// [BubbleTea]: https://github.com/charmbracelet/bubbletea
package navstack

import (
	"errors"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/kevm/bubbleo/window"
)

// Closable is an interface for models that have resources that need to be cleaned up when
// they are no longer needed. The navigation stack checks for this interface when popping items.
type Closable interface {
	Close() error
}

type Model struct {
	stack  []NavigationItem
	window *window.Model
}

// New creates a new navigation stack model. The window is used to
// constrain the view within the container of the navigation stack.
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

// Push pushes a new navigation item onto the stack.
// The new navigation item is given a tea.WindowSizeMsg to ensure it's view can be presented correctly.
// The item's Init method is called and resulting command is processed by [BubbleTea].
// This new item will be the top most item on the stack and thus will be rendered.
func (m *Model) Push(item NavigationItem) tea.Cmd {

	wmsg := m.window.GetWindowSizeMsg()
	nim, cmd := item.Model.Update(wmsg)
	item.Model = nim

	m.stack = append(m.stack, item)
	return tea.Batch(cmd, item.Init())
}

// Pop removes the top most navigation item from the stack.
// If the item implements the Closable interface the Close method is called.
// The new top most item on the stack is given a tea.WindowSizeMsg to ensure it's view can be presented correctly.
// If there are no more items on the stack the application will quit.
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

// Clear pops all the items from the stack.
func (m *Model) Clear() error {
	var errs []error
	for _, item := range m.stack {
		if c, ok := item.Model.(Closable); ok {
			err := c.Close()
			if err != nil {
				errs = append(errs, err)
			}
		}
	}

	m.stack = []NavigationItem{}
	return errors.Join(errs...)
}

// Top returns the top most navigation item on the stack.
func (m Model) Top() *NavigationItem {
	if len(m.stack) == 0 {
		return nil
	}

	top := m.stack[len(m.stack)-1]
	return &top
}

// StackSummary returns a list of titles for each item on the stack.
// This is currently used by the breadcrumb component to render the breadcrumb trail.
func (m Model) StackSummary() []string {
	summary := []string{}
	for _, item := range m.stack {
		summary = append(summary, item.Title)
	}

	return summary
}

// Update processes messages for the top most navigation item on the stack.
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

// View renders the top most navigation item on the stack.
func (m Model) View() string {

	top := m.Top()
	if top == nil {
		return ""
	}

	return top.View()
}
