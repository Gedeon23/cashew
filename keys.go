package main

import (
	"github.com/charmbracelet/bubbles/key"
)

type KeyMap struct {
	Focus               int
	Quit                key.Binding
	Quit_ESC            key.Binding
	Help                key.Binding
	Help_QM             key.Binding
	ExecuteSearch       key.Binding
	NextEntry           key.Binding
	PrevEntry           key.Binding
	FocusNext           key.Binding
	FocusSearch         key.Binding
	FocusSearchAndClear key.Binding
}

func (k KeyMap) ShortHelp() []key.Binding {
	switch k.Focus {
	case FocusSearch:
		return []key.Binding{k.Quit_ESC, k.Help_QM, k.FocusNext}
	case FocusResults:
		return []key.Binding{k.Quit, k.Help, k.NextEntry, k.PrevEntry, k.FocusSearch}
	case FocusDetail:
		return []key.Binding{k.Quit, k.Help, k.FocusSearch}
	default:
		return []key.Binding{k.Quit, k.FocusSearch, k.Help}
	}
}

func (k KeyMap) FullHelp() [][]key.Binding {
	switch k.Focus {
	case FocusSearch:
		return [][]key.Binding{
			{k.ExecuteSearch},
			{k.Quit_ESC, k.Help_QM, k.FocusNext},
		}
	case FocusResults:
		return [][]key.Binding{
			{k.FocusSearch, k.FocusSearchAndClear, k.FocusNext},
			{k.Quit, k.Help, k.NextEntry, k.PrevEntry},
		}
	case FocusDetail:
		return [][]key.Binding{
			{k.FocusSearch, k.FocusSearchAndClear, k.FocusNext},
			{k.Quit, k.Help},
		}
	default:
		return [][]key.Binding{
			{k.FocusSearch, k.ExecuteSearch, k.NextEntry, k.PrevEntry},
			{k.FocusNext, k.Quit, k.Help},
		}
	}
}

func NewDefaultKeyMap() KeyMap {
	return KeyMap{
		Quit: key.NewBinding(
			key.WithKeys("q", "esc"),
			key.WithHelp("q/esc", "quit"),
		),
		Quit_ESC: key.NewBinding(
			key.WithKeys("esc"),
			key.WithHelp("esc", "quit"),
		),
		Help: key.NewBinding(
			key.WithKeys("?", "h"),
			key.WithHelp("?/h", "help"),
		),
		Help_QM: key.NewBinding(
			key.WithKeys("?"),
			key.WithHelp("?", "help"),
		),
		FocusSearch: key.NewBinding(
			key.WithKeys("f", "ctrl+f", "/"),
			key.WithHelp("f", "search"),
		),
		FocusSearchAndClear: key.NewBinding(
			key.WithKeys("F"),
			key.WithHelp("F", "new search"),
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
