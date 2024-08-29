package main

import (
	"log"
	"strings"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var style = lipgloss.NewStyle().Margin(1, 2)

type RecollEntry struct {
	Author   string
	DocTitle string
	File     string
	Url      string
}

func (e RecollEntry) Title() string {
	var icon string = " "
	if e.File[len(e.File)-3:] == "pdf" {
		icon = " "
	}
	return icon + e.DocTitle
}
func (e RecollEntry) Description() string { return " " + e.Author }
func (e RecollEntry) FilterValue() string { return "" + e.Url }

const (
	FocusSearch = iota
	FocusResults
	FocusDetail
)

type model struct {
	search  textinput.Model
	results list.Model
	keys    AdaptiveKeyMap
	help    help.Model
	focus   int
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
	list := list.New(items, NewEntryDelegate(), 0, 0)
	list.Title = "Results"
	list.SetFilteringEnabled(false)
	list.SetShowHelp(false)
	keys := NewAdaptiveKeyMap()
	help := help.New()
	return model{
		search:  search,
		results: list,
		keys:    keys,
		help:    help,
	}
}

func (m model) SetFocus(focus int) {
	m.focus = focus
	m.keys.Focus = focus
}

func (m model) NextFocus() int {
	m.focus = (m.focus + 1) % 3
	m.keys.Focus = m.focus
	return m.focus
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keys.Global.FocusSearch):
			m.SetFocus(FocusSearch)
		case key.Matches(msg, m.keys.Global.FocusNext):
			m.NextFocus()
		case key.Matches(msg, m.keys.Search.ExecuteSearch):
			// TODO Refactor unto Cmd probably?-------------------+
			if !(m.search.Value() == "") {
				m.results.SetItems(Collect(m.search.Value()))
			}
			//----------------------------------------------------+
			m.search.Blur()
			m.SetFocus(FocusResults)
		case key.Matches(msg, m.keys.Global.Quit):
			return m, tea.Quit
		case key.Matches(msg, m.keys.Global.Help):
			prevHeight := strings.Count(m.help.View(m.keys), "\n")
			m.help.ShowAll = !m.help.ShowAll
			newHeight := strings.Count(m.help.View(m.keys), "\n")
			m.results.SetSize(
				m.results.Width(),
				m.results.Height()-(newHeight-prevHeight),
			)
		default:
			var cmd tea.Cmd
			switch m.focus {
			case FocusSearch:
				m.search, cmd = m.search.Update(msg)
			case FocusResults:
				m.results, cmd = m.results.Update(msg)
			}
			cmds = append(cmds, cmd)
		}

	case tea.WindowSizeMsg:
		h, v := style.GetFrameSize()
		m.results.SetSize(
			msg.Width-h,
			msg.Height-v-5,
		)
	}

	// var res_cmd, sea_cmd, hel_cmd tea.Cmd
	// m.results, res_cmd = m.results.Update(msg)
	// cmds = append(cmds, res_cmd)
	// m.search, sea_cmd = m.search.Update(msg)
	// cmds = append(cmds, sea_cmd)
	// m.help, hel_cmd = m.help.Update(msg)
	// cmds = append(cmds, hel_cmd)
	return m, tea.Batch(cmds...)
}

func (m model) View() string {
	if m.err != nil {
		return "\n\nError: " + m.err.Error()
	}

	return style.Render(
		lipgloss.JoinVertical(0,
			m.search.View(),
			"\n",
			m.results.View(),
			m.help.View(m.keys),
		))
}

func main() {
	f, err := tea.LogToFile("debug.log", "debug")
	if err != nil {
		log.Fatalf("err: %v", err)
	}
	defer f.Close()
	p := tea.NewProgram(newModel(), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}
