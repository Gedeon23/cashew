package main

import (
	"github.com/charmbracelet/bubbles/key"
)

type KeyMap struct {
	Focus                  int
	Quit                   key.Binding
	Quit_ESC               key.Binding
	Help                   key.Binding
	Help_QM                key.Binding
	ExecuteSearch          key.Binding
	NextEntry              key.Binding
	PrevEntry              key.Binding
	FocusSearch            key.Binding
	FocusSearchAndClear    key.Binding
	FocusResults           key.Binding
	FocusResultsFromSearch key.Binding
	FocusDetails           key.Binding
	FocusDetailsFromSearch key.Binding
	OpenDocument           key.Binding
}

func (k KeyMap) ShortHelp() []key.Binding {
	switch k.Focus {
	case FocusSearch:
		return []key.Binding{k.Quit_ESC, k.Help_QM, k.FocusDetails, k.FocusResults}
	case FocusResults:
		return []key.Binding{k.Quit, k.Help, k.NextEntry, k.PrevEntry, k.OpenDocument, k.FocusSearch}
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
			{k.Quit_ESC, k.Help_QM, k.FocusDetails, k.FocusResults},
		}
	case FocusResults:
		return [][]key.Binding{
			{k.OpenDocument, k.FocusSearch, k.FocusSearchAndClear, k.FocusDetails},
			{k.Quit, k.Help, k.NextEntry, k.PrevEntry},
		}
	case FocusDetail:
		return [][]key.Binding{
			{k.FocusSearch, k.FocusSearchAndClear, k.FocusResults},
			{k.Quit, k.Help},
		}
	default:
		return [][]key.Binding{
			{k.FocusSearch, k.ExecuteSearch, k.NextEntry, k.PrevEntry},
			{k.FocusResults, k.FocusDetails, k.Quit, k.Help},
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
		FocusResults: key.NewBinding(
			key.WithKeys("s", "ctrl+s"),
			key.WithHelp("s", "focus results"),
		),
		FocusResultsFromSearch: key.NewBinding(
			key.WithKeys("ctrl+s"),
			key.WithHelp("ctrl+s", "focus results"),
		),
		FocusDetails: key.NewBinding(
			key.WithKeys("g", "ctrl+g"),
			key.WithHelp("g", "focus detail"),
		),
		FocusDetailsFromSearch: key.NewBinding(
			key.WithKeys("ctrl+g"),
			key.WithHelp("ctrl+g", "focus detail"),
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
		OpenDocument: key.NewBinding(
			key.WithKeys("o"),
			key.WithHelp("o", "open doc"),
		),
	}
}
