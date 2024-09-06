package main

import (
	"strings"

	"github.com/Gedeon23/cashew/recoll"
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

type Details struct {
	Entry *recoll.Entry
	Pager paginator.Model
	Query string
	Err   error
}

func NewDetails() Details {
	pager := paginator.New()
	pager.Type = paginator.Dots
	pager.PerPage = 1
	pager.TotalPages = 2
	var entry *recoll.Entry

	return Details{
		Entry: entry,
		Pager: pager,
		Query: "",
		Err:   nil,
	}
}

func (d Details) Init() tea.Cmd {
	return nil
}

func (d Details) Update(msg tea.Msg) (Details, tea.Cmd) {
	var cmds []tea.Cmd
	var cmd tea.Cmd

	switch msg.(type) {
	case SnippetsMsg:
		d.Err = msg.(SnippetsMsg).Err
	case SwitchEntryMsg:
		d.Entry = msg.(SwitchEntryMsg).NewEntry
		if d.Pager.Page == SnippetsPage && len(d.Entry.Snippets) == 0 {
			cmd = GetSnipptets(d.Entry, d.Query)
			cmds = append(cmds, cmd)
		}
	default:
		d.Pager, cmd = d.Pager.Update(msg)
		cmds = append(cmds, cmd)
		if d.Pager.Page == SnippetsPage && len(d.Entry.Snippets) == 0 {
			cmd = GetSnipptets(d.Entry, d.Query)
			cmds = append(cmds, cmd)
		}
	}

	return d, tea.Batch(cmds...)
}

func (d Details) View() string {
	if d.Err != nil {
		return d.Err.Error()
	}

	var s strings.Builder
	s.WriteString(centerStyle.Render(d.Pager.View()))
	s.WriteString("\n\n")
	if d.Entry != nil {
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
	}
	return s.String()
}
