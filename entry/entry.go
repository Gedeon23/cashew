package entry

import (
	tea "github.com/charmbracelet/bubbletea"
	"strings"
)

type Recoll struct {
	Author   string
	DocTitle string
	File     string
	Url      string
}

func (e Recoll) Title() string {
	var icon string = " "
	if e.File[len(e.File)-3:] == "pdf" {
		icon = " "
	}
	return icon + e.DocTitle
}
func (e Recoll) Description() string { return " " + e.Author }
func (e Recoll) FilterValue() string { return "" + e.Url }

func (m Recoll) Init() tea.Cmd {
	return nil
}

func (m Recoll) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return m, nil
}

func (m Recoll) View() string {
	var s strings.Builder
	s.WriteString("Title | " + m.DocTitle + "\n\n")
	s.WriteString("Author | " + m.Author + "\n\n")
	s.WriteString("File | " + m.File + "\n\n")
	s.WriteString("Url | " + m.Url + "\n\n")

	return s.String()
}
