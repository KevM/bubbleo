package menu

import (
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/kevm/bubbleo/navstack"
	"github.com/kevm/bubbleo/styles"
	"github.com/kevm/bubbleo/window"
)

type Choice struct {
	Title       string
	Description string
	Model       tea.Model
}

type choiceItem struct {
	title, desc string
	key         Choice
}

func (i choiceItem) Title() string       { return i.title }
func (i choiceItem) Description() string { return i.desc }
func (i choiceItem) FilterValue() string { return i.title + i.desc }

type Model struct {
	Choices  []Choice
	list     list.Model
	navstack *navstack.Model

	selected *Choice
	window   *window.Model
}

// New setups up a new menu model. If no navstack is provided a new one will be created.
// navstack is optional and used to route update/views to the top of the navstack when a selection is made
func New(title string, choices []Choice, selected *Choice, window *window.Model, ns *navstack.Model) Model {
	delegation := list.NewDefaultDelegate()
	items := make([]list.Item, len(choices))
	for i, choice := range choices {
		items[i] = choiceItem{title: choice.Title, desc: choice.Description, key: choice}
	}

	model := Model{
		Choices:  choices,
		list:     list.New(items, delegation, 120, 20),
		navstack: ns,
		selected: selected,
		window:   window,
	}

	model.list.Styles.Title = styles.ListTitleStyle
	model.list.Title = title
	model.list.SetShowPagination(true)
	model.list.SetShowTitle(true)
	model.list.SetFilteringEnabled(false)
	model.list.SetShowFilter(false)
	model.list.SetShowStatusBar(false)
	model.list.SetShowHelp(false)

	//TODO: figure out height long term.
	model.list.SetSize(window.Width, window.Height-window.TopOffset)

	chooseKeyBinding := key.NewBinding(
		key.WithKeys("enter"),
		key.WithHelp("enter", "choose"),
	)
	model.list.AdditionalFullHelpKeys = func() []key.Binding {
		return []key.Binding{chooseKeyBinding}
	}
	model.list.AdditionalShortHelpKeys = func() []key.Binding {
		return []key.Binding{chooseKeyBinding}
	}

	return model
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) RouteToNavstack() bool {
	return m.selected != nil && m.navstack != nil && m.navstack.Top() != nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	if m.RouteToNavstack() {
		cmd := m.navstack.Update(msg)
		return m, cmd
	}

	switch msg := msg.(type) {
	case navstack.PopNavigation, navstack.PushNavigation, navstack.ReloadCurrent:
		if m.navstack != nil {
			cmd := m.navstack.Update(msg)
			return m, cmd
		}
	case tea.KeyMsg:
		switch msg.String() {
		case tea.KeyEsc.String():
			return m, cmdize(navstack.PopNavigation{})
		case tea.KeyEnter.String():
			choice, ok := m.list.SelectedItem().(choiceItem)
			if ok {
				m.selected = &choice.key
				item := navstack.NavigationItem{Title: choice.title, Model: choice.key.Model}
				cmd := cmdize(navstack.PushNavigation{Item: item})
				return m, cmd
			}
		}
	}

	// No selection made yet so update the list
	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m *Model) SetSize(w *window.Model) {
	m.list.SetSize(w.Width, w.Height-w.TopOffset)
}

func (m Model) View() string {

	if m.RouteToNavstack() {
		return m.navstack.View()
	}

	// display menu if choices are present.
	if len(m.Choices) > 0 {
		return "\n" + m.list.View()
	}

	return ""
}

func cmdize[T any](t T) tea.Cmd {
	return func() tea.Msg {
		return t
	}
}
