package main

import (
	"github.com/charmbracelet/bubbles/key"
)

type KeyMap interface {
	ShortHelp()
	FullHelp()
}

type GlobalKeyMap struct {
	Quit        key.Binding
	Help        key.Binding
	FocusSearch key.Binding
	FocusNext   key.Binding
}

func (k GlobalKeyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.Quit, k.FocusSearch, k.Help}
}

func (k GlobalKeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.FocusSearch},
		{k.FocusNext, k.Quit, k.Help},
	}
}

func NewGlobalKeyMap() GlobalKeyMap {
	return GlobalKeyMap{
		Quit: key.NewBinding(
			key.WithKeys("esc"),
			key.WithHelp("", "ESC"),
		),
		Help: key.NewBinding(
			key.WithKeys("?"),
			key.WithHelp("?", "help"),
		),
		FocusSearch: key.NewBinding(
			key.WithKeys("ctrl+f"),
			key.WithHelp("C-f", "search"),
		),
		FocusNext: key.NewBinding(
			key.WithKeys("tab"),
			key.WithHelp("TAB", "next panel"),
		),
	}
}

type SearchKeyMap struct {
	ExecuteSearch key.Binding
}

func (k SearchKeyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.ExecuteSearch}
}

func (k SearchKeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.ExecuteSearch},
	}
}

func NewSearchKeyMap() SearchKeyMap {
	return SearchKeyMap{
		ExecuteSearch: key.NewBinding(
			key.WithKeys("enter"),
			key.WithHelp("enter", "search recoll"),
		),
	}
}

type ResultsKeyMap struct {
	NextEntry key.Binding
	PrevEntry key.Binding
}

func (k ResultsKeyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.NextEntry, k.PrevEntry}
}

func (k ResultsKeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.NextEntry, k.PrevEntry},
	}
}

func NewResultsKeyMap() ResultsKeyMap {
	return ResultsKeyMap{
		NextEntry: key.NewBinding(
			key.WithKeys("j", "down"),
			key.WithHelp("j", "down"),
		),
		PrevEntry: key.NewBinding(
			key.WithKeys("k", "up"),
			key.WithHelp("k", "up"),
		),
	}
}

type AdaptiveKeyMap struct {
	Global  GlobalKeyMap
	Focus   int
	Search  SearchKeyMap
	Results ResultsKeyMap
}

func (k AdaptiveKeyMap) ShortHelp() []key.Binding {
	switch k.Focus {
	case FocusSearch:
		return append(k.Search.ShortHelp(), k.Global.ShortHelp()...)
	case FocusResults:
		return append(k.Results.ShortHelp(), k.Global.ShortHelp()...)
	}
	return nil
}

func (k AdaptiveKeyMap) FullHelp() [][]key.Binding {
	switch k.Focus {
	case FocusSearch:
		return append(k.Search.FullHelp(), k.Global.FullHelp()...)
	case FocusResults:
		return append(k.Results.FullHelp(), k.Global.FullHelp()...)
	}
	return nil
}

func NewAdaptiveKeyMap() AdaptiveKeyMap {
	return AdaptiveKeyMap{
		Global:  NewGlobalKeyMap(),
		Focus:   FocusSearch,
		Search:  NewSearchKeyMap(),
		Results: NewResultsKeyMap(),
	}
}
