package entry

import (
	"fmt"
	"io"
	"strconv"

	"github.com/Gedeon23/cashew/recoll"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

type Delegate struct {
	height     int
	spacing    int
	ItemStyles list.DefaultItemStyles
}

func (d Delegate) Update(msg tea.Msg, m *list.Model) tea.Cmd {
	return nil
}

func (d Delegate) Height() int {
	return d.height
}

func (d Delegate) Spacing() int {
	return d.spacing
}

func (d Delegate) Render(w io.Writer, m list.Model, index int, entry list.Item) {
	if entry, ok := entry.(recoll.Entry); ok {
		var icon string = " "
		if entry.File[len(entry.File)-3:] == "pdf" {
			icon = " "
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

func NewEntryDelegate() Delegate {
	return Delegate{
		height:     2,
		spacing:    1,
		ItemStyles: list.NewDefaultItemStyles(),
	}
}
