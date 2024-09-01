package details

import (
	"strings"

	"github.com/Gedeon23/cashew/entry"
	"github.com/charmbracelet/bubbles/paginator"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const (
	MetaPage = iota
	SnippetsPage
)

var centerStyle = lipgloss.NewStyle().
	Align(lipgloss.Center)

type Model struct {
	Entry entry.Recoll
	Pager paginator.Model
	Query string
}

func New() Model {
	pager := paginator.New()
	pager.Type = paginator.Dots
	pager.PerPage = 1
	pager.TotalPages = 2
	var entry entry.Recoll

	return Model{
		Entry: entry,
		Pager: pager,
		Query: "",
	}
}

func (d Model) Init() tea.Cmd {
	return nil
}

func (d Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	var cmds []tea.Cmd
	var cmd tea.Cmd

	d.Pager, cmd = d.Pager.Update(msg)
	cmds = append(cmds, cmd)

	if d.Pager.Page == SnippetsPage && len(d.Entry.Snippets) == 0 {
		GetSnipptets(d.Entry.Url, d.Query, &d.Entry.Snippets)
	}

	return d, tea.Batch(cmds...)
}

func (d Model) View() string {
	var s strings.Builder
	s.WriteString(centerStyle.Render(d.Pager.View()))
	s.WriteString("\n\n")
	switch d.Pager.Page {
	case MetaPage:
		s.WriteString(d.Entry.View())
	case SnippetsPage:
		if len(d.Entry.Snippets) != 0 {
			for _, snip := range d.Entry.Snippets {
				s.WriteString(snip)
				s.WriteString("\n")
			}
		}
	default:
		s.WriteString(d.Entry.View())
	}
	return s.String()
}
