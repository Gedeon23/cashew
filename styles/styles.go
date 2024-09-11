package styles

import (
	"github.com/charmbracelet/lipgloss"
)

var FocusedForeground = lipgloss.AdaptiveColor{Light: "#EE6FF8", Dark: "#EE6FF8"}
var SelectedForeground = lipgloss.AdaptiveColor{Light: "#1a1a1a", Dark: "#eeeeee"}
var NormalForeground = lipgloss.AdaptiveColor{Light: "#3a3a3a", Dark: "#bbbbbb"}

// TODO cleanup colors into vars

var Root = lipgloss.NewStyle().Margin(1, 2)

var EntryField = lipgloss.NewStyle().
	Background(lipgloss.Color("#AA86E8")).
	Foreground(lipgloss.Color("#FFFFFF")).
	Padding(0, 1).
	Bold(true)

var FocusedSnippet = lipgloss.NewStyle().
	Border(lipgloss.NormalBorder(), false, false, false, true).
	BorderForeground(FocusedForeground).
	Foreground(FocusedForeground).
	Padding(0, 0, 0, 1).
	Bold(true)

var SelectedSnippet = lipgloss.NewStyle().
	Border(lipgloss.NormalBorder(), false, false, false, true).
	BorderForeground(lipgloss.AdaptiveColor{Light: "#F793FF", Dark: "#AD58B4"}).
	Foreground(SelectedForeground).
	Padding(0, 0, 0, 1).
	Bold(true)

var NormalSnippet = lipgloss.NewStyle().
	Foreground(lipgloss.AdaptiveColor{Light: "#1a1a1a", Dark: "#dddddd"}).
	Padding(0, 0, 0, 2)

var FocusedSnippetAfterMatch = lipgloss.NewStyle().
	Foreground(FocusedForeground).
	Bold(true)

var SelectedSnippetAfterMatch = lipgloss.NewStyle().
	Foreground(SelectedForeground).
	Bold(true)

var NormalSnippetAfterMatch = lipgloss.NewStyle().
	Foreground(NormalForeground)

var SnippetMatch = lipgloss.NewStyle().
	Bold(true).
	Background(lipgloss.Color("#FF0000")).
	Foreground(lipgloss.Color("#FFFFFF"))

var SelectedTab = lipgloss.NewStyle().
	Bold(true).
	Foreground(SelectedForeground)

var NormalTab = lipgloss.NewStyle().
	Foreground(NormalForeground)

var FocusedTab = lipgloss.NewStyle().
	Bold(true).
	Foreground(FocusedForeground)

var NormalDebugEntryName = lipgloss.NewStyle().
	Foreground(NormalForeground).
	Bold(true)

var NormalDebugEntryValue = lipgloss.NewStyle().
	Foreground(NormalForeground)

var CriticalDebugEntryName = lipgloss.NewStyle().
	Foreground(lipgloss.Color("#FF0000")).
	Bold(true)

var CriticalDebugEntryValue = lipgloss.NewStyle().
	Foreground(lipgloss.Color("#FF0000")).
	Bold(true)
