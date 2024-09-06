package styles

import (
	"github.com/charmbracelet/lipgloss"
)

var Root = lipgloss.NewStyle().Margin(1, 2)

var EntryField = lipgloss.NewStyle().
	Background(lipgloss.Color("#AA86E8")).
	Foreground(lipgloss.Color("#FFFFFF")).
	Padding(0, 1).
	Bold(true)

var SelectedSnippet = lipgloss.NewStyle().
	Border(lipgloss.NormalBorder(), false, false, false, true).
	BorderForeground(lipgloss.AdaptiveColor{Light: "#F793FF", Dark: "#AD58B4"}).
	Foreground(lipgloss.AdaptiveColor{Light: "#EE6FF8", Dark: "#EE6FF8"}).
	Padding(0, 0, 0, 1)

var NormalSnippet = lipgloss.NewStyle().
	Foreground(lipgloss.AdaptiveColor{Light: "#1a1a1a", Dark: "#dddddd"}).
	Padding(0, 0, 0, 2)

var SelectedSnippetAfterMatch = lipgloss.NewStyle().
	Foreground(lipgloss.AdaptiveColor{Light: "#EE6FF8", Dark: "#EE6FF8"})

var NormalSnippetAfterMatch = lipgloss.NewStyle().
	Foreground(lipgloss.AdaptiveColor{Light: "#1a1a1a", Dark: "#dddddd"})

var SnippetMatch = lipgloss.NewStyle().
	Bold(true).
	Background(lipgloss.Color("#FF0000")).
	Foreground(lipgloss.Color("#FFFFFF"))
