package main

import (
	"github.com/charmbracelet/bubbles/key"
)

type KeyMap struct {
	Quit          key.Binding
	Help          key.Binding
	ExecuteSearch key.Binding
	NextEntry     key.Binding
	PrevEntry     key.Binding
	FocusNext     key.Binding
	FocusSearch   key.Binding
}

func (k KeyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.Quit, k.FocusSearch, k.Help}
}

func (k KeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.FocusSearch, k.ExecuteSearch, k.NextEntry, k.PrevEntry},
		{k.FocusNext, k.Quit, k.Help},
	}
}

func NewDefaultKeyMap() KeyMap {
	return KeyMap{
		Quit: key.NewBinding(
			key.WithKeys("q", "esc"),
			key.WithHelp("q/esc", "quit"),
		),
		Help: key.NewBinding(
			key.WithKeys("?", "h"),
			key.WithHelp("?/h", "help"),
		),
		FocusSearch: key.NewBinding(
			key.WithKeys("f", "ctrl+f", "/"),
			key.WithHelp("f", "search"),
		),
		ExecuteSearch: key.NewBinding(
			key.WithKeys("enter"),
			key.WithHelp("enter", "execute search"),
		),
		NextEntry: key.NewBinding(
			key.WithKeys("j", "ctrl+n"),
			key.WithHelp("j", "next"),
		),
		PrevEntry: key.NewBinding(
			key.WithKeys("k", "ctrl+p"),
			key.WithHelp("k", "previous"),
		),
		FocusNext: key.NewBinding(
			key.WithKeys("tab"),
			key.WithHelp("TAB", "focus next"),
		),
	}
}
