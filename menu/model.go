// Package Menu takes a list of choices allowing the user to select a component
// to push onto the navigation stack. Each choice has a title and a description and
// a component model implementing [tea.Model].
// [tea.Model] https://github.com/charmbracelet/bubbletea/blob/a256e76ff5ff142d747ad833c7aa784113f8558c/tea.go#L39
package menu

import (
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/kevm/bubbleo/navstack"
	"github.com/kevm/bubbleo/styles"
	"github.com/kevm/bubbleo/utils"
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

// MenuStyles is a struct that holds the styles for the menu
// This mostly a passthrough for bubble/list component styles.
type MenuStyles struct {
	ListTitleStyle lipgloss.Style
	ListItemStyles list.DefaultItemStyles
}

type Model struct {
	Choices []Choice

	list     list.Model
	selected *Choice
	delegate list.DefaultDelegate

	width  int
	height int
	help.KeyMap
	keys KeyMap
	help help.Model
}

// New setups up a new menu model
func New(title string, choices []Choice, selected *Choice) Model {

	styles := MenuStyles{
		ListTitleStyle: styles.ListTitleStyle,
		ListItemStyles: list.NewDefaultItemStyles(),
	}

	delegation := list.NewDefaultDelegate()
	delegation.Styles = styles.ListItemStyles

	defaultWidth := 120
	defaultHeight := 20

	model := Model{
		list:     list.New([]list.Item{}, delegation, defaultWidth, defaultHeight),
		selected: selected,
		delegate: delegation,
		keys:     DefaultKeyMap,
		help:     help.New(),
		width:    defaultWidth,
		height:   defaultHeight,
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

	model.SetChoices(choices, selected)

	return model
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m *Model) SetChoices(choices []Choice, selected *Choice) {
	m.Choices = choices

	items := make([]list.Item, len(choices))
	selectedIndex := -1
	for i, choice := range choices {
		if selected != nil && &choice == selected {
			selectedIndex = i
		}
		items[i] = choiceItem{title: choice.Title, desc: choice.Description, key: choice}
	}

	m.list.SetItems(items)
	if selected != nil {
		m.selected = selected
		m.list.Select(selectedIndex)
	}
}

// SetStyles allows you to customize the styles used by the menu. This is mostly a passthrough
// to the bubble/list component used by the menu.
func (m Model) SetStyles(s MenuStyles) {
	m.list.Styles.Title = s.ListTitleStyle
	m.delegate.Styles = s.ListItemStyles
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.SetSize(msg)
	case tea.KeyMsg:
		return m.handleKeyMsg(msg, msg)
	}

	// No selection made yet so update the list
	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

// SelectChoice pushes the selected choice onto the navigation stack. If the choice is nil, nothing happens.
func (m Model) SelectChoice(choice *Choice) (Model, tea.Cmd) {
	if choice == nil {
		return m, nil
	}

	m.selected = choice
	item := navstack.NavigationItem{Title: choice.Title, Model: choice.Model}
	cmd := utils.Cmdize(navstack.PushNavigation{Item: item})

	return m, cmd
}

// SelectedChoice returns the currently selected menu choice
func (m Model) SelectedChoice() *Choice {
	return m.selected
}

// SetSize sets the size of the menu
func (m *Model) SetSize(w tea.WindowSizeMsg) {
	m.width = w.Width
	m.height = w.Height
	m.list.SetSize(w.Width, w.Height)
	m.help.Width = w.Width
}

// View renders the menu. When no choices are present, nothing is rendered.
func (m Model) View() string {
	var help string
	if m.help.ShowAll {
		height := m.height - 5
		m.list.SetSize(m.width, height)
		help = styles.HelpStyle.Render(m.help.View(m.keys))
	}

	// display menu if choices are present.
	if len(m.Choices) > 0 {
		return "\n" + m.list.View() + help
	}

	return ""
}
