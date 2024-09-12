package main

import (
	"flag"
	"fmt"
	"log"
	"os/exec"
	"strconv"
	"strings"

	"github.com/Gedeon23/cashew/recoll"
	"github.com/Gedeon23/cashew/styles"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const (
	FocusSearch = iota
	FocusResults
	FocusDetails
	FocusDebug
)

const (
	MetaTab = iota
	SnippetsTab
)

type model struct {
	Search        textinput.Model
	Results       list.Model
	Tabs          []string
	SelectedEntry recoll.Entry

	Keys       GlobalKeyMap
	Help       help.Model
	DocViewers map[string]string

	SelectedTab     int
	Focus           int
	PrevFocus       int
	SelectedSnippet int
	Err             error
	Logging         bool
}

// Create Initial Model for Application
func newModel(query string, logging bool) model {
	search := textinput.New()
	search.Placeholder = "search…"
	search.Prompt = " "
	search.SetValue(query)
	search.Focus()
	search.CharLimit = 200
	search.Width = 20

	var container []list.Item
	results := list.New(container, NewEntryDelegate(), 0, 0)
	results.SetFilteringEnabled(false)
	results.SetShowTitle(false)
	results.SetShowHelp(false)

	keys := NewGlobalKeyMap()
	help := help.New()

	return model{
		Search:  search,
		Results: results,
		Tabs:    []string{"Metadata", "Snippets"},

		Keys: keys,
		Help: help,

		SelectedTab:     0,
		Focus:           FocusSearch,
		SelectedSnippet: 0,

		Logging: logging,
	}
}

// Set's application focus (if details whichever tab is currently active)
// Keyboard input will be directed to this part of the application

func (m *model) SetFocus(focus int) {
	if m.Focus != focus {
		if m.Focus == FocusSearch {
			m.Search.Blur()
		}
		m.PrevFocus = m.Focus
		m.Focus = focus
		m.Keys.Focus = focus
		if m.Focus == FocusSearch {
			m.Search.Focus()
		}
	}
}

func (m *model) UpdateSelectedEntry() {
	if entry, ok := m.Results.SelectedItem().(recoll.Entry); ok {
		if m.SelectedEntry.Url != entry.Url {
			m.SelectedEntry = entry
			m.SelectedSnippet = 0
		}
	}
}

func (m *model) ExpandHelp() {
	prevHeight := strings.Count(m.Help.View(m.Keys), "\n")
	m.Help.ShowAll = !m.Help.ShowAll
	newHeight := strings.Count(m.Help.View(m.Keys), "\n")
	m.Results.SetSize(
		m.Results.Width(),
		m.Results.Height()-(newHeight-prevHeight),
	)
}

func (m *model) OpenSelected() {
	selected := m.Results.SelectedItem()
	if selected, ok := selected.(recoll.Entry); ok {
		cmd := exec.Command("xdg-open", selected.Url)

		if err := cmd.Start(); err != nil && m.Logging {
			log.Printf("Error: %v\n", err)
		}
	}
}

func (m model) Init() tea.Cmd {
	var cmds []tea.Cmd
	if m.Search.Value() != "" {
		cmds = append(cmds, Collect(m.Search.Value()))
	}
	cmds = append(cmds, GetDocViewers())
	return tea.Batch(cmds...)
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd
	var cmd tea.Cmd

	switch msg := msg.(type) {
	// Keybinds
	case tea.KeyMsg:

		// Global Keybinds
		switch {
		case key.Matches(msg, m.Keys.Quit_ESC):
			return nil, tea.Quit
		case key.Matches(msg, m.Keys.Help):
			m.ExpandHelp()
			return m, nil
		case key.Matches(msg, m.Keys.NextTab):
			m.SelectedTab += 1
			m.SelectedTab %= len(m.Tabs)
			m.Keys.SelectedTab = m.SelectedTab
			m.UpdateSelectedEntry()
			if m.SelectedTab == SnippetsTab && len(m.SelectedEntry.Snippets) == 0 && m.SelectedEntry.Url != "" {
				return m, GetSnipptets(m.SelectedEntry, m.Search.Value())
			}
			return m, nil
		case key.Matches(msg, m.Keys.FocusDebug):
			if m.Focus != FocusDebug {
				m.SetFocus(FocusDebug)
			} else {
				m.SetFocus(m.PrevFocus)
			}
			return m, nil
		}

		// Focus Keybinds
		switch m.Focus {
		case FocusSearch:
			switch {
			case key.Matches(msg, m.Keys.ExecuteSearch) && !(m.Search.Value() == ""):
				return m, Collect(m.Search.Value())
			case key.Matches(msg, m.Keys.FocusResultsFromSearch):
				m.SetFocus(FocusResults)
				return m, nil
			case key.Matches(msg, m.Keys.FocusDetailsFromSearch):
				m.SetFocus(FocusDetails)
				return m, nil
			default:
				m.Search, cmd = m.Search.Update(msg)
				cmds = append(cmds, cmd)
			}
		case FocusResults:
			switch {
			case key.Matches(msg, m.Keys.FocusDetails):
				m.SetFocus(FocusDetails)
			case key.Matches(msg, m.Keys.Quit):
				return m, tea.Quit
			case key.Matches(msg, m.Keys.FocusSearch):
				m.SetFocus(FocusSearch)
			case key.Matches(msg, m.Keys.FocusSearchAndClear):
				m.SetFocus(FocusSearch)
				m.Search.SetValue("")
			case key.Matches(msg, m.Keys.Help):
				m.ExpandHelp()
			case key.Matches(msg, m.Keys.OpenDocument):
				m.OpenSelected()
			// case key.Matches(msg, m.Keys.SelectEntry):
			// 	var yml strings.Builder
			// 	yml.WriteString("title: " + m.SelectedEntry.DocTitle + "\n")
			// 	yml.WriteString("author: " + m.SelectedEntry.Author + "\n")
			// 	yml.WriteString("file: " + m.SelectedEntry.File + "\n")
			// 	yml.WriteString("url: " + m.SelectedEntry.Url + "\n")
			// 	yml.WriteString("snippets: \n")
			// 	for _, snip := range m.SelectedEntry.Snippets {
			// 		yml.WriteString("\t- page: " + snip.Page + "\n\t- text: " + snip.Text)
			// 	}
			// 	fmt.Println(yml.String())
			// 	return m, tea.Quit
			default:
				m.Results, cmd = m.Results.Update(msg)
				cmds = append(cmds, cmd)
				m.UpdateSelectedEntry()
				if m.SelectedTab == SnippetsTab && len(m.SelectedEntry.Snippets) == 0 {
					cmds = append(cmds, GetSnipptets(m.SelectedEntry, m.Search.Value()))
				}
			}
		case FocusDetails:
			switch {
			case key.Matches(msg, m.Keys.FocusResults):
				m.SetFocus(FocusResults)
			case key.Matches(msg, m.Keys.FocusSearch):
				m.SetFocus(FocusSearch)
			case key.Matches(msg, m.Keys.Quit):
				return m, tea.Quit
			case key.Matches(msg, m.Keys.FocusSearchAndClear):
				m.SetFocus(FocusSearch)
				m.Search.SetValue("")
				// case key.Matches(msg, m.keys.OpenDocument):
				// 	m.OpenSelected()
				// default:
				// 	m.details, cmd = m.details.Update(msg)
				// 	cmds = append(cmds, cmd)
			}
			switch m.SelectedTab {
			case SnippetsTab:
				switch {
				case key.Matches(msg, m.Keys.NextSnippet):
					if m.SelectedSnippet < len(m.SelectedEntry.Snippets)-1 {
						m.SelectedSnippet++
					}
				case key.Matches(msg, m.Keys.PrevSnippet):
					if m.SelectedSnippet > 0 {
						m.SelectedSnippet--
					}
				case key.Matches(msg, m.Keys.OpenSnippet):
					return m, OpenSnippet(m.SelectedEntry, m.SelectedSnippet)
				}
			case MetaTab:
				return m, nil
			}
		case FocusDebug:
			if key.Matches(msg, m.Keys.Quit) {
				return m, tea.Quit
			}
		}

	// Update on resize
	case tea.WindowSizeMsg:
		h, v := styles.Root.GetFrameSize()
		m.Results.SetSize(
			msg.Width/2-h,
			msg.Height-v-5,
		)
	// catch query results
	case CollectMsg:
		if msg.Err != nil {
			m.Err = msg.Err
			if m.Logging {
				log.Printf("Error: getting recoll entries, %s", msg.Err)
			}
		} else {
			m.Results.SetItems(msg.Results)
			m.SetFocus(FocusResults)
		}
	// catch snippets
	case SnippetsMsg:
		m.Err = msg.Err
		if msg.Entry.Url == m.SelectedEntry.Url {
			m.Results.SetItem(m.Results.Index(), msg.Entry)
			m.SelectedEntry = msg.Entry
		}
	// get preferred doc applications
	case DocViewerMsg:
		m.DocViewers = msg.Viewers
		if msg.Err != nil {
			m.Err = msg.Err
			if m.Logging {
				log.Printf("Error: determining preferred document viewers, %s", msg.Err)
			}
		}
	case SnippetOpenedMsg:
		if msg.Err != nil {
			m.Err = msg.Err
			if m.Logging {
				log.Printf("Error: opening snippet, %s", msg.Err)
			}
		}
	}

	return m, tea.Batch(cmds...)
}

// Returns application view
func (m model) View() string {
	if m.Focus == FocusDebug {
		var s strings.Builder
		s.WriteString(lipgloss.JoinHorizontal(0,
			RenderDebugEntry("Selected Entry", m.SelectedEntry.DocTitle, false),
			"\t",
			RenderDebugEntry("Snippet Count", strconv.Itoa(len(m.SelectedEntry.Snippets)), false),
			"\t",
			RenderDebugEntry("Selected Snippet", strconv.Itoa(m.SelectedSnippet), false),
			"\t",
			RenderDebugEntry("Selected Tab", strconv.Itoa(m.SelectedTab), false),
		))
		s.WriteString("\n\n")

		s.WriteString(lipgloss.JoinHorizontal(0,
			RenderDebugEntry("Focus", strconv.Itoa(m.Focus), false),
			"\t",
			RenderDebugEntry("Previous Focus", strconv.Itoa(m.PrevFocus), false),
		))
		s.WriteString("\n\n")

		s.WriteString(RenderDebugEntry("Document Viewers", fmt.Sprintf("%s", m.DocViewers), false))
		s.WriteString("\n\n")

		if m.Logging {
			s.WriteString(RenderDebugEntry("Logging", " ", false))
		} else {
			s.WriteString(RenderDebugEntry("Logging", " ", false))
		}

		return styles.Root.Render(s.String())
	} else {

		entry, _ := m.Results.SelectedItem().(recoll.Entry)

		// Details View
		// Switch between MetaData View and Snippets View (might add Annotations eventually)
		var details strings.Builder

		for i, tab := range m.Tabs {
			switch {
			case m.Focus == FocusDetails && m.SelectedTab == i:
				details.WriteString(styles.FocusedTab.Render(tab) + "  ")
			case m.SelectedTab == i:
				details.WriteString(styles.SelectedTab.Render(tab) + "  ")
			default:
				details.WriteString(styles.NormalTab.Render(tab) + "  ")
			}
		}
		details.WriteString("\n\n")

		switch m.SelectedTab {
		case MetaTab:
			details.WriteString(entry.View())
		case SnippetsTab:
			if len(entry.Snippets) != 0 {
				for i, snip := range entry.Snippets {
					details.WriteString(RenderSnippet(entry.Query, m.Focus == FocusDetails, m.SelectedSnippet == i, i, snip))
					details.WriteString("\n")
				}
			}
		}

		// Putting everything together
		//
		return styles.Root.Render(
			lipgloss.JoinVertical(0,
				m.Search.View(),
				"\n",
				lipgloss.JoinHorizontal(0,
					m.Results.View(),
					details.String(),
				),
				m.Help.View(m.Keys),
			))
	}
}

func main() {
	log_flag := flag.Bool("log", false, "turn on error logging, usage: --log=true, file: /tmp/cashew_debug.log, default: false")
	search_flag := flag.String("search", "", "predefine search query, usage: --search='<query>'")
	flag.Parse()
	search := *search_flag
	logging := *log_flag

	if logging {
		f, err := tea.LogToFile("/tmp/cashew_debug.log", "debug")
		if err != nil {
			log.Fatalf("err: %v", err)
		}
		defer f.Close()
	}
	p := tea.NewProgram(newModel(search, logging), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}
