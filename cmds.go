package main

import (
	"github.com/Gedeon23/cashew/recoll"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

type CollectMsg struct {
	Results []list.Item
}

func Collect(term string) tea.Cmd {
	return func() tea.Msg {
		return CollectMsg{Results: recoll.Collect(term)}
	}
}

type SnippetsMsg struct {
	Err error
}

func GetSnipptets(entry *recoll.Entry, term string) tea.Cmd {
	return func() tea.Msg {
		return SnippetsMsg{Err: recoll.GetSnipptets(entry, term)}
	}
}
