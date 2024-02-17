package menu

import (
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/kevm/bubbleo/styles"
	"github.com/kevm/bubbleo/window"
)

type Closable interface {
	Close() error
}

type Choice struct {
	Title         string
	Description   string
	ModelInitFunc func() (tea.Model, tea.Cmd)
}

type choiceItem struct {
	title, desc string
	key         Choice
}

func (i choiceItem) Title() string       { return i.title }
func (i choiceItem) Description() string { return i.desc }
func (i choiceItem) FilterValue() string { return i.title + i.desc }

type Model struct {
	Choices       []Choice
	Selected      *Choice
	SelectedModel *tea.Model
	list          list.Model
	window        *window.Model
}

// New setups up a new menu model
func New(title string, choices []Choice, selected *Choice, width int, height int) Model {
	delegation := list.NewDefaultDelegate()
	items := make([]list.Item, len(choices))
	for i, choice := range choices {
		items[i] = choiceItem{title: choice.Title, desc: choice.Description, key: choice}
	}

	model := Model{
		Choices:  choices,
		list:     list.New(items, delegation, width, height),
		Selected: selected,
		window:   &window.Model{Width: width, Height: height},
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
	switch msg := msg.(type) {
	case ReloadSelected:
		if m.Selected != nil {
			s := *m.Selected
			// TODO handle err
			m.Clear()
			return m.SelectChoiceCmd(s)
		}
	case DismissSelected:
		// TODO handle err
		m.Clear()
		return m, nil
	case tea.WindowSizeMsg:
		m.window.Width = msg.Width
		m.window.Height = msg.Height
		m.list.SetSize(m.window.Width, m.window.Height)
	default:
		if m.SelectedModel != nil {
			// selection made so route updates to the selected model
			sm := *m.SelectedModel
			switch msg := msg.(type) {
			default:
				um, cmd := sm.Update(msg)
				m.SelectedModel = &um
				return m, cmd
			}
		}

		switch msg := msg.(type) {
		case tea.KeyMsg:
			switch msg.Type {
			case tea.KeyEnter:
				choice, ok := m.list.SelectedItem().(choiceItem)
				if ok {
					return m.SelectChoiceCmd(choice.key)
				}
			}
		}

		// No selection made yet so update the list
		var cmd tea.Cmd
		m.list, cmd = m.list.Update(msg)
		return m, cmd
	}

	return m, nil // should never get here
}

func (m *Model) SetSize(w int, h int) {
	m.window.Width = w
	m.window.Height = h
	m.list.SetSize(w, h)
}

func (m Model) View() string {

	if m.SelectedModel != nil {
		// selection made so route view to the selected model
		sm := *m.SelectedModel
		return sm.View()
	}

	// display menu if choices are present.
	if len(m.Choices) > 0 {
		return styles.AppStyle.Render(m.list.View())
	}

	return ""
}

// SelectChoiceCmd selects the given choice.
// It returns the new model and potentially an initialize command to run.
func (m Model) SelectChoiceCmd(choice Choice) (tea.Model, tea.Cmd) {
	selectedModel, cmd := choice.ModelInitFunc()
	m.SelectedModel = &selectedModel
	m.Selected = &choice

	return m, cmd
}

// Clear the selected choice and close the model if it is closable
func (m *Model) Clear() error {

	var err error

	if m.SelectedModel != nil {
		sm := *m.SelectedModel
		if c, ok := sm.(Closable); ok {
			err = c.Close()
		}
	}
	m.Selected = nil
	m.SelectedModel = nil

	return err
}
