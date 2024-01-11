package main

import (
	"fmt"
	"reflect"

	tea "github.com/charmbracelet/bubbletea"
)

// because we want to have a form with several different input items we need
// to be able to activate (Focus) and deactivate (Blur).
// I tried to extend tea.Model but Update can't return a tea.Model because it
// doesn't implement the extended interface.
type Model interface {
	// Init is the first function that will be called. It returns an optional
	// initial command. To not perform an initial command return nil.
	Init() tea.Cmd

	// Update is called when a message is received. Use it to inspect messages
	// and, in response, update the model and/or send a command.
	Update(tea.Msg) (Model, tea.Cmd)

	// View renders the program's UI, which is just a string. The view is
	// rendered after every Update.
	View() string

	// Focus focuses the form item, e.g. enable cursor, switch styles, ...
	Focus() tea.Cmd
}

type workspaceModel struct {
	input Model
}

func newWorkspace() workspaceModel {
	return workspaceModel{
		input: newInput(),
	}
}

func (m workspaceModel) Init() tea.Cmd {
	var cmds []tea.Cmd
	if cmd := m.input.Focus(); cmd != nil {
		cmds = append(cmds, cmd)
	}
	if cmd := m.input.Init(); cmd != nil {
		cmds = append(cmds, cmd)
	}
	if len(cmds) > 1 {
		return tea.Batch(cmds...)
	}
	if len(cmds) > 0 {
		return cmds[0]
	}
	return nil
}

func (m workspaceModel) Update(message tea.Msg) (tea.Model, tea.Cmd) {
	logger.Debug(fmt.Sprintf("Workspace.Update.typeof(message): %s", reflect.TypeOf(message).Name()))

	var cmd tea.Cmd
	m.input, cmd = m.input.Update(message)
	return m, cmd
}

func (m workspaceModel) View() string {
	return m.input.View()
}
