package styles

import (
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/lipgloss"
)

var Root = lipgloss.NewStyle().Margin(1, 2)

var EntryField = lipgloss.NewStyle().
	Background(lipgloss.Color("#AA86E8")).
	Foreground(lipgloss.Color("#FFFFFF")).
	Padding(0, 1).
	Bold(true)

var Snippet = list.NewDefaultItemStyles()
var SelectedSnippet = Snippet.SelectedTitle
var NormalSnippet = Snippet.NormalTitle
