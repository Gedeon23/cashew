package main

import (
	"log"
	"os/exec"
	"strings"

	"github.com/Gedeon23/cashew/recoll"
	"github.com/Gedeon23/cashew/styles"
	tea "github.com/charmbracelet/bubbletea"
)

const (
	MetaTab = iota
	SnippetsTab
)

type Details struct {
	Tabs            []string
	TabIndex        int
	Focused         bool
	Entry           *recoll.Entry
	SelectedSnippet int
	Err             error
}

func NewDetails() Details {
	var entry *recoll.Entry

	return Details{
		Tabs:            []string{"Metadata", "Snippets"},
		TabIndex:        0,
		Focused:         false,
		Entry:           entry,
		SelectedSnippet: 0,
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
		if d.TabIndex == SnippetsTab && len(d.Entry.Snippets) == 0 {
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
			if d.TabIndex == SnippetsTab {
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
		case msg.String() == "tab":
			d.TabIndex = (d.TabIndex + 1) % len(d.Tabs)
			if d.TabIndex == SnippetsTab && len(d.Entry.Snippets) == 0 {
				cmd = GetSnipptets(d.Entry, d.Entry.Query)
				cmds = append(cmds, cmd)
			}
		default:
			if d.TabIndex == SnippetsTab && len(d.Entry.Snippets) == 0 {
				cmd = GetSnipptets(d.Entry, d.Entry.Query)
				cmds = append(cmds, cmd)
			}
		}
	default:
		if d.TabIndex == SnippetsTab && len(d.Entry.Snippets) == 0 {
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
	s.WriteString("\n\n")
	for i, tab := range d.Tabs {
		if i == d.TabIndex && d.Focused {
			s.WriteString(styles.FocusedTab.Render(tab) + "  ")
		} else if i == d.TabIndex {
			s.WriteString(styles.SelectedTab.Render(tab) + "  ")
		} else {
			s.WriteString(styles.NormalTab.Render(tab) + "  ")
		}
	}
	s.WriteString("\n\n")
	if d.Entry != nil {
		switch d.TabIndex {
		case MetaTab:
			s.WriteString(d.Entry.View())
		case SnippetsTab:
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
