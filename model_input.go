package main

import (
	"fmt"
	"reflect"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type inputModel struct {
	ti textinput.Model
}

func newInput() inputModel {
	ti := textinput.New()
	ti.Prompt = "Name: "
	ti.Placeholder = "pollux"
	return inputModel{
		ti: ti,
	}
}

func (m inputModel) Init() tea.Cmd {
	return nil
}

func (m inputModel) Update(message tea.Msg) (Model, tea.Cmd) {
	logger.Debug(fmt.Sprintf("Workspace.Input.Update.typeof(message): %s", reflect.TypeOf(message).Name()))

	var cmd tea.Cmd
	m.ti, cmd = m.ti.Update(message)

	logger.Debug(fmt.Sprintf("Workspace.Input.Update.ti.Focused: %v", m.ti.Focused()))

	return m, cmd
}

func (m inputModel) View() string {
	logger.Debug(fmt.Sprintf("Workspace.Input.View.ti.Focused: %v", m.ti.Focused()))

	return m.ti.View()
}

// Focus sets the focus state on the model. When the model is in focus it can
// receive keyboard input and the cursor will be shown.
func (m inputModel) Focus() tea.Cmd {
	cmd := m.ti.Focus()

	logger.Debug(fmt.Sprintf("Workspace.Input.Focus.ti.Focused: %v", m.ti.Focused()))

	return cmd
}
