package main

import (
	"bufio"
	"bytes"
	"github.com/Gedeon23/cashew/recoll"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"os/exec"
	"strings"
)

type CollectMsg struct {
	Results []list.Item
	Err     error
}

func Collect(term string) tea.Cmd {
	return func() tea.Msg {
		results, err := recoll.Collect(term)
		if err != nil {
			return CollectMsg{Err: err}
		}
		return CollectMsg{Results: results}
	}
}

type SnippetsMsg struct {
	Entry recoll.Entry
	Err   error
}

func GetSnipptets(entry recoll.Entry, term string) tea.Cmd {
	return func() tea.Msg {
		entryWithSnippets, err := recoll.GetSnipptets(entry, term)
		return SnippetsMsg{Entry: entryWithSnippets, Err: err}
	}
}

type SwitchEntryMsg struct {
	NewEntry *recoll.Entry
}

type DocViewerMsg struct {
	Viewers map[string]string
	Err     error
}

func GetDocViewers() tea.Cmd {
	return func() tea.Msg {
		docViewers := make(map[string]string, 1)
		docViewers["application/pdf"] = ""

		for docType := range docViewers {
			cmd := exec.Command("xdg-mime", "query", "default", docType)

			out, err := cmd.CombinedOutput()
			if err != nil {
				return DocViewerMsg{
					Err: err,
				}
			}

			scan := bufio.NewScanner(bytes.NewReader(out))
			scan.Scan()

			docViewers[docType] = scan.Text()
		}

		return DocViewerMsg{
			Viewers: docViewers,
		}
	}
}

type SnippetOpenedMsg struct {
	Err error
}

func OpenSnippet(Entry recoll.Entry, SelectedSnippet int) tea.Cmd {
	return func() tea.Msg {
		cmd := exec.Command("zathura", "--page="+strings.TrimSpace(Entry.Snippets[SelectedSnippet].Page), Entry.Url)
		if err := cmd.Start(); err != nil {
			return SnippetOpenedMsg{Err: err}
		}

		return SnippetOpenedMsg{}
	}
}
