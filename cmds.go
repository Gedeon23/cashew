package main

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/Gedeon23/cashew/recoll"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"os/exec"
)

type CollectMsg struct {
	Results []list.Item
}

func Collect(term string) tea.Cmd {
	return func() tea.Msg {
		return CollectMsg{Results: recoll.Collect(term)}
	}
}

type SnippetsMsg struct {
	Err error
}

func GetSnipptets(entry *recoll.Entry, term string) tea.Cmd {
	return func() tea.Msg {
		return SnippetsMsg{Err: recoll.GetSnipptets(entry, term)}
	}
}

type SwitchEntryMsg struct {
	NewEntry *recoll.Entry
}

type DocViewerMsg struct {
	Viewers map[string]string
}

func GetDocViewers() tea.Cmd {
	return func() tea.Msg {
		docViewers := make(map[string]string, 1)
		docViewers["application/pdf"] = ""

		for docType := range docViewers {
			cmd := exec.Command("xdg-mime", "query", "default", docType)

			out, err := cmd.CombinedOutput()
			if err != nil {
				return fmt.Errorf("Error: could not determine user preferred doc viewers, %s, %s", err.Error(), out)
			}

			scan := bufio.NewScanner(bytes.NewReader(out))
			scan.Scan()

			docViewers[docType] = scan.Text()
		}

		return docViewers
	}
}
