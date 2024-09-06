package main

import (
	"log"
	"os/exec"
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
	Entry           *recoll.Entry
	SelectedSnippet int
	Pager           paginator.Model
	Err             error
}

func NewDetails() Details {
	pager := paginator.New()
	pager.Type = paginator.Dots
	pager.PerPage = 1
	pager.TotalPages = 2
	var entry *recoll.Entry

	return Details{
		Entry:           entry,
		SelectedSnippet: 0,
		Pager:           pager,
		Err:             nil,
	}
}

func (d Details) Init() tea.Cmd {
	return nil
}

func (d Details) Update(msg tea.Msg) (Details, tea.Cmd) {
	var cmds []tea.Cmd
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case SnippetsMsg:
		d.Err = msg.Err
	case SwitchEntryMsg:
		d.Entry = msg.NewEntry
		if d.Pager.Page == SnippetsPage && len(d.Entry.Snippets) == 0 {
			cmd = GetSnipptets(d.Entry, d.Entry.Query)
			cmds = append(cmds, cmd)
			d.SelectedSnippet = 0
		}
	case tea.KeyMsg:
		switch {
		case msg.String() == "j":
			if d.SelectedSnippet < len(d.Entry.Snippets)-1 {
				d.SelectedSnippet++
			}
		case msg.String() == "k":
			if d.SelectedSnippet > 0 {
				d.SelectedSnippet--
			}
		case msg.String() == "o":
			if d.Pager.Page == SnippetsPage {
				cmd := exec.Command("zathura", "--page="+strings.TrimSpace(d.Entry.Snippets[d.SelectedSnippet].Page), d.Entry.Url)
				if err := cmd.Start(); err != nil {
					log.Printf("Error: %v", err)
				}
			} else {
				cmd := exec.Command("xdg-open", d.Entry.Url)
				if err := cmd.Start(); err != nil {
					log.Printf("Error: %v", err)
				}
			}
		default:
			d.Pager, cmd = d.Pager.Update(msg)
			cmds = append(cmds, cmd)
			if d.Pager.Page == SnippetsPage && len(d.Entry.Snippets) == 0 {
				cmd = GetSnipptets(d.Entry, d.Entry.Query)
				cmds = append(cmds, cmd)
			}
		}
	default:
		d.Pager, cmd = d.Pager.Update(msg)
		cmds = append(cmds, cmd)
		if d.Pager.Page == SnippetsPage && len(d.Entry.Snippets) == 0 {
			cmd = GetSnipptets(d.Entry, d.Entry.Query)
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
				for i, snip := range d.Entry.Snippets {
					s.WriteString(RenderSnippet(d.Entry.Query, d.SelectedSnippet == i, i, snip))
					s.WriteString("\n")
				}
			}
		default:
			s.WriteString(d.Entry.View())
		}
	}
	return s.String()
}
