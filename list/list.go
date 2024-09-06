package list

import (
	"github.com/Gedeon23/cashew/recoll"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"strings"
)

type Model struct {
	Items    []recoll.Entry
	Selected int

	Keys KeyMap

	Height int
	Width  int
}

func New() Model {
	return Model{
		Keys:   NewKeyMap(),
		Height: 0,
		Width:  0,
	}
}

func (l Model) Init() tea.Cmd {
	return nil
}

func (l Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	// var cmd tea.Cmd
	// var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, l.Keys.NextEntry):
			l.Selected++
		case key.Matches(msg, l.Keys.PrevEntry):
			l.Selected--
		}
	}

	return l, nil
	// return l, tea.Batch(cmds...)
}

func (l Model) View() string {
	var s strings.Builder

	return s.String()
}
