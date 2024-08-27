package main

import (
	"github.com/charmbracelet/bubbles/list"
	// tea "github.com/charmbracelet/bubbletea"
)

func NewEntryDelegate() list.DefaultDelegate {
	d := list.NewDefaultDelegate()
	return d
}
