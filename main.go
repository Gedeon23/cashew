package main

import (
	"log"
	"strings"

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
	return model{
		search:  search,
		results: list,
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd
	if keyMsg, ok := msg.(tea.KeyMsg); ok {
		switch keyMsg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "enter":
			m.results.SetItems(Collect(m.search.Value()))
			log.Printf("list items: %s", m.results.Items())
		case "ctrl+n", "down", "j":
			if m.selected < len(m.results.Items()) {
				m.selected++
			}
		case "ctrl+p", "up", "k":
			if m.selected > 0 {
				m.selected--
			}
		}
	}
	if windowSizeMsg, ok := msg.(tea.WindowSizeMsg); ok {
		h, v := style.GetFrameSize()
		m.results.SetSize(windowSizeMsg.Width-h, windowSizeMsg.Height-v)
	}

	var res_cmd tea.Cmd
	m.results, res_cmd = m.results.Update(msg)
	cmds = append(cmds, res_cmd)
	var sea_cmd tea.Cmd
	m.search, sea_cmd = m.search.Update(msg)
	cmds = append(cmds, sea_cmd)
	return m, tea.Batch(cmds...)
}

func (m model) View() string {
	var s strings.Builder
	s.WriteString(m.search.View())
	s.WriteString(m.results.View())
	if m.err != nil {
		s.WriteString("\n\nError: " + m.err.Error())
	}
	return style.Render(s.String())
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
