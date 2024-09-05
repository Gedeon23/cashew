package main

import (
	"log"
	"os/exec"
	"strings"

	"github.com/Gedeon23/cashew/details"
	"github.com/Gedeon23/cashew/entry"
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
	results list.Model
	keys    KeyMap
	help    help.Model
	focus   int
	details details.Model
	err     error
}

func newModel() model {
	var items []list.Item
	search := textinput.New()
	search.Placeholder = "search…"
	search.Prompt = " "
	search.Focus()
	search.CharLimit = 200
	search.Width = 20
	list := list.New(items, entry.NewEntryDelegate(), 0, 0)
	list.Title = "Results"
	list.SetFilteringEnabled(false)
	list.SetShowHelp(false)
	keys := NewDefaultKeyMap()
	help := help.New()
	details := details.New()
	return model{
		search:  search,
		results: list,
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
	selected := m.results.SelectedItem()
	m.details.Query = m.search.Value()
	switch selected := selected.(type) {
	case entry.Recoll:
		m.details.Entry = selected
	default:
		m.details.Entry = entry.Recoll{
			DocTitle: "",
			Author:   "",
			File:     "",
			Url:      "",
		}
	}
}

func (m *model) ExpandHelp() {
	prevHeight := strings.Count(m.help.View(m.keys), "\n")
	m.help.ShowAll = !m.help.ShowAll
	newHeight := strings.Count(m.help.View(m.keys), "\n")
	m.results.SetSize(
		m.results.Width(),
		m.results.Height()-(newHeight-prevHeight),
	)
}

func (m *model) OpenSelected() {
	selected := m.results.SelectedItem()
	if selected, ok := selected.(entry.Recoll); ok {
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
				// REFACTOR into Cmd probably?-------------------+
				if !(m.search.Value() == "") {
					m.results.SetItems(Collect(m.search.Value()))
				}
				//----------------------------------------------------+
				m.SetFocus(FocusResults)
				m.UpdateDetails()
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
				m.results, cmd = m.results.Update(msg)
				m.UpdateDetails()
				cmds = append(cmds, cmd)
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
			case key.Matches(msg, m.keys.OpenDocument):
				m.OpenSelected()
			default:
				m.details, cmd = m.details.Update(msg)
				cmds = append(cmds, cmd)
			}
		}
	case tea.WindowSizeMsg:
		h, v := styles.Root.GetFrameSize()
		m.results.SetSize(
			msg.Width/2-h,
			msg.Height-v-5,
		)
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
					m.results.View(),
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
