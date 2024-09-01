package entry

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"strings"
)

var fieldStyle = lipgloss.NewStyle().
	Background(lipgloss.Color("#AA86E8")).
	Foreground(lipgloss.Color("#FFFFFF")).
	Padding(0, 1).
	Bold(true)

type Recoll struct {
	Author   string
	DocTitle string
	File     string
	Url      string

	Snippets []string
}

func (e Recoll) FilterValue() string { return "" + e.Url }

func (m Recoll) Init() tea.Cmd {
	return nil
}

func (m Recoll) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return m, nil
}

func (m Recoll) View() string {
	var s strings.Builder
	s.WriteString(fieldStyle.Render(" Title") + "\t" + m.DocTitle + "\n\n")
	s.WriteString(fieldStyle.Render("Author") + "\t" + m.Author + "\n\n")
	s.WriteString(fieldStyle.Render(" File ") + "\t" + m.File + "\n\n")
	s.WriteString(fieldStyle.Render(" Url  ") + "\t" + m.Url + "\n\n")

	return s.String()
}
