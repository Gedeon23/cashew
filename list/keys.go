package list

import (
	"github.com/charmbracelet/bubbles/key"
)

type KeyMap struct {
	NextEntry key.Binding
	PrevEntry key.Binding
}

func NewKeyMap() KeyMap {
	return KeyMap{
		NextEntry: key.NewBinding(
			key.WithKeys("ctrl+n", "j", "down"),
			key.WithHelp("next entry", "j"),
		),
		PrevEntry: key.NewBinding(
			key.WithKeys("ctrl+p", "k", "up"),
			key.WithHelp("prev entry", "k"),
		),
	}
}

func (k KeyMap) ShortHelp() []key.Binding {
	return []key.Binding{
		k.NextEntry,
		k.PrevEntry,
	}
}

func (k KeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{
			k.NextEntry,
			k.PrevEntry,
		},
	}
}
