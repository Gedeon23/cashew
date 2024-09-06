package recoll

import (
	"fmt"
	"github.com/Gedeon23/cashew/styles"
	"strings"
)

type Entry struct {
	Author   string
	DocTitle string
	File     string
	Url      string

	Snippets []string
}

func (e Entry) FilterValue() string { return "" + e.Url }

// func (m Entry) Init() tea.Cmd {
// 	return nil
// }

// func (m Entry) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
// 	return m, nil
// }

func (m Entry) View() string {
	var s strings.Builder
	s.WriteString(styles.EntryField.Render(" Title") + "\t" + m.DocTitle + "\n\n")
	s.WriteString(styles.EntryField.Render("Author") + "\t" + m.Author + "\n\n")
	s.WriteString(styles.EntryField.Render(" File ") + "\t" + m.File + "\n\n")
	s.WriteString(styles.EntryField.Render(" Url  ") + "\t" + m.Url + "\n\n")
	s.WriteString(fmt.Sprintf("Snippets (-> %s): %s", &m, m.Snippets))

	return s.String()
}
