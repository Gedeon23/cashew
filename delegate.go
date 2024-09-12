package main

import (
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/Gedeon23/cashew/recoll"
	"github.com/Gedeon23/cashew/styles"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

type EntryDelegate struct {
	height     int
	spacing    int
	ItemStyles list.DefaultItemStyles
}

func (d EntryDelegate) Update(msg tea.Msg, m *list.Model) tea.Cmd {
	return nil
}

func (d EntryDelegate) Height() int {
	return d.height
}

func (d EntryDelegate) Spacing() int {
	return d.spacing
}

func (d EntryDelegate) Render(w io.Writer, m list.Model, index int, entry list.Item) {
	if entry, ok := entry.(recoll.Entry); ok {
		var icon string = " "
		if entry.File[len(entry.File)-3:] == "pdf" {
			icon = " "
		}
		if entry.File[len(entry.File)-3:] == "csv" {
			icon = " "
		}
		title := icon + entry.DocTitle
		if len(title) > (m.Width() - 3) {
			title = title[:m.Width()-4] + "…"
		}

		var snipDisplay string
		snipCount := len(entry.Snippets)
		if snipCount != 0 {
			snipDisplay = "  " + strconv.Itoa(snipCount)
		} else {
			snipDisplay = ""
		}
		auth := " " + entry.Author + snipDisplay
		if len(auth) > (m.Width() - 3) {
			auth = auth[:m.Width()-4] + "…"
		}

		if index == m.Index() {
			title = d.ItemStyles.SelectedTitle.Render(title)
			auth = d.ItemStyles.SelectedDesc.Render(auth)
		} else {
			title = d.ItemStyles.NormalTitle.Render(title)
			auth = d.ItemStyles.NormalDesc.Render(auth)
		}
		fmt.Fprintf(w, "%s\n%s", title, auth)
	} else {
		return
	}

}

func NewEntryDelegate() EntryDelegate {
	return EntryDelegate{
		height:     2,
		spacing:    1,
		ItemStyles: list.NewDefaultItemStyles(),
	}
}

// type SnippetDelegate struct {
// 	height     int
// 	spacing    int
// 	ItemStyles list.DefaultItemStyles
// }

// func (d SnippetDelegate) Update(msg tea.Msg, m *list.Model) tea.Cmd {
// 	return nil
// }

// func (d SnippetDelegate) Height() int {
// 	return d.height
// }

// func (d SnippetDelegate) Spacing() int {
// 	return d.spacing
// }

// func (d SnippetDelegate) Render(w io.Writer, m list.Model, index int, snippet list.Item) {
// 	if snippet, ok := snippet.(recoll.Snippet); ok {
// 		text := " " + snippet.Page + " " + snippet.Text
// 		if len(text) > (m.Width() - 3) {
// 			text = text[:m.Width()-4] + "…"
// 		}
// 		if index == m.Index() {
// 			text = d.ItemStyles.SelectedTitle.Render(text)
// 		} else {
// 			text = d.ItemStyles.NormalTitle.Render(text)
// 		}
// 		fmt.Fprintf(w, text)
// 	}
// }

// func NewSnippetDelegate() SnippetDelegate {
// 	return SnippetDelegate{
// 		height:     1,
// 		spacing:    0,
// 		ItemStyles: list.NewDefaultItemStyles(),
// 	}
// }

func RenderSnippet(query string, focused bool, selected bool, index int, snippet recoll.Snippet) string {
	// if len(text) > (m.Width() - 3) {
	// 	text = text[:m.Width()-4] + "…"
	// }

	var match, before, after string
	queryIndex := strings.Index(strings.ToLower(snippet.Text), strings.ToLower(query))
	if queryIndex != -1 {
		switch {
		case selected && focused:
			before = styles.FocusedSnippet.Render(" " + snippet.Page + " " + snippet.Text[:queryIndex])
			after = styles.FocusedSnippetAfterMatch.Render(snippet.Text[queryIndex+len(query):])
			match = styles.SnippetMatch.Render(snippet.Text[queryIndex : queryIndex+len(query)])
		case selected:
			before = styles.SelectedSnippet.Render(" " + snippet.Page + " " + snippet.Text[:queryIndex])
			after = styles.SelectedSnippetAfterMatch.Render(snippet.Text[queryIndex+len(query):])
			match = styles.SnippetMatch.Render(snippet.Text[queryIndex : queryIndex+len(query)])
		default:
			before = styles.NormalSnippet.Render(" " + snippet.Page + " " + snippet.Text[:queryIndex])
			after = styles.NormalSnippetAfterMatch.Render(snippet.Text[queryIndex+len(query):])
			match = styles.SnippetMatch.Render(snippet.Text[queryIndex : queryIndex+len(query)])
		}
		return before + match + after
	} else {
		switch {
		case selected && focused:
			return styles.FocusedSnippet.Render(" " + snippet.Page + " " + snippet.Text)
		case selected:
			return styles.SelectedSnippet.Render(" " + snippet.Page + " " + snippet.Text)
		default:
			return styles.NormalSnippet.Render(" " + snippet.Page + " " + snippet.Text)
		}
	}
}

func RenderDebugEntry(name string, value string, critical bool) string {
	if critical {
		return styles.CriticalDebugEntryName.Render(name) + "\n" + styles.CriticalDebugEntryValue.Render(value)
	}
	return styles.NormalDebugEntryName.Render(name) + "\n" + styles.NormalDebugEntryValue.Render(value)
}
