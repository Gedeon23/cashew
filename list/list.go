package list

import (
	tea "github.com/charmbracelet/bubbletea"
)

type Item interface {
	ListView()
}

type Model struct {
	Items    *[]Item
	Selected int

	Height int
	Width  int
}

func (l Model) Init() tea.Cmd {
	return nil
}

func (l Model) Update()
