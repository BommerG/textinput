package main

import (
	"fmt"
	"reflect"

	tea "github.com/charmbracelet/bubbletea"
)

type mainModel struct {
	workspace tea.Model
}

func initialModel() mainModel {
	return mainModel{
		workspace: newWorkspace(),
	}
}

func (m mainModel) Init() tea.Cmd {
	return m.workspace.Init()
}

func (m mainModel) Update(message tea.Msg) (tea.Model, tea.Cmd) {
	logger.Debug(fmt.Sprintf("Update.typeof(message): %s", reflect.TypeOf(message).Name()))

	switch msg := message.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter, tea.KeyCtrlC, tea.KeyEsc:
			return m, tea.Quit
		}
	}

	var cmd tea.Cmd
	m.workspace, cmd = m.workspace.Update(message)
	return m, cmd
}

func (m mainModel) View() string {
	return m.workspace.View()
}
