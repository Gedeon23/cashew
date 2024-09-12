package main

import (
	"github.com/charmbracelet/bubbles/key"
)

type GlobalKeyMap struct {
	Focus                  int
	SelectedTab            int
	Quit                   key.Binding
	Quit_ESC               key.Binding
	Help                   key.Binding
	ExecuteSearch          key.Binding
	NextTab                key.Binding
	NextEntry              key.Binding
	PrevEntry              key.Binding
	FocusSearch            key.Binding
	FocusSearchAndClear    key.Binding
	FocusResults           key.Binding
	FocusResultsFromSearch key.Binding
	FocusDetails           key.Binding
	FocusDetailsFromSearch key.Binding
	OpenDocument           key.Binding
	NextSnippet            key.Binding
	PrevSnippet            key.Binding
	OpenSnippet            key.Binding
	FocusDebug             key.Binding
	SelectEntry            key.Binding
}

func (k GlobalKeyMap) ShortHelp() []key.Binding {
	switch k.Focus {
	case FocusSearch:
		return []key.Binding{k.Quit_ESC, k.Help, k.FocusDetails, k.FocusResults}
	case FocusResults:
		return []key.Binding{k.Quit, k.Help, k.NextEntry, k.PrevEntry, k.OpenDocument, k.FocusSearch}
	case FocusDetails:
		return []key.Binding{k.Quit, k.Help, k.FocusSearch}
	default:
		return []key.Binding{k.Quit, k.FocusSearch, k.Help}
	}
}

func (k GlobalKeyMap) FullHelp() [][]key.Binding {
	switch k.Focus {
	case FocusSearch:
		return [][]key.Binding{
			{k.ExecuteSearch},
			{k.Quit_ESC, k.Help, k.FocusDetails, k.FocusResults},
		}
	case FocusResults:
		return [][]key.Binding{
			{k.OpenDocument, k.FocusSearch, k.FocusSearchAndClear, k.FocusDetails},
			{k.Quit, k.Help, k.NextEntry, k.PrevEntry},
		}
	case FocusDetails:
		if k.SelectedTab == SnippetsTab {
			return [][]key.Binding{
				{k.FocusSearch, k.FocusResults, k.PrevSnippet},
				{k.Quit, k.Help, k.OpenSnippet, k.NextSnippet},
			}
		} else {
			return [][]key.Binding{
				{k.FocusSearch, k.FocusSearchAndClear, k.FocusResults},
				{k.Quit, k.Help},
			}
		}
	default:
		return [][]key.Binding{
			{k.FocusSearch, k.ExecuteSearch, k.NextEntry, k.PrevEntry},
			{k.FocusResults, k.FocusDetails, k.Quit, k.Help},
		}
	}
}

func NewGlobalKeyMap() GlobalKeyMap {
	return GlobalKeyMap{
		Quit: key.NewBinding(
			key.WithKeys("q", "esc"),
			key.WithHelp("q/esc", "quit"),
		),
		Quit_ESC: key.NewBinding(
			key.WithKeys("esc"),
			key.WithHelp("esc", "quit"),
		),
		Help: key.NewBinding(
			key.WithKeys("?"),
			key.WithHelp("?", "help"),
		),
		NextTab: key.NewBinding(
			key.WithKeys("tab"),
			key.WithHelp("tab", "next tab"),
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
		NextSnippet: key.NewBinding(
			key.WithKeys("j", "ctrl+n"),
			key.WithHelp("j", "next snippet"),
		),
		PrevSnippet: key.NewBinding(
			key.WithKeys("k", "ctrl+p"),
			key.WithHelp("k", "prev snippet"),
		),
		OpenSnippet: key.NewBinding(
			key.WithKeys("o"),
			key.WithHelp("o", "open snippet"),
		),
		FocusDebug: key.NewBinding(
			key.WithKeys("!"),
			key.WithHelp("!", "open debug"),
		),
		SelectEntry: key.NewBinding(
			key.WithKeys("enter"),
			key.WithHelp("Enter", "select entry"),
		),
	}
}
