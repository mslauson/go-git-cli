package thelper

import (
	"fmt"
	"io"
	"strings"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type TeaChoices struct {
	Choices  []string         // items on the to-do list
	Cursor   int              // which to-do list item our cursor is pointing at
	Selected map[int]struct{} // which to-do items are selected
}

func (tc TeaChoices) Init() tea.Cmd {
	return nil
}

func (tc TeaChoices) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:

		switch msg.String() {

		// These keys should exit the program.
		case "ctrl+c", "q":
			return tc, tea.Quit

		// The "up" and "k" keys move the cursor up
		case "up", "k":
			if tc.Cursor > 0 {
				tc.Cursor--
			}

		// The "down" and "j" keys move the cursor down
		case "down", "j":
			if tc.Cursor < len(tc.Choices)-1 {
				tc.Cursor++
			}

		// The "enter" key and the spacebar (a literal space) toggle
		// the selected state for the item that the cursor is pointing at.
		case "enter", " ":
			_, ok := tc.Selected[tc.Cursor]
			if ok {
				delete(tc.Selected, tc.Cursor)
			} else {
				tc.Selected[tc.Cursor] = struct{}{}
			}
		}
	}

	// Return the updated teachoices to the Bubble Tea runtime for processing.
	// Note that we're not returning a command.
	return tc, nil
}

func (tc TeaChoices) View() string {
	// The header
	s := "What should we buy at the market?\n\n"

	// Iterate over our choices
	for i, choice := range tc.Choices {

		// Is the cursor pointing at this choice?
		cursor := " " // no cursor
		if tc.Cursor == i {
			cursor = ">" // cursor!
		}

		// Is this choice selected?
		checked := " " // not selected
		if _, ok := tc.Selected[i]; ok {
			checked = "x" // selected!
		}

		// Render the row
		s += fmt.Sprintf("%s [%s] %s\n", cursor, checked, choice)
	}

	// The footer
	s += "\nPress q to quit.\n"

	// Send the UI for rendering
	return s
}

const listHeight = 14

var (
	TitleStyle        = lipgloss.NewStyle().MarginLeft(2)
	ItemStyle         = lipgloss.NewStyle().PaddingLeft(4)
	SelectedItemStyle = lipgloss.NewStyle().PaddingLeft(2).Foreground(lipgloss.Color("170"))
	PaginationStyle   = list.DefaultStyles().PaginationStyle.PaddingLeft(4)
	HelpStyle         = list.DefaultStyles().HelpStyle.PaddingLeft(4).PaddingBottom(1)
	QuitTextStyle     = lipgloss.NewStyle().Margin(1, 0, 2, 4)
)

type Item string

func (i Item) FilterValue() string { return "" }

type ItemDelegate struct{}

func (d ItemDelegate) Height() int                             { return 1 }
func (d ItemDelegate) Spacing() int                            { return 0 }
func (d ItemDelegate) Update(_ tea.Msg, _ *list.Model) tea.Cmd { return nil }
func (d ItemDelegate) Render(w io.Writer, m list.Model, index int, listItem list.Item) {
	i, ok := listItem.(Item)
	if !ok {
		return
	}

	str := fmt.Sprintf("%d. %s", index+1, i)

	fn := ItemStyle.Render
	if index == m.Index() {
		fn = func(s ...string) string {
			return SelectedItemStyle.Render("> " + strings.Join(s, " "))
		}
	}

	fmt.Fprint(w, fn(str))
}

type ListModel struct {
	List     list.Model
	Choice   string
	Quitting bool
}

func (lm ListModel) Init() tea.Cmd {
	return nil
}

func (lm ListModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		lm.List.SetWidth(msg.Width)
		return lm, nil

	case tea.KeyMsg:
		switch keypress := msg.String(); keypress {
		case "ctrl+c":
			lm.Quitting = true
			return lm, tea.Quit

		case "enter":
			i, ok := lm.List.SelectedItem().(Item)
			if ok {
				lm.Choice = string(i)
			}
			return lm, tea.Quit
		}
	}

	var cmd tea.Cmd
	lm.List, cmd = lm.List.Update(msg)
	return lm, cmd
}

func (lm ListModel) View() string {
	if lm.Choice != "" {
		return QuitTextStyle.Render(fmt.Sprintf("Selected %s", lm.Choice))
	}
	if lm.Quitting {
		return QuitTextStyle.Render("Aborting...")
	}
	return "\n" + lm.List.View()
}
