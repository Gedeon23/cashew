package main

import (
	"log"

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

type model struct {
	search   textinput.Model
	results  list.Model
	keys     KeyMap
	help     help.Model
	selected int
	err      error
}

func newModel() model {
	var items []list.Item
	search := textinput.New()
	search.Placeholder = " search…"
	search.Focus()
	search.CharLimit = 200
	search.Width = 20
	list := list.New(items, NewEntryDelegate(), 0, 0)
	list.Title = "Results"
	list.SetFilteringEnabled(false)
	list.SetShowHelp(false)
	keys := NewDefaultKeyMap()
	help := help.New()
	return model{
		search:  search,
		results: list,
		keys:    keys,
		help:    help,
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd
	if keyMsg, ok := msg.(tea.KeyMsg); ok {
		switch {
		case key.Matches(keyMsg, m.keys.FocusSearch):
			m.search.Focus()
		case key.Matches(keyMsg, m.keys.ExecuteSearch):
			m.results.SetItems(Collect(m.search.Value()))
			m.search.Blur()
		case key.Matches(keyMsg, m.keys.Quit):
			return m, tea.Quit
		case key.Matches(keyMsg, m.keys.Help):
			m.help.ShowAll = !m.help.ShowAll
		}
	}
	if windowSizeMsg, ok := msg.(tea.WindowSizeMsg); ok {
		h, v := style.GetFrameSize()
		m.results.SetSize(windowSizeMsg.Width-h, windowSizeMsg.Height-v-3)
	}

	var res_cmd, sea_cmd tea.Cmd
	m.results, res_cmd = m.results.Update(msg)
	cmds = append(cmds, res_cmd)
	m.search, sea_cmd = m.search.Update(msg)
	cmds = append(cmds, sea_cmd)
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
