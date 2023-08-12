package thelper

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
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