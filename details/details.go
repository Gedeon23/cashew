package details

import (
	"strings"

	"github.com/Gedeon23/cashew/entry"
	"github.com/charmbracelet/bubbles/paginator"
	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	Entry entry.Recoll
	Pager paginator.Model
}

func New(title string, author string, file string, url string) Model {
	pager := paginator.New()
	pager.Type = paginator.Dots
	pager.PerPage = 1
	var entry entry.Recoll

	return Model{
		Entry: entry,
		Pager: pager,
	}
}

func (d *Model) SetEntry(e entry.Recoll) {

}

func (d Model) Init() tea.Cmd {
	return nil
}

func (d Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd
	var cmd tea.Cmd

	d.Pager, cmd = d.Pager.Update(msg)
	cmds = append(cmds, cmd)

	return d, tea.Batch(cmds...)
}

func (d Model) View() string {
	var s strings.Builder
	switch d.Pager.Page {
	default:
		s.WriteString(d.Entry.View())
	}
	s.WriteString(d.Pager.View())
	return s.String()
}
