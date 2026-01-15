package models

import tea "github.com/charmbracelet/bubbletea"

type BubbleTea struct {
	Choices  []string         // items on the to-do list
	Cursor   int              // which to-do list item our cursor is pointing at
	Selected map[int]struct{} // which to-do items are selected
}

func (m BubbleTea) Init() tea.Cmd {
	// Just return `nil`, which means "no I/O right now, please."
	return nil
}
