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

	delegate list.DefaultDelegate
	list     list.Model
	width    int
	height   int
	help.KeyMap
	keys KeyMap
	help help.Model
}

// Option is an optional [Model] configuration function.
type Option func(*Model)

// WithShowPagination returns an [Option]
// whether the list shows pagination indictors.
func WithShowPagination(show bool) Option {
	return func(mod *Model) { mod.list.SetShowPagination(show) }
}

// WithShowTitle returns an [Option]
// whether to show the list title.
func WithShowTitle(show bool) Option {
	return func(mod *Model) { mod.list.SetShowTitle(show) }
}

// WithFilteringEnabled returns an [Option]
// whether the list allows filtering.
func WithFilteringEnabled(enabled bool) Option {
	return func(mod *Model) { mod.list.SetFilteringEnabled(enabled) }
}

// WithShowFilter returns an [Option]
// whether the list shows the active filter.
func WithShowFilter(show bool) Option {
	return func(mod *Model) { mod.list.SetShowFilter(show) }
}

// WithShowStatusBar returns an [Option]
// whether to show the list's status bar.
func WithShowStatusBar(show bool) Option {
	return func(mod *Model) { mod.list.SetShowStatusBar(show) }
}

// WithShowHelp returns an [Option]
// whether to always show help for active keybindings.
func WithShowHelp(show bool) Option {
	return func(mod *Model) { mod.list.SetShowHelp(show) }
}

// WithAdditionalFullHelpKeys returns an [Option]
// that replaces the lists's [list.Model.AdditionalFullHelpKeys].
//
// The list's current slice is passed as argument to the given keys function.
// The keys function can then choose to return a new or modified slice.
func WithAdditionalFullHelpKeys(keys func([]key.Binding) []key.Binding) Option {
	return func(mod *Model) {
		prev := mod.list.AdditionalFullHelpKeys()
		mod.list.AdditionalFullHelpKeys = func() []key.Binding {
			return keys(prev)
		}
	}
}

// WithAdditionalShortHelpKeys returns an [Option]
// that replaces the lists's [list.Model.AdditionalShortHelpKeys].
//
// The list's current slice is passed as argument to the given keys function.
// The keys function can then choose to return a new or modified slice.
func WithAdditionalShortHelpKeys(keys func([]key.Binding) []key.Binding) Option {
	return func(mod *Model) {
		prev := mod.list.AdditionalShortHelpKeys()
		mod.list.AdditionalShortHelpKeys = func() []key.Binding {
			return keys(prev)
		}
	}
}

// New setups up a new menu model
func New(title string, choices []Choice, selected *Choice, options ...Option) Model {

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

	for _, opt := range options {
		opt(&model)
	}

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
		// Make sure we do not trap keybindings required for editing filters.
		// While editing, all messages should be forwarded on to [list.Model].
		if !m.list.SettingFilter() {
			if mod, cmd := m.handleKeyMsg(msg, msg); cmd != nil {
				return mod, cmd
			}
		}
	}
	// No selection made yet so update the list
	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

// SelectChoice pushes the selected choice onto the navigation stack. If the choice is nil, nothing happens.
func (m Model) SelectChoice(choice Choice) (Model, tea.Cmd) {
	item := navstack.NavigationItem{Title: choice.Title, Model: choice.Model}
	cmd := utils.Cmdize(navstack.PushNavigation{Item: item})

	return m, cmd
}

// SetSize sets the size of the menu
func (m *Model) SetSize(w tea.WindowSizeMsg) {
	m.width = w.Width
	m.height = w.Height
	m.list.SetSize(w.Width, w.Height)
	m.help.Width = w.Width
}

func (m *Model) SetShowTitle(display bool) {
	m.list.SetShowTitle(display)
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
