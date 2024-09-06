package main

import (
	"log"
	"os/exec"
	"strings"

	"github.com/Gedeon23/cashew/recoll"
	"github.com/Gedeon23/cashew/styles"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const (
	FocusSearch = iota
	FocusResults
	FocusDetail
)

type model struct {
	search  textinput.Model
	list    list.Model
	keys    KeyMap
	help    help.Model
	focus   int
	details Details
	err     error
}

func newModel() model {
	search := textinput.New()
	search.Placeholder = "search…"
	search.Prompt = " "
	search.Focus()
	search.CharLimit = 200
	search.Width = 20

	var results []list.Item
	list := list.New(results, NewEntryDelegate(), 0, 0)
	list.SetFilteringEnabled(false)
	list.SetShowTitle(false)
	list.SetShowHelp(false)

	keys := NewDefaultKeyMap()
	help := help.New()

	details := NewDetails()
	return model{
		search:  search,
		list:    list,
		keys:    keys,
		help:    help,
		details: details,
	}
}

func (m *model) SetFocus(Focus int) {
	if m.focus == FocusSearch && Focus != FocusSearch {
		m.search.Blur()
	}
	m.focus = Focus
	m.keys.Focus = Focus
	if Focus == FocusSearch {
		m.search.Focus()
	}
}

func (m *model) NextFocus() {
	if m.focus == FocusSearch {
		m.search.Blur()
	}
	m.focus = (m.focus + 1) % 3
	m.keys.Focus = m.focus
	if m.focus == FocusSearch {
		m.search.Focus()
	}
}

func (m *model) UpdateDetails() {
	selected := m.list.SelectedItem()
	switch selected := selected.(type) {
	case recoll.Entry:
		m.details.Update(SwitchEntryMsg{NewEntry: &selected})
	}
}

func (m *model) ExpandHelp() {
	prevHeight := strings.Count(m.help.View(m.keys), "\n")
	m.help.ShowAll = !m.help.ShowAll
	newHeight := strings.Count(m.help.View(m.keys), "\n")
	m.list.SetSize(
		m.list.Width(),
		m.list.Height()-(newHeight-prevHeight),
	)
}

func (m *model) OpenSelected() {
	selected := m.list.SelectedItem()
	if selected, ok := selected.(recoll.Entry); ok {
		cmd := exec.Command("xdg-open", selected.Url)

		if err := cmd.Start(); err != nil {
			log.Printf("Error: %v\n", err)
		}
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch m.focus {
		case FocusSearch:
			switch {
			case key.Matches(msg, m.keys.FocusNext):
				m.NextFocus()
			case key.Matches(msg, m.keys.ExecuteSearch):
				if !(m.search.Value() == "") {
					cmd = Collect(m.search.Value())
					cmds = append(cmds, cmd)
				}
			case key.Matches(msg, m.keys.Quit_ESC):
				return m, tea.Quit
			case key.Matches(msg, m.keys.Help_QM):
				m.ExpandHelp()
			default:
				m.search, cmd = m.search.Update(msg)
				cmds = append(cmds, cmd)
			}
		case FocusResults:
			switch {
			case key.Matches(msg, m.keys.FocusNext):
				m.NextFocus()
			case key.Matches(msg, m.keys.Quit):
				return m, tea.Quit
			case key.Matches(msg, m.keys.FocusSearch):
				m.SetFocus(FocusSearch)
			case key.Matches(msg, m.keys.FocusSearchAndClear):
				m.SetFocus(FocusSearch)
				m.search.SetValue("")
			case key.Matches(msg, m.keys.Help):
				m.ExpandHelp()
			case key.Matches(msg, m.keys.OpenDocument):
				m.OpenSelected()
			default:
				m.list, cmd = m.list.Update(msg)
				cmds = append(cmds, cmd)
				if entry, ok := m.list.SelectedItem().(recoll.Entry); ok {
					m.details, cmd = m.details.Update(SwitchEntryMsg{NewEntry: &entry})
					cmds = append(cmds, cmd)
				}
			}
		case FocusDetail:
			switch {
			case key.Matches(msg, m.keys.FocusNext):
				m.NextFocus()
			case key.Matches(msg, m.keys.Quit):
				return m, tea.Quit
			case key.Matches(msg, m.keys.FocusSearch):
				m.SetFocus(FocusSearch)
			case key.Matches(msg, m.keys.FocusSearchAndClear):
				m.SetFocus(FocusSearch)
				m.search.SetValue("")
			case key.Matches(msg, m.keys.Help_QM):
				m.ExpandHelp()
			// case key.Matches(msg, m.keys.OpenDocument):
			// 	m.OpenSelected()
			default:
				m.details, cmd = m.details.Update(msg)
				cmds = append(cmds, cmd)
			}
		}
	case tea.WindowSizeMsg:
		h, v := styles.Root.GetFrameSize()
		m.list.SetSize(
			msg.Width/2-h,
			msg.Height-v-5,
		)
	case CollectMsg:
		m.list.SetItems(msg.Results)
		m.SetFocus(FocusResults)
		if entry, ok := m.list.SelectedItem().(recoll.Entry); ok {
			m.details, cmd = m.details.Update(SwitchEntryMsg{NewEntry: &entry})
			cmds = append(cmds, cmd)
		}
	case SnippetsMsg:
		m.details, cmd = m.details.Update(msg)
		cmds = append(cmds, cmd)
	}

	return m, tea.Batch(cmds...)
}

func (m model) View() string {
	if m.err != nil {
		return "\n\nError: " + m.err.Error()
	}

	return styles.Root.Render(
		lipgloss.JoinVertical(0,
			lipgloss.JoinHorizontal(0,
				lipgloss.JoinVertical(0,
					m.search.View(),
					"\n",
					m.list.View(),
				),
				m.details.View(),
			),
			m.help.View(m.keys),
		))
}

func main() {
	f, err := tea.LogToFile("/tmp/cashew_debug.log", "debug")
	if err != nil {
		log.Fatalf("err: %v", err)
	}
	defer f.Close()
	p := tea.NewProgram(newModel(), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}
