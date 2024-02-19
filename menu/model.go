package menu

import (
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/kevm/bubbleo/navstack"
	"github.com/kevm/bubbleo/styles"
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
}

// New setups up a new menu model
func New(title string, choices []Choice, selected *Choice, width int, height int) Model {
	delegation := list.NewDefaultDelegate()
	items := make([]list.Item, len(choices))
	for i, choice := range choices {
		items[i] = choiceItem{title: choice.Title, desc: choice.Description, key: choice}
	}

	navstack := navstack.New()

	model := Model{
		Choices:  choices,
		list:     list.New(items, delegation, width, height),
		navstack: &navstack,
	}

	model.list.Styles.Title = styles.ListTitleStyle
	model.list.Title = title
	model.list.SetShowPagination(true)
	model.list.SetShowTitle(true)
	model.list.SetFilteringEnabled(false)
	model.list.SetShowFilter(false)
	model.list.SetShowStatusBar(false)
	model.list.SetShowHelp(false)

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

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	if m.navstack.Top() != nil {
		cmd := m.navstack.Update(msg)
		return m, cmd
	}

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case tea.KeyEnter.String():
			choice, ok := m.list.SelectedItem().(choiceItem)
			if ok {
				choice.key.Model.Init()
				item := navstack.NavigationItem{Title: choice.title, Model: choice.key.Model}
				cmd := m.navstack.Push(item)
				return m, cmd
			}
		}
	}

	// No selection made yet so update the list
	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m *Model) SetSize(w int, h int) {
	m.list.SetSize(w, h)
}

func (m Model) View() string {

	if m.navstack.Top() != nil {
		return m.navstack.View()
	}

	// display menu if choices are present.
	if len(m.Choices) > 0 {
		return "\n" + m.list.View()
	}

	return ""
}
